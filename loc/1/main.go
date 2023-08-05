package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func main() {
	script.FindFiles(".").
		MatchRegexp(regexp.MustCompile(".go$")).Stdout()
	lines, err := script.FindFiles(".").
		MatchRegexp(regexp.MustCompile(".go$")).
		Concat().
		RejectRegexp(regexp.MustCompile("^$")).
		CountLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("You've written %d lines of Go in this project. Nice work!\n", lines)
}
