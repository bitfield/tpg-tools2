package findgo

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Files(path string) (count int) {
	fsys := os.DirFS(path)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			count++
		}
		return nil
	})
	return count
}
