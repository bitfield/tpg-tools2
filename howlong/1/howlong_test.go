package howlong_test

import (
	"testing"
	"time"

	"github.com/bitfield/howlong"
)

func TestRunReportsCorrectElapsedTimeForCommand(t *testing.T) {
	t.Parallel()
	target := 100 * time.Millisecond
	elapsed, err := howlong.Run("sleep", "0.1")
	if err != nil {
		t.Fatal(err)
	}
	epsilon := target - elapsed
	if epsilon.Abs() > 300*time.Millisecond {
		t.Fatalf("want %s, got %s (not close enough)", target, elapsed)
	}
}
