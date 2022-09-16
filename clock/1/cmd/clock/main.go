package main

import (
	"fmt"
	"time"

	"clock"
)

func main() {
	fmt.Println(clock.Format(time.Now()))
}
