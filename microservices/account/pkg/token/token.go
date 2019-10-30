package token

import (
	"crypto/rand"
)

type Token struct{}

func (Token) Generate() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return string(b)
}
