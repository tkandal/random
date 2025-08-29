package random

import (
	"crypto/rand"
	"fmt"
)

func String(l int) string {
	b := make([]byte, l)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
