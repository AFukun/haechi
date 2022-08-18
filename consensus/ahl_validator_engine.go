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
	height          uint
	localBatch      []types.AhlTx
	intershardBatch []types.AhlTx
	pendingTxHash   map[common.Hash]struct{}
	coordinatorIP   string
	logger          log.Logger
}

func NewAhlValidatorEngine(coordinatorIP string) *AhlValidatorEngine {
	return &AhlValidatorEngine{
		db:              db.NewSimpleKVDatabase(),
		height:          0,
		localBatch:      []types.AhlTx{},
		intershardBatch: []types.AhlTx{},
		pendingTxHash:   make(map[common.Hash]struct{}),
		coordinatorIP:   coordinatorIP,
		logger:          log.New("ahl"),
	}
}

func (e *AhlValidatorEngine) BeginBlock() {
	e.localBatch = []types.AhlTx{}
	e.intershardBatch = []types.AhlTx{}
	e.height++
	e.logger.Info("init block", "height", e.height)
}

func (e *AhlValidatorEngine) EndBlock() {
	e.logger.Info("end block",
		"height", e.height,
		"num_local_tx", len(e.localBatch),
		"num_cross_shard_tx", len(e.intershardBatch),
		"num_pending_tx", len(e.pendingTxHash))
}

// Validation Code
// 0: valid
// 1: invalid
func (e *AhlValidatorEngine) Validate(txs string) uint32 {
	tx, err := types.DecodeAhlTxBase64String(txs)
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
	if len(e.localBatch) != 0 {
		txHashList := make([]string, 0)
		for _, tx := range e.localBatch {
			e.db.Put(tx.Value, tx.Data)
			txHashList = append(txHashList, tx.HashString())
		}
		e.logger.Info("excuted local tx",
			"num_tx", len(e.localBatch),
			"tx_list", txHashList)
	}
	if len(e.pendingTxHash) != 0 {
		txHashList := make([]string, 0)
		for tx := range e.pendingTxHash {
			txHashList = append(txHashList, tx.ToHexString())
		}
		e.logger.Info("pending tx",
			"num_tx", len(e.pendingTxHash),
			"tx_list", txHashList)
	}
}

func (e *AhlValidatorEngine) Communicate() {
	if e.coordinatorIP != "" {
		txHashList := make([]string, 0)
		for _, tx := range e.intershardBatch {
			_, err := tools.SendTxString(e.coordinatorIP, tx.EncodeToBase64String())
			if err != nil {
				e.logger.Error("error when sending tx",
					"ip", e.coordinatorIP,
					"err", err,
					"tx", tx.HashString())
			}
			txHashList = append(txHashList, tx.HashString())
		}
		e.logger.Info("sent cross shard tx",
			"num_success", len(txHashList),
			"num_failed", len(e.intershardBatch)-len(txHashList),
			"tx_list", txHashList)
	}
}
