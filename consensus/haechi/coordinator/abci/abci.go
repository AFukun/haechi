package abci

import (
	"encoding/binary"
	"fmt"
	"time"

	// "log"
	// "time"

	haechiNode "github.com/AFukun/haechi/consensus/haechi/coordinator/validator"
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

var _ abcitypes.Application = (*HaechiBeaconApplication)(nil)

type HaechiBeaconApplication struct {
	abcitypes.BaseApplication
	// mu   sync.Mutex
	Node *haechiNode.ValidatorInterface
	// intraTxBatch *badger.Txn
}

func NewHaechiBeaconApplication(node *haechiNode.ValidatorInterface) *HaechiBeaconApplication {
	return &HaechiBeaconApplication{
		Node: node,
	}
}

func (HaechiBeaconApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (HaechiBeaconApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

// receive: blocktimestamp=111000,fromid=1,toid=1,type=0,from=ABCD1,to=DCBA1,value=0,data=NONE,nonce=0,blockheight=1000,index=0>fromid=1,toid=2,type=0,from=ABCD1,to=DCBA2,value=0,data=NONE,nonce=1,blockheight=1000,index=1>
func (app *HaechiBeaconApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	app.Node.UpdateShardCrosslinkMsgs(req.Tx)
	app.Node.UpdateOrderParameters(req.Tx)
	return abcitypes.ResponseCheckTx{Code: abcicode.CodeTypeOK, GasWanted: 1}
}

func (app *HaechiBeaconApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	if app.Node.BCState.Height == 0 {
		app.Node.BCState.Height++
		return abcitypes.ResponseBeginBlock{}
	}
	return abcitypes.ResponseBeginBlock{}
}

func (app *HaechiBeaconApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	if app.Node.StartOrder() && app.Node.Start_Order {
		// MicroBench: test time difference of shard's CrossLinks
		temp_output := fmt.Sprintf("current block height is %v, receive CrossLink completely, at time %v", app.Node.BCState.Height, time.Now())
		fmt.Println(temp_output)
		app.Node.Start_Order = false
		// TODO: how to securely concurrently diliver call lists?
		if app.Node.Leader {
			go app.Node.DeliverCallLists()
		}

	}
	var events []abcitypes.Event
	// var event_type string
	return abcitypes.ResponseDeliverTx{Code: abcicode.CodeTypeOK, Events: events}
}

func (app *HaechiBeaconApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *HaechiBeaconApplication) Commit() abcitypes.ResponseCommit {
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.Node.BCState.Size)
	app.Node.BCState.AppHash = appHash
	app.Node.BCState.Height++
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (app *HaechiBeaconApplication) Query(reqQuery abcitypes.RequestQuery) (resQuery abcitypes.ResponseQuery) {
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
