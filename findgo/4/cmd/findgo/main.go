package main

import (
	"fmt"
	"os"

	"github.com/bitfield/findgo"
)

func main() {
	fmt.Println(findgo.Files(os.DirFS(os.Args[1])))
}
