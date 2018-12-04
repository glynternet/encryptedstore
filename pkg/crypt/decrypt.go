package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"

	"github.com/pkg/errors"
)

// Decryptor is used to decrypt a payload with a given key
type Decryptor interface {
	Decrypt(key, payload []byte) ([]byte, error)
}

// DecryptorFn is any pure function or closure that can be used to decrypt a payload with a given key
type DecryptorFn func(key, payload []byte) ([]byte, error)

// Decrypt ensures that DecryptorFn satistfies the Decryptor interface
func (dfn DecryptorFn) Decrypt(key, payload []byte) ([]byte, error) {
	return dfn(key, payload)
}

// Decrypt decrypts an aes payload using the given key
func Decrypt(key, payload []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "creating new aws cipher")
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(c, iv[:])

	encryptedIn := bytes.NewBuffer(payload)

	var decryptedOut bytes.Buffer
	writer2 := &cipher.StreamWriter{S: stream, W: &decryptedOut}
	if _, err := io.Copy(writer2, encryptedIn); err != nil {
		return nil, errors.Wrap(err, "copying buffer")
	}

	return decryptedOut.Bytes(), nil
}
