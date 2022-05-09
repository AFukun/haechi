//-----------------------------------------------------------------------------
// crosslink is the cross-shard message used in Haechi

package common

import (
	"bytes"
	"log"
)

// set the length of address as 160 bits, as Ethereum does,
// current length = 16 bits
const AddressLengthByte uint = 2

type CrossLink struct {
	shardId    []byte
	fromAdd    []byte
	toAdd      []byte
	parameters []byte
}

// transfer a cross-shard transaction into a crosslink
// e.g., a tx in Haechi has a form like "type=htx,fromShardId=1,toShardId=2,fromAdd=aa,toAdd=bb,value=0,parameters=2"
func TxToCrossLink(ctx []byte) (cl CrossLink, isValid bool) {
	txElements := bytes.Split(ctx, []byte(","))
	for _, txElement := range txElements {
		kv := bytes.Split(txElement, []byte("="))
		if string(kv[0]) == "toShardId" {
			copy(cl.shardId, kv[1])
		} else if string(kv[0]) == "fromAdd" {
			copy(cl.fromAdd, kv[1])
		} else if string(kv[0]) == "toAdd" {
			copy(cl.toAdd, kv[1])
		} else if string(kv[0]) == "parameters" {
			copy(cl.parameters, kv[1])
		} else {
			log.Panicf("invalid format of transaction")
			return cl, false
		}
	}
	return cl, true
}

// TODO: convert a crosslink to a normal transaction
func CrossLinkToTx(cl CrossLink) []byte {
	// TODO
	return []byte("test")
}
