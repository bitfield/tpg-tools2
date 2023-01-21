package count_test

import (
	"bytes"
	"github.com/bitfield/count"
	"testing"
)

func TestLinesCountsLinesInInput(t *testing.T) {
	t.Parallel()
	c := count.NewCounter()
	c.Input = bytes.NewBufferString("1\n2\n3")
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
