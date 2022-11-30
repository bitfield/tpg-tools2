package count2_test

import (
	"bytes"
	count "count2"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	c := count.NewCounter()
	c.Input = bytes.NewBufferString("1\n2\n3")
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
