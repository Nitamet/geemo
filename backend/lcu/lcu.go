// Package lcu provides a client to work with the LCU API
package lcu

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Nitamet/geemo/backend/shell"
	"io"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Game modes
const (
	gameModeNone    = "NONE"
	gameModeAram    = "ARAM"
	gameModeClassic = "CLASSIC"
)

// Summoner spell ids
const (
	SummonerSpellIdUrfMark  = 39
	SummonerSpellIdAramMark = 32
)

// LCU states
const (
	StateNotLaunched = "NotLaunched"
	StateNotInLobby  = "NotInLobby"
	StateInLobby     = "InLobby"
	StateInGame      = "InGame"
)

// Client represents a LCU client
type Client struct {
	port     int          // Port of the LCU instance
	token    string       // Token used to authenticate requests
	http     *http.Client // HTTP client configured to work with the LCU (ignores SSL errors, etc...)
	summoner *Summoner    // Logged in summoner
}

// Regexps used to get the port and token from the LCU command line arguments
var portRegExp = regexp.MustCompile(`--app-port=(\d+)`)
var tokenRegExp = regexp.MustCompile(`--remoting-auth-token=([a-zA-Z0-9-_]+)`)

// TryToGetLCU tries to get the LCU instance and returns a LCU client if it succeeds
func TryToGetLCU() *Client {
	cmdOutput, running := getLCUArgs()

	if !running {
		return nil
	}

	port, ok := getPortFromArgs(cmdOutput)
	if !ok {
		return nil
	}

	token, ok := getTokenFromArgs(cmdOutput)
	if !ok {
		return nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr, Timeout: 5 * time.Second}

	return &Client{
		port:  port,
		token: token,
		http:  httpClient,
	}
}

// getLCUArgs returns the command line arguments of the LCU instance and whether an instance is running
func getLCUArgs() (string, bool) {
	var output string

	switch os := runtime.GOOS; os {
	case "linux":
		output = shell.Execute("ps x -o args | grep 'LeagueClientUx.exe' | grep -v grep")
	case "windows":
		lcuCommand := "Get-CimInstance Win32_Process -Filter \"name = 'LeagueClientUX.exe'\" | Select-Object -ExpandProperty CommandLine"
		gameClientCommand := "Get-CimInstance Win32_Process -Filter \"name = 'League Of Legends.exe'\" | Select-Object -ExpandProperty CommandLine"

		output = shell.Execute(fmt.Sprintf("%s; %s", lcuCommand, gameClientCommand))
	default:
		log.Panicf("Unsupported OS %s", os)
	}

	if output == "" {
		return "", false
	}

	return output, true
}

// IsGameClientRunning checks whether the LoL Game Client is running or not
func (c *Client) isGameClientRunning() bool {
	var output string

	switch os := runtime.GOOS; os {
	case "linux":
		output = shell.Execute("ps x -o args | grep 'League of Legends.exe'")

		return strings.Contains(output, "-GameID")
	case "windows":
		output = shell.Execute("Get-CimInstance Win32_Process -Filter \"name = 'League Of Legends.exe'\" | Select-Object -ExpandProperty CommandLine")
	default:
		log.Panicf("Unsupported OS %s", os)
	}

	return output != ""
}

// getPortFromArgs returns the port of the LCU instance and whether it succeeded or not
func getPortFromArgs(args string) (int, bool) {
	portArg := portRegExp.FindStringSubmatch(args)

	if len(portArg) < 2 {
		return 0, false
	}

	port, _ := strconv.Atoi(portArg[1])

	return port, true
}

// getTokenFromArgs returns the token of the LCU instance and whether it succeeded or not
func getTokenFromArgs(args string) (string, bool) {
	tokenArg := tokenRegExp.FindStringSubmatch(args)

	if len(tokenArg) < 2 {
		return "", false
	}

	return tokenArg[1], true
}

// UpdateState updates the state of the LCU instance and returns current state
func (c *Client) UpdateState() string {
	if c.isGameClientRunning() {
		return StateInGame
	}

	_, running := getLCUArgs()
	if !running {
		return StateNotLaunched
	}

	ok := c.IsInLobby()
	if ok {
		return StateInLobby
	}

	return StateNotInLobby
}

// request sends a request to the LCU API
func (c *Client) request(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://127.0.0.1:%d/%s", c.port, path), body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Accept":        {"*/*"},
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("riot:"+c.token)))},
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// get sends a GET request to the LCU API
func (c *Client) get(path string) (*http.Response, error) {
	return c.request("GET", path, nil)
}

