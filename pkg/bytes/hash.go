package bytes

// hash is used as the key for the internal storage of the HashMap.
// Task note: just a string for the purposes of this exercise but something
// more robust would be required in a production system.
// Task note: A generated has must be completely unique
type hash string

func newHash(id []byte) hash {
	return hash(string(id))
}

