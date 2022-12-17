package abci

import (
	"encoding/binary"
	"fmt"
	"time"

	// "log"
	"strconv"
	// "time"

	ahlNode "github.com/AFukun/haechi/consensus/ahl/shard/validator"
	abcicode "github.com/tendermint/tendermint/abci/example/code"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

var (
	// stateKey        = []byte("stateKey")
	kvPairPrefixKey = []byte("kvPairKey:")

	ProtocolVersion uint64 = 0x1
)

func prefixKey(key []byte) []byte {
	return append(kvPairPrefixKey, key...)
}

var _ abcitypes.Application = (*AhlShardApplication)(nil)

type AhlShardApplication struct {
	abcitypes.BaseApplication
	// mu   sync.Mutex
	Node *ahlNode.ValidatorInterface
	// intraTxBatch *badger.Txn
}

func NewAhlShardApplication(node *ahlNode.ValidatorInterface) *AhlShardApplication {
	return &AhlShardApplication{
		Node: node,
	}
}

func (AhlShardApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (AhlShardApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *AhlShardApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return abcitypes.ResponseCheckTx{Code: abcicode.CodeTypeOK, GasWanted: 1}
}

func (app *AhlShardApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (app *AhlShardApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	_, tx_json := ahlNode.ResolveTx(req.Tx)
	var err1, err2 error
	var events []abcitypes.Event
	var event_type string
	new_tx := ahlNode.TransactionType{
		From_shard: tx_json.From_shard,
		To_shard:   tx_json.To_shard,
		Tx_type:    tx_json.Tx_type,
		From:       tx_json.From,
		To:         tx_json.To,
		Value:      tx_json.Value,
		Data:       tx_json.Data,
		Nonce:      tx_json.Nonce,
		TX_id:      tx_json.TX_id,
	}
	new_tx1 := ahlNode.TransactionType{
		From_shard: tx_json.From_shard,
		To_shard:   tx_json.To_shard,
		Tx_type:    tx_json.Tx_type,
		From:       tx_json.From,
		To:         tx_json.To,
		Value:      tx_json.Value,
		Data:       tx_json.Data,
		Nonce:      tx_json.Nonce,
		TX_id:      tx_json.TX_id,
	}
	if tx_json.Tx_type == ahlNode.IntraShard_TX {
		event_type = "intra-shard transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
	} else if tx_json.Tx_type == ahlNode.InterShard_TX_Verify {
		event_type = "inter-shard verification transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("lock"))
		new_tx.Tx_type = ahlNode.InterShard_TX_Execute
		new_tx1.Tx_type = ahlNode.InterShard_TX_Commit_sender
		_, exe_tx := ahlNode.Deserilization(new_tx)
		_, com_tx := ahlNode.Deserilization(new_tx1)
		if app.Node.Leader {
			go app.Node.DeliverExecutionTx(exe_tx, new_tx.To_shard)
			go app.Node.DeliverCommitTx(com_tx)
		}
	} else if tx_json.Tx_type == ahlNode.InterShard_TX_Execute {
		event_type = "inter-shard execution transaction"
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("lock"))
		new_tx.Tx_type = ahlNode.InterShard_TX_Commit_receiver
		_, com_tx := ahlNode.Deserilization(new_tx)
		if app.Node.Leader {
			go app.Node.DeliverCommitTx(com_tx)
		}
	} else if tx_json.Tx_type == ahlNode.InterShard_TX_Update_sender {
		event_type = "inter-shard update transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		// Trace: cross-shard tx confirmation latency
		if string(tx_json.From) == "CROS" {
			fmt.Println("cross-shard trace, nonce is", tx_json.Nonce)
			fmt.Println("cross-shard trace, end time is", time.Now())
		}
	} else if tx_json.Tx_type == ahlNode.InterShard_TX_Update_receiver {
		event_type = "inter-shard update transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
	}
	if err1 != nil || err2 != nil {
		panic(err1)
	}
	events = []abcitypes.Event{
		{
			Type: event_type,
			Attributes: []abcitypes.EventAttribute{
				{Key: "from", Value: string(tx_json.From), Index: true},
				{Key: "to", Value: string(tx_json.To), Index: true},
				{Key: "value", Value: strconv.Itoa(int(tx_json.Value)), Index: true},
				{Key: "data", Value: string(tx_json.Data), Index: true},
			},
		},
	}
	return abcitypes.ResponseDeliverTx{Code: abcicode.CodeTypeOK, Events: events}
}

func (app *AhlShardApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *AhlShardApplication) Commit() abcitypes.ResponseCommit {
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.Node.BCState.Size)
	app.Node.BCState.AppHash = appHash
	app.Node.BCState.Height++
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (app *AhlShardApplication) Query(reqQuery abcitypes.RequestQuery) (resQuery abcitypes.ResponseQuery) {
	if reqQuery.Prove {
		value, err := app.Node.BCState.Database.Get(prefixKey(reqQuery.Data))
		if err != nil {
			panic(err)
		}
		if value == nil {
			resQuery.Log = "does not exist"
		} else {
			resQuery.Log = "exists"
		}
		resQuery.Index = -1 // TODO make Proof return index
		resQuery.Key = reqQuery.Data
		resQuery.Value = value
		resQuery.Height = int64(app.Node.BCState.Height)

		return
	}

	resQuery.Key = reqQuery.Data
	value, err := app.Node.BCState.Database.Get(prefixKey(reqQuery.Data))
	if err != nil {
		panic(err)
	}
	if value == nil {
		resQuery.Log = "does not exist"
	} else {
		resQuery.Log = "exists"
	}
	resQuery.Value = value
	resQuery.Height = int64(app.Node.BCState.Height)

	return resQuery
}
