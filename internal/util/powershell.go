//go:build windows

package util

import (
	"os/exec"
	"syscall"
)

func getCmd() *exec.Cmd {
	cmd := exec.Command("powershell")
	// Hide powershell window
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: windowsCreateNoWindowFlag}

	return cmd
}
