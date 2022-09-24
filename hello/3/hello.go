package hello3

import (
	"fmt"
	"io"
)

func PrintTo(w io.Writer) {
	fmt.Fprintln(w, "Hello, world")
}
