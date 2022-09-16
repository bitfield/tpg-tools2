package clock

import (
	"fmt"
	"time"
)

func Format(tm time.Time) string {
	return fmt.Sprintf("It's %d minutes past %d.", tm.Minute(), tm.Hour())
}
