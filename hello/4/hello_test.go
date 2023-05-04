package hello_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/hello"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	p := &hello.Printer{
		Output: buf,
	}
	p.Print()
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
