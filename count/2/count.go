package count

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type counter struct {
	Input io.Reader
}

func NewCounter() *counter {
	return &counter{
		Input: os.Stdin,
	}
}

func (c *counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	return lines
}

func Main() {
	fmt.Println(NewCounter().Lines())
}
