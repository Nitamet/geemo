//go:build linux

package util

import "os/exec"

func getCmd() *exec.Cmd {
	return exec.Command("bash")
}