// post sends a POST request to the LCU API
func (c *Client) post(path string, body io.Reader) (*http.Response, error) {
	return c.request("POST", path, body)
}

// put sends a PUT request to the LCU API
func (c *Client) put(path string, body io.Reader) (*http.Response, error) {
	return c.request("PUT", path, body)
}

// patch sends a PATCH request to the LCU API
func (c *Client) patch(path string, body io.Reader) (*http.Response, error) {
	return c.request("PATCH", path, body)
}

// delete sends a DELETE request to the LCU API
func (c *Client) delete(path string) (*http.Response, error) {
	return c.request("DELETE", path, nil)
}

// Summoner represents a summoner
type Summoner struct {
	AccountId     int64  `json:"accountId"`
	Name          string `json:"displayName"`
	ProfileIconId int    `json:"profileIconId"`
	SummonerId    int64  `json:"summonerId"`
}

// CurrentSummoner returns the current logged in summoner
func (c *Client) CurrentSummoner() Summoner {
	if c.summoner != nil {
		return *c.summoner
	}

	var summoner Summoner

	resp, err := c.get("lol-summoner/v1/current-summoner")
	if err != nil {
		log.Panic(err)
	}

	defer closeBody(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(body, &summoner)
	if err != nil {
		log.Panic(err)
	}

	c.summoner = &summoner

	return summoner
}

// IsInLobby checks whether the current summoner is in lobby
func (c *Client) IsInLobby() bool {
	resp, err := c.get("lol-lobby/v2/lobby")

	// Don't panic if we get an error, it means that the client has not been started yet
	if err != nil {
		println(err.Error())
		return false
	}

	defer closeBody(resp)

	return resp.StatusCode == 200
}

// GetCurrentGameMode returns both the LCU game mode name and the formatted game mode name
func (c *Client) GetCurrentGameMode() (string, string) {
	resp, err := c.get("lol-lobby/v2/lobby")
	if err != nil {
		return gameModeNone, gameModeNone
	}

	defer closeBody(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var lobbyInfo struct {
		GameConfig struct {
			GameMode string `json:"gameMode"`
			PickType string `json:"pickType"`
		} `json:"gameConfig"`
	}

	err = json.Unmarshal(body, &lobbyInfo)
	if err != nil {
		log.Panic(err)
	}

	if lobbyInfo.GameConfig.GameMode == gameModeClassic {
		return lobbyInfo.GameConfig.GameMode, "Normal"
	}

	return lobbyInfo.GameConfig.GameMode, lobbyInfo.GameConfig.GameMode
}

// GetCurrentChampion returns the current champion id and whether a champion is selected
func (c *Client) GetCurrentChampion() (int, bool) {
	resp, err := c.get("lol-champ-select/v1/current-champion")
	if err != nil {
		log.Panic(err)
	}

	defer closeBody(resp)

	if resp.StatusCode != 200 {
		return -1, false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var championId int
	err = json.Unmarshal(body, &championId)
	if err != nil {
		log.Panic(err)
	}

	return championId, true
}

// GetAssignedRole returns the current assigned role and whether a role was assigned
func (c *Client) GetAssignedRole() (string, bool) {
	log.Println("Getting assigned role")

	resp, err := c.get("lol-lobby-team-builder/champ-select/v1/session")
	if err != nil {
		log.Panic(err)
	}

	defer closeBody(resp)

	if resp.StatusCode != 200 {
		log.Println("Can't get assigned role, got status code " + strconv.Itoa(resp.StatusCode))

		return "", false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var session struct {
		MyTeam []struct {
			AssignedPosition string `json:"assignedPosition"`
			SummonerId       int64  `json:"summonerId"`
		} `json:"myTeam"`
	}
	err = json.Unmarshal(body, &session)
	if err != nil {
		log.Panic(err)
	}

	position := ""
	for _, summoner := range session.MyTeam {
		if summoner.SummonerId == c.CurrentSummoner().SummonerId {
			position = summoner.AssignedPosition
			break
		}
	}

	if position == "" {
		log.Println("No role assigned")

		return "", false
	}

	log.Println("Assigned role: " + position)

	return position, true
}

// ApplySummonerSpells applies the given summoner spells
func (c *Client) ApplySummonerSpells(firstSpellId int, secondSpellId int) {
	// TODO: Temporary fix for ARAM
	if firstSpellId == SummonerSpellIdUrfMark {
		firstSpellId = SummonerSpellIdAramMark
	}
	if secondSpellId == SummonerSpellIdUrfMark {
		secondSpellId = SummonerSpellIdAramMark
	}

	spellsJson := struct {
		Spell1Id int `json:"spell1Id"`
		Spell2Id int `json:"spell2Id"`
	}{
		Spell1Id: firstSpellId,
		Spell2Id: secondSpellId,
	}

	encodedSpells, _ := json.Marshal(spellsJson)
	resp, err := c.patch("lol-champ-select/v1/session/my-selection", bytes.NewBuffer(encodedSpells))
	if err != nil {
		log.Panic(err)
	}

	defer closeBody(resp)

	if resp.StatusCode != 204 {
		log.Panicf("Error while applying summoner spells, got status code %d", resp.StatusCode)
	}

	log.Println("Applied summoner spells")
}

// Item represents a LoL item
type Item struct {
	Count int    `json:"count"`
	Id    string `json:"id"`
}

// ItemBlock represents a group of items
type ItemBlock struct {
	Items []Item `json:"items"`
	Type  string `json:"type"`
}

// ItemSet represents a set of item blocks associated with a champions and maps
type ItemSet struct {
	AssociatedChampions []int       `json:"associatedChampions"`
	AssociatedMaps      []int       `json:"associatedMaps"`
	Blocks              []ItemBlock `json:"blocks"`
	Title               string      `json:"title"`
}

// ItemSets represents a group of item sets
type ItemSets struct {
	AccountId int64     `json:"accountId"`
	ItemSets  []ItemSet `json:"itemSets"`
	Timestamp int64     `json:"timestamp"`
}

// ApplyItemSet applies the given item set
func (c *Client) ApplyItemSet(itemSet ItemSet) {
	currentSummoner := c.CurrentSummoner()

	itemSets := ItemSets{
		AccountId: currentSummoner.AccountId,
		ItemSets:  []ItemSet{itemSet},
		Timestamp: time.Now().Unix(),
	}

	encodedItemSets, _ := json.Marshal(itemSets)
	println(fmt.Sprintf("Encoded item sets: %s", encodedItemSets))
	resp, err := c.put(fmt.Sprintf("lol-item-sets/v1/item-sets/%d/sets", c.CurrentSummoner().SummonerId), bytes.NewBuffer(encodedItemSets))
	if err != nil {
		log.Panic()
	}

	defer closeBody(resp)

	if resp.StatusCode != 201 {
		log.Panicf("Error while applying item set, got status code %d", resp.StatusCode)
	}

	log.Println("Applied item set")
}

// RunePage represents a LoL rune page
type RunePage struct {
	Name            string `json:"name"`
	PrimaryStyleId  int    `json:"primaryStyleId"`
	SubStyleId      int    `json:"subStyleId"`
	SelectedPerkIds []int  `json:"selectedPerkIds"`
	Current         bool   `json:"current"`
}

// ApplyRunes applies the given rune page
func (c *Client) ApplyRunes(runes RunePage) {
	c.deleteCurrentRunePage()

	encodedRunes, _ := json.Marshal(runes)
	println(fmt.Sprintf("Encoded runes: %s", encodedRunes))
	resp, err := c.post("lol-perks/v1/pages", bytes.NewBuffer(encodedRunes))
	if err != nil {
		log.Panic(err)
	}

	defer closeBody(resp)

	if resp.StatusCode != 200 {
		log.Panicf("Error while applying runes, got status code %d", resp.StatusCode)
	}

	log.Println("Applied runes")
}

// deleteCurrentRunePage deletes the current rune page
func (c *Client) deleteCurrentRunePage() {
	resp, err := c.get("lol-perks/v1/currentpage")
	if err != nil {
		log.Panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var currentPage struct {
		Id int64 `json:"id"`
	}
	err = json.Unmarshal(body, &currentPage)
	if err != nil {
		log.Panic(err)
	}

	resp, err = c.delete(fmt.Sprintf("lol-perks/v1/pages/%d", currentPage.Id))
	if err != nil {
		log.Panic(err)
	}

	defer closeBody(resp)

	if resp.StatusCode != 204 {
		log.Printf("Error while deleting current rune page, got status code %d", resp.StatusCode)
	}

	log.Println("Deleted current rune page")
}

func closeBody(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Panic(err)
	}
}
