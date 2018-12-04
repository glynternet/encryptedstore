package storage

import (
	"fmt"
	"github.com/glynternet/encryptedstore/pkg/bytes"
	"github.com/glynternet/encryptedstore/pkg/crypt"
	"github.com/pkg/errors"
)


// Encrypted is a storage mechanism that will encrypt your payloads for storage
type Encrypted struct {
	encryptedStore bytes.HashMap
	encryptor crypt.Encryptor
	decryptor crypt.Decryptor
}

func (s *Encrypted) Store(id, payload []byte) (aesKey []byte, err error) {
	if s.encryptor == nil {
		s.encryptor = crypt.EncryptorFn(crypt.Encrypt)
	}
	e, k, err := crypt.Encrypt(payload)
	if err != nil {
		return nil, errors.Wrap(err, "encrypting payload")
	}
	s.encryptedStore.Store(id, e)
	return k, nil
}

func (s *Encrypted) Retrieve(id, key []byte) ([]byte, error) {
	if s.decryptor == nil {
		s.decryptor = crypt.DecryptorFn(crypt.Decrypt)
	}
	encrypted, ok := s.encryptedStore.Retrieve(id)
	if !ok {
		return nil, fmt.Errorf("no item with id: %s", string(id))
	}
	bs, err := crypt.Decrypt(key, encrypted)
	return bs, errors.Wrap(err, "decrypting payload")
}
