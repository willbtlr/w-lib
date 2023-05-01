package strutil

import (
	"crypto/rand"
	"encoding/hex"
)

func RandHex(n int) (string, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	return hex.EncodeToString(buf), err
}

func MustRandHex(n int) string {
	s, err := RandHex(n)
	if err != nil {
		panic(err)
	}
	return s
}

