package lcu

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
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
	httpClient := &http.Client{Transport: tr}

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
	pattern := regexp.MustCompile("--remoting-auth-token=([a-zA-Z0-9-]+)")
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

	return "NotInLobby"
}

func (c *Client) get(path string) []byte {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%d/%s", c.port, path), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Accept":        {"*/*"},
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("riot:"+c.token)))},
	}

	resp, err := c.http.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

type Summoner struct {
	Name          string `json:"displayName"`
	ProfileIconId int    `json:"profileIconId"`
}

func (c *Client) CurrentSummoner() Summoner {
	var summoner Summoner

	data := c.get("lol-summoner/v1/current-summoner")
	err := json.Unmarshal(data, &summoner)
	if err != nil {
		log.Fatalln(err)
	}

	return summoner
}
