package findgo_test

import (
	"testing"

	"github.com/bitfield/findgo"
	"github.com/google/go-cmp/cmp"
)

func TestFilesCorrectlyListsFilesInTree(t *testing.T) {
	t.Parallel()
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files("testdata/tree")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
