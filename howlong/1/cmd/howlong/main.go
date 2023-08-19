package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bitfield/howlong"
)

const Usage = `Usage: howlong COMMAND [ARGS...]

Runs COMMAND with ARGS and reports the elapsed wall-clock time.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		os.Exit(0)
	}
	elapsed, err := howlong.Run(os.Args[1], os.Args[2:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("(time: %s)\n", elapsed.Round(time.Millisecond))
}
