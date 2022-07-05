package abci

import (
	"encoding/binary"
	"fmt"

	elrondNode "github.com/AFukun/haechi/consensus/elrond/coordinator/validator"
	abcicode "github.com/tendermint/tendermint/abci/example/code"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	// "github.com/dgraph-io/badger/v3"
)

var (
	// stateKey        = []byte("stateKey")
	kvPairPrefixKey = []byte("kvPairKey:")

	ProtocolVersion uint64 = 0x1
)

func prefixKey(key []byte) []byte {
	return append(kvPairPrefixKey, key...)
}

var _ abcitypes.Application = (*ElrondApplication)(nil)

type ElrondApplication struct {
	abcitypes.BaseApplication
	// mu   sync.Mutex
	Node *elrondNode.ValidatorInterface
	// intraTxBatch *badger.Txn
}

func NewElrondApplication(node *elrondNode.ValidatorInterface) *ElrondApplication {
	return &ElrondApplication{
		Node: node,
	}
}

func (ElrondApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

// func (app *ElrondApplication) Info(_ context.Context, req *abcitypes.RequestInfo) (*abcitypes.ResponseInfo, error) {
// 	// app.mu.Lock()
// 	// defer app.mu.Unlock()
// 	return &abcitypes.ResponseInfo{
// 		Data:             fmt.Sprintf("{\"size\":%v}", app.Node.BCState.Size),
// 		Version:          "v1",
// 		AppVersion:       ProtocolVersion,
// 		LastBlockHeight:  app.Node.BCState.Height,
// 		LastBlockAppHash: app.Node.BCState.AppHash,
// 	}, nil
// }
func (ElrondApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *ElrondApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	fmt.Println("Elrond: Intra_shard CheckTx...")
	// code, _ := elrondNode.ResolveTx(req.Tx)
	return abcitypes.ResponseCheckTx{Code: abcicode.CodeTypeOK, GasWanted: 1}
}

func (app *ElrondApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	fmt.Println("Elrond: Intra_shard BeginBlock...")
	// app.intraTxBatch = app.Node.BCState.Database.NewTransaction(true)
	return abcitypes.ResponseBeginBlock{}
}

func (app *ElrondApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	_, tx_json := elrondNode.ResolveTx(req.Tx)
	var err1, err2 error
	events := []abcitypes.Event{
		{
			Type: "app",
			Attributes: []abcitypes.EventAttribute{
				{Key: "creator", Value: "Cosmoshi Netowoko", Index: true},
				{Key: "key", Value: string(tx_json.From), Index: true},
				{Key: "index_key", Value: "index is working", Index: true},
				{Key: "noindex_key", Value: "index is working", Index: false},
			},
		},
	}
	new_tx := elrondNode.TransactionType{
		Tx_type: tx_json.Tx_type,
		From:    tx_json.From,
		To:      tx_json.To,
		Value:   tx_json.Value,
		Data:    tx_json.Data,
	}
	if tx_json.Tx_type == elrondNode.IntraShard_TX {
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Verify {
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("lock"))
		new_tx.Tx_type = elrondNode.InterShard_TX_Execute
		_, exec_tx := elrondNode.Deserilization(new_tx)
		if app.Node.Leader {
			app.Node.DeliverExecutionTx(exec_tx)
		}
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Execute {
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("lock"))
		new_tx.Tx_type = elrondNode.InterShard_TX_Commit
		_, commit_tx := elrondNode.Deserilization(new_tx)
		if app.Node.Leader {
			app.Node.DeliverCommitTx(commit_tx)
		}
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Commit {
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		new_tx.Tx_type = elrondNode.InterShard_TX_Update
		_, update_tx := elrondNode.Deserilization(new_tx)
		if app.Node.Leader {
			app.Node.DeliverUpdateTx(update_tx)
		}
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Update {
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
		app.Node.BCState.Size++
		if err1 == nil || err2 == nil {
			return abcitypes.ResponseDeliverTx{Code: abcicode.CodeTypeOK, Events: events}
		}

	}
	if err1 == nil || err2 == nil {
		return abcitypes.ResponseDeliverTx{Code: abcicode.CodeTypeOK, Events: events}
	}
	return abcitypes.ResponseDeliverTx{Code: abcicode.CodeTypeUnknownError, Events: events}
}

func (app *ElrondApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	fmt.Println("Elrond: Intra_shard EndBlock...")
	return abcitypes.ResponseEndBlock{}
}

func (app *ElrondApplication) Commit() abcitypes.ResponseCommit {
	fmt.Println("Elrond: Intra_shard Commit...")
	// if err := app.intraTxBatch.Commit(); err != nil {
	// 	log.Panicf("Error writing to database, unable to commit block: %v", err)
	// }
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.Node.BCState.Size)
	app.Node.BCState.AppHash = appHash
	app.Node.BCState.Height++
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (app *ElrondApplication) Query(reqQuery abcitypes.RequestQuery) (resQuery abcitypes.ResponseQuery) {
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
		resQuery.Height = app.Node.BCState.Height

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
	resQuery.Height = app.Node.BCState.Height

	return resQuery
}
