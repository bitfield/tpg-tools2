package hello3_test

import (
	"bytes"
	hello "hello3"
	"testing"
)

func TestPrintTo_PrintsHelloMessageToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
