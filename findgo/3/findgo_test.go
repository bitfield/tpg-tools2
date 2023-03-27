package findgo_test

import (
	"testing"

	"github.com/bitfield/findgo"
)

func TestFiles_CorrectlyCountsFilesInTree(t *testing.T) {
	t.Parallel()
	want := 4
	got := findgo.Files("testdata")
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
