package crypt

import (
	"crypto/rand"
	"github.com/pkg/errors"
)

func generateNewKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, errors.Wrap(err, "reading random")
	}
	return key, nil
}

