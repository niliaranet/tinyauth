package server

import (
	"crypto/rand"
	"encoding/base32"
	"io"
)

func GenerateSecret(length int) ([]byte, error) {
	secretBytes := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, secretBytes)
	if err != nil {
		return nil, err
	}

	encoding := base32.StdEncoding.WithPadding(base32.NoPadding)
	secret := []byte(encoding.EncodeToString(secretBytes))
	return secret, nil
}
