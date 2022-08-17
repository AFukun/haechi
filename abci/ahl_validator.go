package abci

import (
	"github.com/AFukun/haechi/consensus"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type AhlValidatorApplication struct {
	engine *consensus.AhlValidatorEngine
}

var _ abcitypes.Application = (*AhlValidatorApplication)(nil)

func NewAhlValidatorApplication(coordinatorIP string) *AhlValidatorApplication {
	return &AhlValidatorApplication{consensus.NewAhlValidatorEngine(coordinatorIP)}
}

func (AhlValidatorApplication) Info(abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *AhlValidatorApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	code := app.engine.Validate(string(req.Tx))
	return abcitypes.ResponseDeliverTx{Code: code}
}

func (app *AhlValidatorApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	code := app.engine.Validate(string(req.Tx))
	return abcitypes.ResponseCheckTx{Code: code}
}

func (app *AhlValidatorApplication) Commit() abcitypes.ResponseCommit {
	app.engine.Excute()
	app.engine.Communicate()
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (AhlValidatorApplication) Query(abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{}
}

func (AhlValidatorApplication) InitChain(abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (AhlValidatorApplication) BeginBlock(abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (AhlValidatorApplication) EndBlock(abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (AhlValidatorApplication) ListSnapshots(abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (AhlValidatorApplication) OfferSnapshot(abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (AhlValidatorApplication) LoadSnapshotChunk(abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (AhlValidatorApplication) ApplySnapshotChunk(abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}
