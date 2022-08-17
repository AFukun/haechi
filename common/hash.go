package common

import (
	"crypto/sha256"
)

const (
	HashSize = sha256.Size
)

type Hash [HashSize]byte

func NewHash(data []byte) Hash {
	return sha256.Sum256(data)
}
