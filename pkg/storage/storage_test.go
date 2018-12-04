package storage_test

import (
	"github.com/glynternet/encryptedstore/pkg/client"
	"github.com/glynternet/encryptedstore/pkg/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ensure that Encrypted satisfies the task's required interface
var _ client.Client = &storage.Encrypted{}

func TestStoreRetrieve(t *testing.T) {
	var e storage.Encrypted

	id := []byte("qwert")
	payload := []byte("yuiop")

	key, err := e.Store(id, payload)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	t.Run("retreve with invalid key", func(t *testing.T) {
		invalidKey := append([]byte("invalid-"), key...)
		decrypted, err := e.Retrieve(id, invalidKey)
		assert.Error(t, err)
		assert.Nil(t, decrypted)
	})

	t.Run("retrieve with valid key", func(t *testing.T) {
		decrypted, err := e.Retrieve(id, key)
		assert.NoError(t, err)
		assert.Equal(t, payload,decrypted)
	})

	t.Run("retrieve with non-existant id", func(t *testing.T) {
		invalidID := append([]byte("invalid-"), id...)
		decrypted, err := e.Retrieve(invalidID, key)
		assert.Error(t, err)
		assert.Nil(t, decrypted)
	})
}