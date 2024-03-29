package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("output.dat")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	err = write(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type safeWriter struct {
	w     io.Writer
	Error error
}

func (sw *safeWriter) Write(data []byte) {
	if sw.Error != nil {
		return
	}
	_, err := sw.w.Write(data)
	if err != nil {
		sw.Error = err
	}
}

func write(w io.Writer) error {
	metadata := []byte("hello\n")
	sw := safeWriter{w: w}
	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)
	return sw.Error
}
