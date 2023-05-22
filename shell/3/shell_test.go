package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"shell"

	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestCmdFromString_CreatesExpectedCmd(t *testing.T) {
	t.Parallel()
	cmd, err := shell.CmdFromString("/bin/ls -l main.go\n")
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(args, got) {
		t.Error(cmp.Diff(args, got))
	}
}

func TestNewSession_CreatesExpectedSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun_ProducesExpectedOutput(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("echo hello\n\n")
	out := new(bytes.Buffer)
	session := shell.NewSession(in, out, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nBe seeing you!\n"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
