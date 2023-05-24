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

func (c *Client) get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%d/%s", c.port, path), nil)
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

func (c *Client) post(path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("https://127.0.0.1:%d/%s", c.port, path), body)
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

func (c *Client) delete(path string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://127.0.0.1:%d/%s", c.port, path), nil)
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

type Summoner struct {
	Name          string `json:"displayName"`
	ProfileIconId int    `json:"profileIconId"`
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
	return true
	resp, err := c.get("lol-lobby/v2/lobby")
	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}

func (c *Client) SelectedChampion() (int, bool) {
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
