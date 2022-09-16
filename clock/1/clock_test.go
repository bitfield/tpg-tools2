package clock_test

import (
	"testing"
	"time"

	"clock"

	"github.com/google/go-cmp/cmp"
)

func TestFormatReturnsGivenTimeAsString(t *testing.T) {
	t.Parallel()
	testTime, err := time.Parse(time.RFC3339, "2022-09-16T15:18:23Z")
	if err != nil {
		t.Fatal(err)
	}
	want := "It's 18 minutes past 15."
	got := clock.Format(testTime)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
