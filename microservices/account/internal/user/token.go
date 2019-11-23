package user

import "crypto/rand"

func generateToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return string(b)
}
