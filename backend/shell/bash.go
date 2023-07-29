//go:build linux

package shell

import "os/exec"

func Execute(command string) string {
	cmd := exec.Command("bash", "-c", command)

	cmdOutput, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(cmdOutput)
}
