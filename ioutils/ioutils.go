package ioutils

import (
	"bufio"
	"os"
)

func StdinLines() (ret []string, err error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return
	}

	if info.Size() == 0 {
		return nil, nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return
}
