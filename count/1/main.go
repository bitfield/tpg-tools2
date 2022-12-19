package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := 0
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines++
	}
	fmt.Println(lines)
}
