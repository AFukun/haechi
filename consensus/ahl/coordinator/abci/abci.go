package abci

import (
	"encoding/binary"

	// "log"
	"strconv"

	ahlNode "github.com/AFukun/haechi/consensus/ahl/coordinator/validator"
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

var _ abcitypes.Application = (*AhlBeaconApplication)(nil)

type AhlBeaconApplication struct {
	abcitypes.BaseApplication
	// mu   sync.Mutex
	Node *ahlNode.ValidatorInterface
	// intraTxBatch *badger.Txn
}

func NewAhlBeaconApplication(node *ahlNode.ValidatorInterface) *AhlBeaconApplication {
	return &AhlBeaconApplication{
		Node: node,
	}
}

func (AhlBeaconApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (AhlBeaconApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *AhlBeaconApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	// tln("beacon receive tx " + string(req.Tx))
	// tln("beacon receive tx, current time is: " + time.Now().String())
	return abcitypes.ResponseCheckTx{Code: abcicode.CodeTypeOK, GasWanted: 1}
}

func (app *AhlBeaconApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (app *AhlBeaconApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	_, tx_json := ahlNode.ResolveTx(req.Tx)
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
	if tx_json.Tx_type == ahlNode.InterShard_TX_Verify {
		// // tln("This is a verification transaction")
		event_type = "inter-shard verification transaction"
		_, cs_tx := ahlNode.Deserilization(new_tx)
		if app.Node.Leader {
			go app.Node.DeliverCrossShardTx(cs_tx, new_tx.From_shard)
		}
	} else if tx_json.Tx_type == ahlNode.InterShard_TX_Commit_sender {
		// // tln("This is a commit transaction from sender")
		event_type = "sender commit verification transaction"
		app.Node.Tx_set[tx_json.TX_id%ahlNode.Process_Length] += 1
	} else if tx_json.Tx_type == ahlNode.InterShard_TX_Commit_receiver {
		// // tln("This is a commit transaction from receiver")
		event_type = "receiver commit verification transaction"
		app.Node.Tx_set[tx_json.TX_id%ahlNode.Process_Length] += 1
	}

	if app.Node.Tx_set[tx_json.TX_id%ahlNode.Process_Length] >= 2 {
		// // tln("Beacon chain commits a transaction")
		app.Node.Tx_set[tx_json.TX_id%ahlNode.Process_Length] = 0
		new_tx.Tx_type = ahlNode.InterShard_TX_Update_sender
		new_tx1.Tx_type = ahlNode.InterShard_TX_Update_receiver
		_, update_sender_tx := ahlNode.Deserilization(new_tx)
		_, update_receiver_tx := ahlNode.Deserilization(new_tx1)
		if app.Node.Leader {
			go app.Node.DeliverCommitTx(update_sender_tx, new_tx.From_shard)
			go app.Node.DeliverCommitTx(update_receiver_tx, new_tx.To_shard)
		}
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

func (app *AhlBeaconApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *AhlBeaconApplication) Commit() abcitypes.ResponseCommit {
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.Node.BCState.Size)
	app.Node.BCState.AppHash = appHash
	app.Node.BCState.Height++
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (app *AhlBeaconApplication) Query(reqQuery abcitypes.RequestQuery) (resQuery abcitypes.ResponseQuery) {
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
