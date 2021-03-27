package xredis

import (
	"crypto/sha1"
	"encoding/hex"
)

type (
	// A Script is a set of custom atomic actions to perform in the scache.
	Script interface {
		// Source returns the script data.
		Source() string
		// Hash returns the SHA1 of the script.
		Hash() string
	}

	script struct {
		src, hash string
	}
)

// NewScript returns a new Script.
func NewScript(src string) Script {
	sum := sha1.Sum([]byte(src))
	return &script{
		src:  src,
		hash: hex.EncodeToString(sum[:]),
	}
}
func (s *script) Source() string {
	return s.src
}

func (s *script) Hash() string {
	return s.hash
}
