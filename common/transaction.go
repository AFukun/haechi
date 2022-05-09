//-----------------------------------------------------------------------------
// Four types of transactions in Haechi: {ntx, htx, vtx, etx}
// ntx: normal transaction used for bft consensus of tendermint
// htx: haechi transaction
// vtx: verification transaction that is sent to sender's shard for account verification
// etx: execution transaction that is sent to contract's shard for contract execution

package common

import (
	"bytes"
	"log"
)

type HaechiTransaction struct {
	FromShardId []byte
	FromAdd     []byte // is equal to "key" in ShardBFTApplication
	FromValue   []byte // is equal to "value" in ShardBFTApplication
	ToShardId   []byte
	ToAdd       []byte // is equal to "key" in ShardBFTApplication
	ToValue     []byte // is equal to "value" in ShardBFTApplication
}

type VerifyTransaction struct {
	FromShardId []byte
	FromAdd     []byte // is equal to "key" in ShardBFTApplication
	FromValue   []byte // is equal to "value" in ShardBFTApplication
}

type ExecutionTransaction struct {
	ToShardId []byte
	ToAdd     []byte // is equal to "key" in ShardBFTApplication
	ToValue   []byte // is equal to "value" in ShardBFTApplication
}

// e.g., a tx in Haechi has a form like "type=htx,fromShardId=1,toShardId=2,fromAdd=aa,toAdd=bb,value=0,parameters=2"
func ShardTxToHtransaction(ctx []byte) (htx HaechiTransaction, isValid bool) {
	txElements := bytes.Split(ctx, []byte(","))
	for _, txElement := range txElements {
		kv := bytes.Split(txElement, []byte("="))
		if string(kv[0]) == "FromShardId" {
			copy(htx.FromShardId, kv[1])
		} else if string(kv[0]) == "FromAdd" {
			copy(htx.FromAdd, kv[1])
		} else if string(kv[0]) == "FromValue" {
			copy(htx.FromValue, kv[1])
		} else if string(kv[0]) == "ToShardId" {
			copy(htx.ToShardId, kv[1])
		} else if string(kv[0]) == "ToShardId" {
			copy(htx.ToShardId, kv[1])
		} else if string(kv[0]) == "ToAdd" {
			copy(htx.ToAdd, kv[1])
		} else if string(kv[0]) == "ToValue" {
			copy(htx.ToValue, kv[1])
		} else {
			log.Panicf("invalid format of transaction")
			return htx, false
		}
	}
	return htx, true
}

func DivideTransaction(htx HaechiTransaction) (vTx VerifyTransaction, eTx ExecutionTransaction) {
	copy(vTx.FromShardId, htx.FromShardId)
	copy(vTx.FromAdd, htx.FromAdd)
	copy(vTx.FromValue, htx.FromValue)
	copy(eTx.ToShardId, htx.ToShardId)
	copy(eTx.ToAdd, htx.ToAdd)
	copy(eTx.ToValue, htx.ToValue)
	return vTx, eTx
}

func VerifyTxToShardTx(vtx VerifyTransaction) (shardTx []byte) {
	tempStr := "type=vtx,fromShardId=" + string(vtx.FromShardId) + ",toShardId=0,fromAdd=" + string(vtx.FromAdd) + ",toAdd=0,value=" + string(vtx.FromValue) + ",parameters=0"
	shardTx = []byte(tempStr)
	return shardTx
}

func ExecutionTxToShardTx(etx ExecutionTransaction) (shardTx []byte) {
	tempStr := "type=etx,fromShardId=0,toShardId=" + string(etx.ToShardId) + ",fromAdd=aa,toAdd=" + string(etx.ToAdd) + ",value=0,parameters=" + string(etx.ToValue)
	shardTx = []byte(tempStr)
	return shardTx
}
