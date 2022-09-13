package hello

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GreetUser(input io.Reader, output io.Writer) {
	name := "you"
	fmt.Fprintln(output, "What Is Your Name")
	scanner := bufio.NewScanner(input)
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Fprintf(output, "Hello, %s\n", name)
}

func Main() int {
	GreetUser(os.Stdin, os.Stdout)
	return 0
}