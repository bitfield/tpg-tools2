package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	var count int
	fsys := os.DirFS("testdata/tree")
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			count++
		}
		return nil
	})
	fmt.Println(count)
}
