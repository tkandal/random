package random

import (
	"crypto/rand"
	"encoding/base64"
)

func String(l int) string {
	b := make([]byte, l)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
