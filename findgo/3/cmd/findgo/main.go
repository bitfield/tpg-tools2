package main

import (
	"fmt"
	"os"

	"github.com/bitfield/findgo"
)

func main() {
	paths := findgo.Files(os.Args[1])
	for _, p := range paths {
		fmt.Println(p)
	}
}
