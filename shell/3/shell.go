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
	DryRun         bool
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Stdin:  in,
		Stdout: out,
		Stderr: errs,
		DryRun: false,
	}
}

func (s *Session) Run() {
	fmt.Fprintf(s.Stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(s.Stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s\n> ", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s> ", output)
	}
	fmt.Fprintln(s.Stdout, "\nBe seeing you!")
}

func CmdFromString(cmdLine string) (*exec.Cmd, error) {
	args := strings.Fields(cmdLine)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func Main() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
