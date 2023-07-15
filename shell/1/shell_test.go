package shell_test

import (
	"testing"

	"github.com/bitfield/shell"

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
	cmd, err := shell.CmdFromString("/bin/ls -l main.go")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
