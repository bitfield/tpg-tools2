package writer

import (
	"flag"
	"fmt"
	"os"
)

const Usage = `Usage: writefile -size SIZE_BYTES PATH

Creates the file PATH, containing SIZE_BYTES bytes, all zero.

Example: writefile -size 1000 zeroes.dat`

func Main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return
	}
	size := flag.Int("size", 0, "Size in bytes")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	err := WriteToFile(flag.Args()[0], make([]byte, *size))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0o600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0o600)
}
