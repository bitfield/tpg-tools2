package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type matcher struct {
	input  io.Reader
	output io.Writer
	text   string
}

type option func(*matcher) error

func WithInput(input io.Reader) option {
	return func(c *matcher) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithSearchTextFromArgs(args []string) option {
	return func(m *matcher) error {
		if len(args) < 1 {
			return nil
		}
		m.text = args[0]
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(m *matcher) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		m.output = output
		return nil
	}
}

func NewMatcher(opts ...option) (*matcher, error) {
	m := &matcher{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (m *matcher) PrintMatchingLines() {
	input := bufio.NewScanner(m.input)
	for input.Scan() {
		if strings.Contains(input.Text(), m.text) {
			fmt.Fprintln(m.output, input.Text())
		}
	}
}

func Main() {
	m, err := NewMatcher(
		WithSearchTextFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	m.PrintMatchingLines()
}
