package util

import (
	"crypto/sha256"
	"fmt"
)

func HashString(str string) string {
	data := []byte(str)

	hasher := sha256.New()

	hasher.Write(data)

	hash := hasher.Sum(nil)

	return fmt.Sprintf("%x", hash)
}
