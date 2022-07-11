package abci

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	haechiNode "github.com/AFukun/haechi/consensus/haechi/shard/validator"
	haechitypes "github.com/AFukun/haechi/core/types"
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

var _ abcitypes.Application = (*HaechiShardApplication)(nil)

type HaechiShardApplication struct {
	abcitypes.BaseApplication
	// mu   sync.Mutex
	Node *haechiNode.ValidatorInterface
	// intraTxBatch *badger.Txn
}

func NewHaechiShardApplication(node *haechiNode.ValidatorInterface) *HaechiShardApplication {
	return &HaechiShardApplication{
		Node: node,
	}
}

func (HaechiShardApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (HaechiShardApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *HaechiShardApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	fmt.Println("Haechi shard: Intra_shard CheckTx...")
	// code, _ := haechiNode.ResolveTx(req.Tx)
	return abcitypes.ResponseCheckTx{Code: abcicode.CodeTypeOK, GasWanted: 1}
}

func (app *HaechiShardApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	fmt.Println("Haechi shard: Intra_shard BeginBlock...")
	// app.intraTxBatch = app.Node.BCState.Database.NewTransaction(true)
	return abcitypes.ResponseBeginBlock{}
}

func (app *HaechiShardApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	fmt.Println("Haechi shard: Intra_shard DeliverTx...")
	_, tx_json := haechiNode.ResolveTx(req.Tx)
	var err1, err2 error
	var events []abcitypes.Event
	var event_type string
	if tx_json.Tx_type == haechiNode.IntraShard_TX {
		// fmt.Println("This is an intra-shard transaction")
		event_type = "intra-shard transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("0"))
		err2 = app.Node.BCState.Database.Set(prefixKey(tx_json.To), []byte("0"))
		app.Node.BCState.Size++
	} else if tx_json.Tx_type == haechiNode.InterShard_TX_Verify {
		app.Node.BCState.Index++
		event_type = "inter-shard transaction"
		err1 = app.Node.BCState.Database.Set(prefixKey(tx_json.From), []byte("lock"))
		app.Node.Current_cl = string(req.Tx)
		app.Node.Current_cl += ",blockheight="
		app.Node.Current_cl += strconv.Itoa(int(app.Node.BCState.Height))
		app.Node.Current_cl += ",index="
		app.Node.Current_cl += strconv.Itoa(int(app.Node.BCState.Index))
		app.Node.Current_cl += ";"
	} else if tx_json.Tx_type == haechiNode.CrossShard_Call_List {
		// _, ccl_txs := haechiNode.CclToTxs(req.Tx)
		_, ccl_txs := haechitypes.TxToCCL(req.Tx, app.Node.Shard_id)
		for {
			ccl_tx, _ := ccl_txs.Call_txs.Dequeue()
			q_tx_json := ccl_tx.(haechiNode.TransactionType)
			err1 = app.Node.BCState.Database.Set(prefixKey(q_tx_json.From), []byte("0"))
			err2 = app.Node.BCState.Database.Set(prefixKey(q_tx_json.To), []byte("0"))
			if ccl_txs.Call_txs.Empty() || err1 != nil {
				break
			}
			app.Node.BCState.Size++
		}
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

func (app *HaechiShardApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	fmt.Println("Haechi shard: Intra_shard EndBlock...")
	return abcitypes.ResponseEndBlock{}
}

func (app *HaechiShardApplication) Commit() abcitypes.ResponseCommit {
	fmt.Println("Haechi shard: Intra_shard Commit...")
	current_timestamp := time.Now().Unix()
	if app.Node.Leader {
		go app.Node.DeliverCrossLink(current_timestamp)
	}
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, int64(app.Node.BCState.Size))
	app.Node.BCState.AppHash = appHash
	app.Node.BCState.Height++
	app.Node.BCState.Index = 0
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (app *HaechiShardApplication) Query(reqQuery abcitypes.RequestQuery) (resQuery abcitypes.ResponseQuery) {
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