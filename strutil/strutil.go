package strutil

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
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

func EnsureSuffix(s, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}
	return s
}

func EnsurePrefix(s, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}
