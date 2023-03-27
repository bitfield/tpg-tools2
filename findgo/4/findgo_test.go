package findgo_test

import (
	"archive/zip"
	"os"
	"testing"
	"testing/fstest"

	"github.com/bitfield/findgo"
)

func TestFiles_CorrectlyCountsFilesInTree(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/tree")
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFiles_CorrectlyCountsFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFiles_CorrectlyCountsFilesInZIPArchive(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/files.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findgo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findgo.Files(fsys)
	}
}
