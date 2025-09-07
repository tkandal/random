package random

import (
	"crypto/rand"
	"os"
)

const (
	devRandom = "/dev/random"
)

// Bytes returns a slice of random bytes of length l.
// The maximum length is truncated to MaxStringLen.
func Bytes(l int) []byte {
	n := l
	if n > MaxStringLen {
		n = MaxStringLen
	}
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}

// SecureBytes returns a slice of cryptographically secure random bytes of length l.
// The maximum length is truncated to MaxStringLen.
func SecureBytes(l int) ([]byte, error) {
	f, err := os.OpenFile(devRandom, os.O_RDONLY, os.FileMode(0444))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	n := l
	if n > MaxStringLen {
		n = MaxStringLen
	}
	b := make([]byte, n)
	o, err := f.Read(b)
	if err != nil {
		return nil, err
	}
	var m int
	for o < n {
		m, err = f.Read(b[o:])
		if err != nil {
			return nil, err
		}
		o += m
	}
	return b, nil
}
