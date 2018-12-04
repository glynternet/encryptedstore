package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"

	"github.com/pkg/errors"
)

// Encryptor is used to enecrypt a payload, returning the encrypted payload and key
type Encryptor interface {
	Encrypt(payload []byte) (encrypted, key []byte, err error)
}

// EncryptorFn is any pure function or closure that can be used to enecrypt a
// payload, returning the encrypted payload and key
type EncryptorFn func(payload []byte) (encrypted, key []byte, err error)

// Encrypt ensures that EncryptorFn satistfies the Encryptor interface
func (eFn EncryptorFn) Encrypt(payload []byte) (encrypted, key []byte, err error) {
	return eFn(payload)
}

// Encrypt encrypts a payload using AES , returning the encrypted payload with
// the key that was used to encrypt it.
func Encrypt(payload []byte) (encrypted, key []byte, err error) {
	k, err := generateNewKey()
	if err != nil {
		return nil, nil, errors.Wrap(err, "generating new key for storage")
	}

	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating new aes cipher")
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(c, iv[:])

	plain := bytes.NewBuffer(payload)

	var eBuff bytes.Buffer
	writer := &cipher.StreamWriter{S: stream, W: &eBuff}
	if _, err := io.Copy(writer, plain); err != nil {
		return nil, nil, errors.Wrap(err, "copying buffer")
	}

	return eBuff.Bytes(), k, nil
}
