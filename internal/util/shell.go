package util

import (
	"io"
	"log"
	"os/exec"
	"runtime"
	"syscall"
)

const windowsCreateNoWindowFlag = 0x08000000

type Shell struct {
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Os     string
}

func CreateShell() *Shell {
	var cmd *exec.Cmd

	os := runtime.GOOS

	switch os {
	case "windows":
		cmd = exec.Command("powershell")
		// Hide powershell window
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: windowsCreateNoWindowFlag}
	case "linux":
		cmd = exec.Command("bash")
	default:
		log.Fatalln("Unsupported OS")
	}

	shellStdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}

	shellStdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	return &Shell{
		Cmd:    cmd,
		Stdin:  shellStdin,
		Stdout: shellStdout,
		Os:     os,
	}
}

func (s *Shell) ExecuteCommand(command string) string {
	// Clear previous output
	switch s.Os {
	case "windows":
		_, err := s.Stdin.Write([]byte("cls\r\n"))
		if err != nil {
			log.Fatalln(err)
		}
	case "linux":
		_, err := s.Stdin.Write([]byte("clear\r\n"))
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("Unsupported OS")
	}

	_, err := s.Stdin.Write([]byte(command + "\r\n"))
	if err != nil {
		log.Fatalln(err)
	}

	buffer := make([]byte, 16384)

	// We can't use ReadAll because it waits for the end of the stream
	// and the stream will never end because the shell is still running
	_, err = s.Stdout.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	return string(buffer)
}

func (s *Shell) Close() error {
	err := s.Stdin.Close()
	if err != nil {
		return err
	}

	err = s.Stdout.Close()
	if err != nil {
		return err
	}

	return nil
}
