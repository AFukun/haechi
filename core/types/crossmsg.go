package types

import (
	"bytes"
	"strconv"

	aq "github.com/emirpasic/gods/queues/arrayqueue"
)

const (
	Addr_Length uint8 = 4
	Data_Length uint8 = 4
)

type CrossLink struct {
	Block_timestamp int64
	Shard_id        uint8
	Tx_type         uint8
	From            []byte
	To              []byte
	Value           uint32
	Data            []byte
	Nonce           uint32 // TODO: enable contineous tx requests by setting vary nonce
	Block_height    uint32
	Index           uint32
}

// blocktimestamp=111000,fromid=1,shardid=1,type=0,from=ABCD1,to=DCBA1,value=0,data=NONE,nonce=0,blockheight=1000,index=0;fromid=1,shardid=2,type=0,from=ABCD1,to=DCBA2,value=0,data=NONE,nonce=1,blockheight=1000,index=1;
func RequestToCrossLinks(ctx []byte) (uint32, []CrossLink) {
	var cls []CrossLink

	reqs := bytes.Split(ctx, []byte(";"))
	cls = make([]CrossLink, len(reqs))
	for _, req := range reqs {
		txElements := bytes.Split(req, []byte(","))
		var cl CrossLink
		for _, txElement := range txElements {
			kv := bytes.Split(txElement, []byte("="))
			switch {
			case string(kv[0]) == "blocktimestamp":
				temp_type64, _ := strconv.ParseInt(string(kv[1]), 10, 64)
				cl.Block_timestamp = temp_type64
			case string(kv[0]) == "shardid":
				temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				cl.Shard_id = uint8(temp_type64)
			case string(kv[0]) == "type":
				temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				cl.Tx_type = uint8(temp_type64)
			case string(kv[0]) == "from":
				cl.From = make([]byte, Addr_Length)
				copy(cl.From, kv[1])
			case string(kv[0]) == "to":
				cl.To = make([]byte, Addr_Length)
				copy(cl.To, kv[1])
			case string(kv[0]) == "value":
				// temp_value := string(kv[1])
				temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				cl.Value = uint32(temp_value64)
			case string(kv[0]) == "data":
				cl.Data = make([]byte, Data_Length)
				copy(cl.Data, kv[1])
			case string(kv[0]) == "nonce":
				temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				cl.Nonce = uint32(temp_value64)
			case string(kv[0]) == "blockheight":
				temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				cl.Block_height = uint32(temp_value64)
			case string(kv[0]) == "index":
				temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				cl.Index = uint32(temp_value64)
			default:
				return 1, cls
			}
		}
		cls = append(cls, cl)
	}
	return 0, cls
}

// TODO: convert a crosslink to a cross_shard call
func CrossLinkToTx(cl CrossLink) TransactionType {
	// TODO
	var ccl TransactionType
	ccl.Tx_type = cl.Tx_type
	ccl.From = cl.From
	ccl.To = cl.To
	ccl.Value = cl.Value
	ccl.Data = cl.Data
	ccl.Nonce = cl.Nonce
	return ccl
}

type TransactionType struct {
	Tx_type uint8
	From    []byte
	To      []byte
	Value   uint32
	Data    []byte
	Nonce   uint32 // TODO: enable contineous tx requests by setting vary nonce
}

func IsCrossShardTx(tx []byte) bool {
	txElements := bytes.Split(tx, []byte(","))
	for _, txElement := range txElements {
		kv := bytes.Split(txElement, []byte("="))
		if string(kv[0]) == "type" {
			temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			if temp_type64 == 1 {
				return true
			}
		}
	}
	return false
}

// the implement of CCL in our work
type CrossShardCallList struct {
	Shard_id uint8
	Call_txs *aq.Queue
}

// type=5,from=ABCD1,to=DCBA1,value=0,data=NONE,nonce=1;type=5,from=ABCD2,to=DCBA2,value=0,data=NONE,nonce=2;from=ABCD3,to=DCBA3,value=0,data=NONE,nonce=3;
func TxToCCL(tx_byte []byte, shard_id uint8) (uint32, *CrossShardCallList) {
	var ccl CrossShardCallList
	ccl.Shard_id = shard_id
	ccl.Call_txs = aq.New()
	txs := bytes.Split(tx_byte, []byte(";"))
	var temp_tx_json TransactionType
	for _, tx := range txs {
		tx_elements := bytes.Split(tx, []byte(","))
		for _, tx_element := range tx_elements {
			kv := bytes.Split(tx_element, []byte("="))
			switch {
			case string(kv[0]) == "type":
				temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				temp_tx_json.Tx_type = uint8(temp_type64)
			case string(kv[0]) == "from":
				temp_tx_json.From = make([]byte, Addr_Length)
				copy(temp_tx_json.From, kv[1])
			case string(kv[0]) == "to":
				temp_tx_json.To = make([]byte, Addr_Length)
				copy(temp_tx_json.To, kv[1])
			case string(kv[0]) == "value":
				// temp_value := string(kv[1])
				temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				temp_tx_json.Value = uint32(temp_value64)
			case string(kv[0]) == "data":
				temp_tx_json.Data = make([]byte, Data_Length)
				copy(temp_tx_json.Data, kv[1])
			case string(kv[0]) == "nonce":
				temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
				temp_tx_json.Nonce = uint32(temp_value64)
			}
		}
		ccl.Call_txs.Enqueue(temp_tx_json)
	}
	return 0, &ccl
}

type CrossShardCallLists []CrossShardCallList
