package bytes

import "sync"

// HashMap is used for storage and retrieval of []byte payloads, keyed by a []byte key
// Task note: We could use functional options and a constructor to allow the user to
// provide their own hashing function.
type HashMap struct {
	hashFn func([]byte) hash
	m      map[hash][]byte
	lock   sync.RWMutex
}


// Store will insert or update any payload
// Task note: We could return more info here to say whether it existed previously
func (lm *HashMap) Store(key, payload []byte) {
	lm.lock.Lock()
	defer lm.lock.Unlock()
	if lm.m == nil {
		lm.m = make(map[hash][]byte)
	}
	if lm.hashFn == nil {
		lm.hashFn = newHash
	}
	lm.m[lm.hashFn(key)] = payload
}

// Retrieve fetches the given payload for the key and returns it, if it exists.
// The bool returned will be true if the payload was previously stored.
func (lm *HashMap) Retrieve(key []byte) ([]byte, bool) {
	lm.lock.RLock()
	defer lm.lock.RUnlock()
	if lm.m == nil {
		return nil, false
	}
	bs, ok := lm.m[lm.hashFn(key)]
	return bs, ok
}
