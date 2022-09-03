package sliceutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Choose[T fmt.Stringer](xs []T) (T, bool) {
	s := bufio.NewScanner(os.Stdin)
	for {
		for i, x := range xs {
			fmt.Printf("%d: %s", i+1, x)
		}

		fmt.Print("> ")
		s.Scan()
		line := s.Text()

		if line == "q" {
			return nil, false
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("invalid number")
			continue
		}

		if n > 0 && n < len(xs)-1 {
			fmt.Println("choice out of range")
			continue
		}

		return xs[n-1], true
	}
}

