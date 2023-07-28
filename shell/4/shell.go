package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	Transcript     io.Writer
	Combined       io.Writer
	DryRun         bool
	TranscriptPath string
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Stdin:      in,
		Stdout:     out,
		Stderr:     errs,
		DryRun:     false,
		Transcript: io.Discard,
	}
}

func (s *Session) Run() {
	s.Combined = io.MultiWriter(s.Stdout, s.Transcript)
	fmt.Fprintf(s.Combined, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Fprintln(s.Transcript, line)
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(s.Combined, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Combined, "%s\n> ", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
			fmt.Fprintln(s.Transcript, "error:", err)
		}
		fmt.Fprintf(s.Combined, "%s> ", output)
	}
	fmt.Fprintln(s.Combined, "\nBe seeing you!")
}

func CmdFromString(cmdLine string) (*exec.Cmd, error) {
	args := strings.Fields(cmdLine)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func Main() int {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	transcript, err := os.Create("transcript.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer transcript.Close()
	session.Transcript = transcript
	session.Run()
	fmt.Println("[output file is 'transcript.txt']")
	return 0
}
