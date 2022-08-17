package consensus

import (
	"github.com/AFukun/haechi/common"
	"github.com/AFukun/haechi/core/db"
	"github.com/AFukun/haechi/core/types"
	"github.com/AFukun/haechi/log"
	"github.com/AFukun/haechi/tools"
)

type AhlValidatorEngine struct {
	db              *db.SimpleKVDatabase
	localBatch      []types.AhlTx
	intershardBatch []types.AhlTx
	pendingTxHash   map[common.Hash]struct{}
	coordinatorIP   string
}

func NewAhlValidatorEngine(coordinatorIP string) *AhlValidatorEngine {
	return &AhlValidatorEngine{
		db:              db.NewSimpleKVDatabase(),
		localBatch:      []types.AhlTx{},
		intershardBatch: []types.AhlTx{},
		pendingTxHash:   make(map[common.Hash]struct{}),
		coordinatorIP:   coordinatorIP,
	}
}

// Validation Code
// 0: validated
// 1: error
func (e *AhlValidatorEngine) Validate(txs string) uint32 {
	tx, err := types.DecodeAhlTxString(txs)
	if err != nil {
		return 1
	}
	if tx.FromShard != tx.ToStard {
		hash := common.NewHash(tx.Bytes())
		_, exist := e.pendingTxHash[hash]
		if exist {
			e.localBatch = append(e.localBatch, tx)
			delete(e.pendingTxHash, hash)
		} else {
			e.pendingTxHash[hash] = struct{}{}
			e.intershardBatch = append(e.intershardBatch, tx)
		}
	} else {
		e.localBatch = append(e.localBatch, tx)
	}

	return 0
}

func (e *AhlValidatorEngine) Excute() {
	for _, tx := range e.localBatch {
		e.db.Put(tx.Value, tx.Data)
	}
	e.localBatch = []types.AhlTx{}
}

func (e *AhlValidatorEngine) Communicate() {
	if e.coordinatorIP != "" {
		for _, tx := range e.intershardBatch {
			_, err := tools.SendTxString(e.coordinatorIP, tx.EncodeToString())
			if err != nil {
				log.Error("failed to send cross-shard tx", "ip", e.coordinatorIP, "err", err)
			}
		}
	}
}
