package abci

import (
	"encoding/binary"
	"log"
	"strconv"

	elrondNode "github.com/AFukun/haechi/consensus/elrond/coordinator/validator"
	haechitypes "github.com/AFukun/haechi/core/types"
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

func (ElrondApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *ElrondApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return abcitypes.ResponseCheckTx{Code: abcicode.CodeTypeOK, GasWanted: 1}
}

func (app *ElrondApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (app *ElrondApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	_, tx_json := elrondNode.ResolveTx(req.Tx)
	var err1, err2 error
	var events []abcitypes.Event
	var event_type string
	new_tx := haechitypes.TransactionType{
		From_shard: tx_json.From_shard,
		To_shard:   tx_json.To_shard,
		Tx_type:    tx_json.Tx_type,
		From:       tx_json.From,
		To:         tx_json.To,
		Value:      tx_json.Value,
		Data:       tx_json.Data,
	}
	if tx_json.Tx_type == elrondNode.IntraShard_TX {
		log.Println("This is an intra-shard transaction")
		event_type = "intra-shard transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
		app.Node.BCState.Size++
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Verify {
		log.Println("This is a verification transaction")
		event_type = "inter-shard verification transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("lock"))
		new_tx.Tx_type = elrondNode.InterShard_TX_Execute
		_, exec_tx := elrondNode.Deserilization(new_tx)
		if app.Node.Leader {
			go app.Node.DeliverExecutionTx(exec_tx, new_tx.To_shard)
		}

	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Execute {
		log.Println("This is an execution transaction")
		event_type = "inter-shard execution transaction"
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("lock"))
		new_tx.Tx_type = elrondNode.InterShard_TX_Commit
		_, commit_tx := elrondNode.Deserilization(new_tx)
		if app.Node.Leader {
			go app.Node.DeliverCommitTx(commit_tx, new_tx.From_shard)
		}
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Commit {
		log.Println("This is a commit transaction")
		event_type = "inter-shard commit transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		new_tx.Tx_type = elrondNode.InterShard_TX_Update
		_, update_tx := elrondNode.Deserilization(new_tx)
		if app.Node.Leader {
			go app.Node.DeliverUpdateTx(update_tx, new_tx.To_shard)
		}
	} else if tx_json.Tx_type == elrondNode.InterShard_TX_Update {
		log.Println("This is an update transaction")
		event_type = "inter-shard update transaction"
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
		app.Node.BCState.Size++
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

func (app *ElrondApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *ElrondApplication) Commit() abcitypes.ResponseCommit {
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
