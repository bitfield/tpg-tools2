package writer

import (
	"flag"
	"fmt"
	"os"
)

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0600)
}

func Main() int {
	size := flag.Int("size", 0, "Size in bytes")
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		return 1
	}
	err := WriteToFile(flag.Args()[0], make([]byte, *size))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return 0
}
