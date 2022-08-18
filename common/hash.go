package common

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	HashSize = sha256.Size
)

type Hash [HashSize]byte

func (h Hash) ToHexString() string {
	return hex.EncodeToString(h[:])
}

func NewHash(data []byte) Hash {
	return sha256.Sum256(data)
}
