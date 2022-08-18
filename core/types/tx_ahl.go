package types

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/AFukun/haechi/common"
)

const (
	AhlTxByteSize   = 22
	AhlTxEncodeSize = 32
)

type AhlTx struct {
	From      common.Address
	To        common.Address
	FromShard uint8
	ToStard   uint8
	Value     uint32
	Data      uint32
	Nonce     uint32
}

func (tx AhlTx) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, tx)

	return buf.Bytes()
}

func (tx AhlTx) HashString() string {
	hash := common.NewHash(tx.Bytes())
	return hash.ToHexString()
}

func (tx AhlTx) EncodeToBase64String() string {
	return base64.StdEncoding.EncodeToString(tx.Bytes())
}

func DecodeAhlTxBase64String(encodedString string) (AhlTx, error) {
	var tx AhlTx
	txBytes, err := base64.StdEncoding.DecodeString(encodedString)
	binary.Read(bytes.NewBuffer(txBytes), binary.LittleEndian, &tx)

	return tx, err
}
