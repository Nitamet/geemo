package lcu

import (
	"os/exec"
	"regexp"
	"strconv"
)

type Client struct {
	port  int
	token string
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

	return &Client{
		port:  port,
		token: token,
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
	pattern := regexp.MustCompile("--riotclient-auth-token=([a-zA-Z0-9-]+)")
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
