package validator

import (
	"bytes"
	"log"
	"net"
	"net/http"
	"strconv"

	hctypes "github.com/AFukun/haechi/types"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	dbm "github.com/tendermint/tm-db"
)

const (
	Addr_Length    uint8  = 4
	Data_Length    uint8  = 4
	Process_Length uint32 = 20000
)
const (
	Type_Num                      uint8 = 7
	IntraShard_TX                 uint8 = 0
	InterShard_TX_Verify          uint8 = 1
	InterShard_TX_Execute         uint8 = 2
	InterShard_TX_Commit_sender   uint8 = 3
	InterShard_TX_Commit_receiver uint8 = 4
	InterShard_TX_Update_sender   uint8 = 5
	InterShard_TX_Update_receiver uint8 = 6
)

type TransactionType struct {
	From_shard uint8 // the sender's shard
	To_shard   uint8 // the receiver's shard
	Tx_type    uint8
	From       []byte
	To         []byte
	Value      uint32
	Data       []byte
	Nonce      uint32 // TODO: enable contineous tx requests by setting vary nonce
	TX_id      uint32
}

type BlockchainState struct {
	Database dbm.DB
	Size     int64
	Height   int64
	AppHash  []byte
}

func NewBlockchainState(name string, dir string) *BlockchainState {
	var bcstate BlockchainState
	var err error
	bcstate.Database, err = dbm.NewDB(name, dbm.GoLevelDBBackend, dir)
	bcstate.Height = 1
	bcstate.Size = 0
	if err != nil {
		log.Fatalf("Create database error: %v", err)
	}
	return &bcstate
}

type ShardCrosslinkMsg struct {
	CL *aq.Queue // queue used to store CrossLink
}

type ValidatorInterface struct {
	BCState             *BlockchainState
	shard_num           uint8
	Leader              bool
	input_addr          hctypes.HaechiAddress
	output_shards_addrs []hctypes.HaechiAddress
	Tx_set              [Process_Length]uint8
}

func NewValidatorInterface(bcstate *BlockchainState, shard_num uint8, leader bool, in_addr hctypes.HaechiAddress, out_addrs []hctypes.HaechiAddress) *ValidatorInterface {
	var new_validator ValidatorInterface
	new_validator.BCState = bcstate
	new_validator.shard_num = shard_num
	new_validator.Leader = leader
	new_validator.input_addr = in_addr
	new_validator.output_shards_addrs = make([]hctypes.HaechiAddress, shard_num)
	for i := uint8(0); i < shard_num; i++ {
		new_validator.output_shards_addrs[i].Ip = out_addrs[i].Ip
		new_validator.output_shards_addrs[i].Port = out_addrs[i].Port
	}
	for i := uint32(0); i < Process_Length; i++ {
		new_validator.Tx_set[i] = 0
	}
	return &new_validator
}

func (nw *ValidatorInterface) DeliverCrossShardTx(tx []byte, shardid uint8) {
	tx_str := string(tx)
	receiver_addr := net.JoinHostPort(nw.output_shards_addrs[shardid].Ip.String(), strconv.Itoa(int(nw.output_shards_addrs[shardid].Port)))
	request := receiver_addr
	request += "/broadcast_tx_commit?tx=\""
	request += tx_str
	request += "\""
	http.Get("http://" + request)
	// _, err := http.Get("http://" + request)
	// if err != nil {
	// 	fmt.Println("Error: deliver execution tx error when request a curl")
	// }
}

func (nw *ValidatorInterface) DeliverCommitTx(tx []byte, shardid uint8) {
	tx_str := string(tx)
	receiver_addr := net.JoinHostPort(nw.output_shards_addrs[shardid].Ip.String(), strconv.Itoa(int(nw.output_shards_addrs[shardid].Port)))
	request := receiver_addr
	request += "/broadcast_tx_commit?tx=\""
	request += tx_str
	request += "\""
	http.Get("http://" + request)
	// _, err := http.Get("http://" + request)
	// if err != nil {
	// 	fmt.Println("Error: deliver execution tx error when request a curl")
	// }
}

func Serilization(tx []byte) (uint32, TransactionType) {
	var tx_json TransactionType
	// tx_json.Tx_type = 4
	txElements := bytes.Split(tx, []byte(","))
	if len(txElements) == 0 {
		return 1, tx_json
	}
	for _, txElement := range txElements {
		kv := bytes.Split(txElement, []byte("="))
		switch {
		case string(kv[0]) == "fromid":
			temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			tx_json.From_shard = uint8(temp_type64)
		case string(kv[0]) == "toid":
			temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			tx_json.To_shard = uint8(temp_type64)
		case string(kv[0]) == "type":
			temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			tx_json.Tx_type = uint8(temp_type64)
		case string(kv[0]) == "from":
			tx_json.From = make([]byte, Addr_Length)
			copy(tx_json.From, kv[1])
		case string(kv[0]) == "to":
			tx_json.To = make([]byte, Addr_Length)
			copy(tx_json.To, kv[1])
		case string(kv[0]) == "value":
			// temp_value := string(kv[1])
			temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			tx_json.Value = uint32(temp_value64)
		case string(kv[0]) == "data":
			tx_json.Data = make([]byte, Data_Length)
			copy(tx_json.Data, kv[1])
		case string(kv[0]) == "nonce":
			temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			tx_json.Nonce = uint32(temp_value64)
		case string(kv[0]) == "txid":
			temp_value64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			tx_json.TX_id = uint32(temp_value64)
		}
	}
	return 0, tx_json
}

func Deserilization(tx TransactionType) (uint32, []byte) {
	var tempStr string = ""
	tempStr += "fromid="
	tempStr += strconv.Itoa(int(tx.From_shard))
	tempStr += ",toid="
	tempStr += strconv.Itoa(int(tx.To_shard))
	tempStr += ",type="
	tempStr += strconv.Itoa(int(tx.Tx_type))
	tempStr += ",from="
	tempStr += string(tx.From)
	tempStr += ",to="
	tempStr += string(tx.To)
	tempStr += ",value="
	tempStr += strconv.Itoa(int(tx.Value))
	tempStr += ",data="
	tempStr += string(tx.Data)
	tempStr += ",nonce="
	tempStr += strconv.Itoa(int(tx.Nonce))
	tempStr += ",txid="
	tempStr += strconv.Itoa(int(tx.TX_id))
	tx_byte := []byte(tempStr)
	return 0, tx_byte
}

func ResolveTx(_tx []byte) (uint32, TransactionType) {
	isValid, tx := Serilization(_tx)
	if isValid != 0 {
		return 1, tx
	}
	if tx.Tx_type < Type_Num {
		return 0, tx
	} else {
		return 1, tx
	}
	// txType := IntraShard_TX
	// return txType
}
