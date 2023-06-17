package lcu

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

const (
	gameModeNone    = "NONE"
	gameModeAram    = "ARAM"
	gameModeClassic = "CLASSIC"
)

const (
	blindPickStrategy = "SimulPickStrategy"
	draftPickStrategy = "DraftModeSinglePickStrategy"
	blindPickName     = "Blind Pick"
	draftPickName     = "Draft Pick"
)

type Client struct {
	port  int
	token string
	http  *http.Client
}

func TryToGetLCU() *Client {
	cmdOutput := lookForLCUInstance()

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

func lookForLCUInstance() string {
	cmd := exec.Command("bash", "-c", "ps x -o args | grep 'LeagueClientUx'")
	bytes, _ := cmd.Output()

	return string(bytes)
}

func getPortFromArgs(args string) (int, bool) {
	pattern := regexp.MustCompile("--app-port=([0-9]{1,5})")
	portArg := pattern.FindStringSubmatch(args)

	if len(portArg) < 2 {
		return 0, false
	}

	port, _ := strconv.Atoi(portArg[1])

	return port, true
}

func getTokenFromArgs(args string) (string, bool) {
	pattern := regexp.MustCompile("--remoting-auth-token=([a-zA-Z0-9-_]+)")
	tokenArg := pattern.FindStringSubmatch(args)

	if len(tokenArg) < 2 {
		return "", false
	}

	return tokenArg[1], true
}

func (c *Client) UpdateState() string {
	cmdOutput := lookForLCUInstance()
	if cmdOutput == "" {
		return "NotLaunched"
	}

	ok := c.IsInLobby()
	if ok {
		return "InLobby"
	}

	return "NotInLobby"
}

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

func (c *Client) get(path string) (*http.Response, error) {
	return c.request("GET", path, nil)
}

func (c *Client) post(path string, body io.Reader) (*http.Response, error) {
	return c.request("POST", path, body)
}

func (c *Client) put(path string, body io.Reader) (*http.Response, error) {
	return c.request("PUT", path, body)
}

func (c *Client) patch(path string, body io.Reader) (*http.Response, error) {
	return c.request("PATCH", path, body)
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.request("DELETE", path, nil)
}

type Summoner struct {
	AccountId     int64  `json:"accountId"`
	Name          string `json:"displayName"`
	ProfileIconId int    `json:"profileIconId"`
	SummonerId    int64  `json:"summonerId"`
}

func (c *Client) CurrentSummoner() Summoner {
	var summoner Summoner
	println("Get current summoner")
	resp, _ := c.get("lol-summoner/v1/current-summoner")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &summoner)
	if err != nil {
		log.Fatalln(err)
	}
	println(fmt.Sprintf("Summoner: %v", summoner))
	return summoner
}

func (c *Client) IsInLobby() bool {
	resp, err := c.get("lol-lobby/v2/lobby")
	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}

func (c *Client) GetCurrentGameMode() (string, string) {
	resp, err := c.get("lol-lobby/v2/lobby")
	if err != nil {
		return gameModeNone, gameModeNone
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var lobbyInfo struct {
		GameConfig struct {
			GameMode string `json:"gameMode"`
			PickType string `json:"pickType"`
		} `json:"gameConfig"`
	}

	err = json.Unmarshal(body, &lobbyInfo)

	if lobbyInfo.GameConfig.GameMode == gameModeClassic {
		pickStrategy := ""
		switch lobbyInfo.GameConfig.PickType {
		case blindPickStrategy:
			pickStrategy = blindPickName
		case draftPickStrategy:
			pickStrategy = draftPickName
		}

		return lobbyInfo.GameConfig.GameMode, fmt.Sprintf("%s (%s)", "Normal", pickStrategy)
	}

	return lobbyInfo.GameConfig.GameMode, lobbyInfo.GameConfig.GameMode
}

func (c *Client) SelectedChampion() (int, bool) {
	return 266, true
	resp, _ := c.get("lol-champ-select/v1/current-champion")

	if resp.StatusCode != 200 {
		return -1, false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var championId int
	err = json.Unmarshal(body, &championId)
	if err != nil {
		log.Fatalln(err)
	}

	return championId, true
}

func (c *Client) ApplySummonerSpells(firstSpellId int, secondSpellId int) error {
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
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("error applying summoner spells")
	}

	return nil
}

type Item struct {
	Count int    `json:"count"`
	Id    string `json:"id"`
}
type ItemBlock struct {
	Items []Item `json:"items"`
	Type  string `json:"type"`
}
type ItemSet struct {
	AssociatedChampions []int       `json:"associatedChampions"`
	AssociatedMaps      []int       `json:"associatedMaps"`
	Blocks              []ItemBlock `json:"blocks"`
	Title               string      `json:"title"`
}
type ItemSets struct {
	AccountId int64     `json:"accountId"`
	ItemSets  []ItemSet `json:"itemSets"`
	Timestamp int64     `json:"timestamp"`
}

func (c *Client) ApplyItemSet(itemSet ItemSet) error {
	println("Apply item set")

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
		println(err.Error())
		return err
	}

	if resp.StatusCode != 201 {
		return errors.New("could not apply item set")
	}

	println("Applied item set")

	return nil
}

type RunePage struct {
	Name            string `json:"name"`
	PrimaryStyleId  int    `json:"primaryStyleId"`
	SubStyleId      int    `json:"subStyleId"`
	SelectedPerkIds []int  `json:"selectedPerkIds"`
	Current         bool   `json:"current"`
}

func (c *Client) ApplyRunes(runes RunePage) error {
	c.deleteCurrentRunePage()

	encodedRunes, _ := json.Marshal(runes)
	println(fmt.Sprintf("Encoded runes: %s", encodedRunes))
	resp, err := c.post("lol-perks/v1/pages", bytes.NewBuffer(encodedRunes))
	if err != nil {
		println("return")
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("could not apply runes")
	}

	return nil
}

func (c *Client) deleteCurrentRunePage() {
	resp, err := c.get("lol-perks/v1/currentpage")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var currentPage struct {
		Id int `json:"id"`
	}
	err = json.Unmarshal(body, &currentPage)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err = c.delete(fmt.Sprintf("lol-perks/v1/pages/%d", currentPage.Id))
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 204 {
		log.Fatalln("could not delete current rune page")
	}
}
