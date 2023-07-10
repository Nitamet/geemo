package util

import (
	"io"
	"log"
	"os/exec"
	"runtime"
)

const windowsCreateNoWindowFlag = 0x08000000

type Shell struct {
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Os     string
}

func CreateShell() *Shell {
	os := runtime.GOOS

	cmd := getCmd()

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
	newline := "\n"

	// Clear previous output
	switch s.Os {
	case "windows":
		newline = "\r\n"
		_, err := s.Stdin.Write([]byte("cls" + newline))
		if err != nil {
			log.Fatalln(err)
		}
	case "linux":
		_, err := s.Stdin.Write([]byte("clear" + newline))
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("Unsupported OS")
	}

	_, err := s.Stdin.Write([]byte(command + newline))
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
