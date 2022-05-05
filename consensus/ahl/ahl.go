//-----------------------------------------------------------------------------
// the implement of AHL cross-shard consensus
// where there is an extra shard, AHLCoordinator, for coordination

package ahl

import (
	htx "haechi/structure/transaction"
	"strconv"
)

type AHLCoordinator struct {
	shardIPs map[int]string
	listAdd  string // TCP socket address receiving cross-shard transactions
}

func (ac *AHLCoordinator) DeliverTxToShards(tx []byte) {
	hTx, _ := htx.ShardTxToHtransaction(tx)
	vtx, etx := htx.DivideTransaction(hTx)
	vtxByte := htx.VerifyTxToShardTx(vtx)
	etxByte := htx.ExecutionTxToShardTx(etx)
	fromShardId, _ := strconv.Atoi(string(vtx.FromShardId))
	toShardId, _ := strconv.Atoi(string(etx.ToShardId))
	ac.ForwardTx(fromShardId, vtxByte)
	ac.ForwardTx(toShardId, etxByte)
}

// TODO: forward a transaction to a sender's (or contract's) shard
func (ac *AHLCoordinator) ForwardTx(shardId int, tx []byte) {
	// shardIP := ac.shardIPs[shardId]
}
