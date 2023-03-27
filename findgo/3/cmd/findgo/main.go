package main

import (
	"fmt"
	"os"

	"github.com/bitfield/findgo"
)

func main() {
	fmt.Println(findgo.Files(os.Args[1]))
}
