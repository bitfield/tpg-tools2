package findgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
