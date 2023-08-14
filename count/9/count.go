package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	files  []io.Reader
	input  io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return lines
}

func (c *counter) Words() int {
	words := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return words
}

func (c *counter) Bytes() int {
	bytes := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanBytes)
	for input.Scan() {
		bytes++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return bytes
}

func Main() int {
	lineMode := flag.Bool("lines", false, "Count lines, not words")
	byteMode := flag.Bool("bytes", false, "Count bytes, not words")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-lines | -bytes] [files...]\n", os.Args[0])
		fmt.Println("Counts words (or lines or bytes) in named files or standard input.\nFlags:")
		flag.PrintDefaults()
	}
	flag.Parse()
	c, err := NewCounter(
		WithInputFromArgs(flag.Args()),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	switch {
	case *lineMode && *byteMode:
		fmt.Fprintln(os.Stderr, "Please specify either '-lines' or '-bytes', but not both.")
		return 1
	case *lineMode:
		fmt.Println(c.Lines())
	case *byteMode:
		fmt.Println(c.Bytes())
	default:
		fmt.Println(c.Words())
	}
	return 0
}
