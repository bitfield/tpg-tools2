package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/bitfield/shell"
	"github.com/rogpeppe/go-internal/testscript"

	"github.com/google/go-cmp/cmp"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"shell": shell.Main,
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

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
		Stdin:      os.Stdin,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		DryRun:     false,
		Transcript: io.Discard,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRunProducesExpectedOutput(t *testing.T) {
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

func TestRunProducesExpectedTranscript(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("echo hello\n\n")
	transcript := new(bytes.Buffer)
	session := shell.NewSession(in, io.Discard, io.Discard)
	session.DryRun = true
	session.Transcript = transcript
	session.Run()
	want := "> echo hello\necho hello\n> \n> \nBe seeing you!\n"
	got := transcript.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
