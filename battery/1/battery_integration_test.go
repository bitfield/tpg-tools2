//go:build integration

package battery_test

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/bitfield/battery"
)

func TestGetPmsetOutput_CapturesCmdOutput(t *testing.T) {
	t.Parallel()
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		t.Skipf("unable to run 'pmset' command: %v", err)
	}
	if !bytes.Contains(data, []byte("InternalBattery")) {
		t.Skip("no battery fitted")
	}
	text, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParsePmsetOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge: %d%%", status.ChargePercent)
}
