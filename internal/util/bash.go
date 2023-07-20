//go:build linux

package util

import "os/exec"

func Execute(command string) string {
	cmd := exec.Command("bash", "-c", command)

	cmdOutput, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(cmdOutput)
}
