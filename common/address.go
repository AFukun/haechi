package common

import (
	"encoding/hex"
	"errors"
)

const (
	AddressByteSize = 4
)

type Address [AddressByteSize]byte

func (addr Address) ToHexString() string {
	return hex.EncodeToString(addr[:])
}

func HexStringToAddress(hexString string) (Address, error) {
	var addr [AddressByteSize]byte

	if len(hexString) != AddressByteSize*2 {
		err := errors.New("invalid string length")
		return addr, err
	}
	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return addr, err

	}
	copy(addr[:], decodedBytes[0:AddressByteSize])

	return addr, nil
}
