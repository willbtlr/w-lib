package slices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Choose[T fmt.Stringer](xs []T) (*T, bool) {
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

		return &xs[n-1], true
	}
}

func Find[T any](xs []T, f func(T) bool) (T, bool) {
	for _, x := range xs {
		if f(x) {
			return x, true
		}
	}

	return nil, false
}

func FindIndex[T any](xs []T, f func(T) bool) (int, bool) {
	for i, x := range xs {
		if f(x) {
			return i, true
		}
	}

	return -1, false
}

func Filter[T any](xs []T, f func(T) bool) []T {
	var ret []T
	for _, x := range xs {
		if f(x) {
			ret = append(ret, x)
		}
	}
	return ret
}

func Map[T any, U any](xs []T, f func(T) U) []U {
	var ret []U
	for _, x := range xs {
		ret = append(ret, f(x))
	}
	return ret
}

func MapWithIndex[T any, U any](xs []T, f func(int, T) U) []U {
	var ret []U
	for i, x := range xs {
		ret = append(ret, f(i, x))
	}
	return ret
}

func Reduce[T any, U any](xs []T, f func(U, T) U, init U) U {
	ret := init
	for _, x := range xs {
		ret = f(ret, x)
	}
	return ret
}

func ReduceWithIndex[T any, U any](xs []T, f func(int, U, T) U, init U) U {
	ret := init
	for i, x := range xs {
		ret = f(i, ret, x)
	}
	return ret
}

func Any[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}
	return false
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

func Contains[T comparable](xs []T, x T) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}
