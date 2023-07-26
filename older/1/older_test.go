package older_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/bitfield/older"
	"github.com/google/go-cmp/cmp"
)

func TestFiles_ReturnsFilesOlderThanGivenDuration(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {ModTime: time.Now()},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
	}
	got := older.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
