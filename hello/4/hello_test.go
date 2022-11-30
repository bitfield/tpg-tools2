package hello4_test

import (
	"bytes"
	hello "hello4"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	fakeTerminal := &bytes.Buffer{}
	p := &hello.Printer{
		Output: fakeTerminal,
	}
	p.Print()
	want := "Hello, world\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
