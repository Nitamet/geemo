//go:build windows

package util

import (
	"os/exec"
	"syscall"
)

const windowsCreateNoWindowFlag = 0x08000000

func Execute(command string) string {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: windowsCreateNoWindowFlag}

	cmdOutput, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(cmdOutput)
}
