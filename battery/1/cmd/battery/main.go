package main

import (
	"fmt"
	"os"

	"github.com/bitfield/battery"
)

func main() {
	status, err := battery.GetStatus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't read battery status: %v", err)
	}
	fmt.Printf("Battery %d%% charged\n", status.ChargePercent)
}
