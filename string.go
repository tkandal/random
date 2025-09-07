package random

import (
	"fmt"
)

const (
	// MaxStringLen is the maximum length of a random string that can be generated.
	MaxStringLen = 256
)

// String returns a random hexadecimal string of length l.
// If l is odd, it is rounded up to the next even number.
// The maximum length is truncated to MaxStringLen.
func String(l int) string {
	n := l
	if n%2 != 0 {
		n++
	}
	if n > MaxStringLen {
		n = MaxStringLen
	}
	n /= 2
	return fmt.Sprintf("%x", Bytes(n))
}

// SecureString returns a cryptographically secure random hexadecimal string of length l.
// If l is odd, it is rounded up to the next even number.
// The maximum length is truncated to MaxStringLen.
func SecureString(l int) (string, error) {
	n := l
	if n%2 != 0 {
		n++
	}
	if n > MaxStringLen {
		n = MaxStringLen
	}
	n /= 2
	b, err := SecureBytes(n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
