package validator

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	haechitypes "github.com/AFukun/haechi/core/types"
	dbm "github.com/tendermint/tm-db"
)

/*
A transaction is defined as a json including:
{
	tx_type uint32
	from	[20]byte
	to		[20]byte
	value	uint32
	data	[20]byte
}
e.g., sent by user as: "type=0,from=ABCD,to=DCBA,value=0,data=NONE"
*/
const (
	Addr_Length uint8 = 4
	Data_Length uint8 = 4
)
const (
	Type_Num              uint8 = 5
	IntraShard_TX         uint8 = 0
	InterShard_TX_Verify  uint8 = 1
	InterShard_TX_Execute uint8 = 2
	InterShard_TX_Commit  uint8 = 3
	InterShard_TX_Update  uint8 = 4
)

// var (
// 	stateKey = []byte("stateKey")
// )

type BlockchainState struct {
	Database dbm.DB
	Size     int64
	Height   int64
	AppHash  []byte
}

func NewBlockchainState(name string, dir string) *BlockchainState {
	var bcstate BlockchainState
	var err error
	// bcstate, _ := loadState(dbm.DB.(name, dir))
	bcstate.Database, err = dbm.NewDB(name, dbm.GoLevelDBBackend, dir)
	bcstate.Height = 0
	bcstate.Size = 0
	if err != nil {
		log.Fatalf("Create database error: %v", err)
	}
	return &bcstate
}

type ValidatorInterface struct {
	BCState           *BlockchainState
	Leader            bool
	ip_input_shard    net.IP // send commit tx to the input shard
	port_input_shard  uint16
	ip_output_shard   net.IP // send execute tx to the output shard
	port_output_shard uint16
}

func NewValidatorInterface(bcstate *BlockchainState, leader bool, ip_in net.IP, port_in uint16, ip_out net.IP, port_out uint16) *ValidatorInterface {
	return &ValidatorInterface{
		BCState:           bcstate,
		Leader:            leader,
		ip_input_shard:    ip_in,
		port_input_shard:  port_in,
		ip_output_shard:   ip_out,
		port_output_shard: port_out,
	}
}

func (nw *ValidatorInterface) DeliverExecutionTx(tx []byte) {
	tx_str := string(tx)
	receiver_addr := net.JoinHostPort(nw.ip_output_shard.String(), strconv.Itoa(int(nw.port_output_shard)))
	request := receiver_addr
	request += "/broadcast_tx_commit?tx=\""
	request += tx_str
	request += "\""
	_, err := http.Get("http://" + request)
	if err != nil {
		fmt.Println("Error: deliver execution tx error when request a curl")
	}
}

func (nw *ValidatorInterface) DeliverCommitTx(tx []byte) {
	tx_str := string(tx)
	sender_addr := net.JoinHostPort(nw.ip_input_shard.String(), strconv.Itoa(int(nw.port_input_shard)))
	request := sender_addr
	request += "/broadcast_tx_commit?tx=\""
	request += tx_str
	request += "\""
	_, err := http.Get("http://" + request)
	if err != nil {
		fmt.Println("Error: deliver execution tx error when request a curl")
	}
}

func (nw *ValidatorInterface) DeliverUpdateTx(tx []byte) {
	// output_addr := net.JoinHostPort(nw.ip_output_shard.String(), strconv.Itoa(int(nw.port_output_shard)))
	tx_str := string(tx)
	receiver_addr := net.JoinHostPort(nw.ip_output_shard.String(), strconv.Itoa(int(nw.port_output_shard)))
	request := receiver_addr
	request += "/broadcast_tx_commit?tx=\""
	request += tx_str
	request += "\""
	_, err := http.Get("http://" + request)
	if err != nil {
		fmt.Println("Error: deliver execution tx error when request a curl")
	}
}

func Serilization(tx []byte) (uint32, haechitypes.TransactionType) {
	var tx_json haechitypes.TransactionType
	// tx_json.Tx_type = 4
	txElements := bytes.Split(tx, []byte(","))
	if len(txElements) == 0 {
		return 1, tx_json
	}
	for _, txElement := range txElements {
		kv := bytes.Split(txElement, []byte("="))
		switch {
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
		}
	}
	return 0, tx_json
}

func Deserilization(tx haechitypes.TransactionType) (uint32, []byte) {
	var tempStr string = ""
	tempStr += "type="
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
	tx_byte := []byte(tempStr)
	return 0, tx_byte
}

func ResolveTx(_tx []byte) (uint32, haechitypes.TransactionType) {
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

func GetTxType(_tx []byte) uint8 {
	_, tx := Serilization(_tx)
	return tx.Tx_type
}

func AddOperation(add []byte, value uint32) []byte {
	temp_add := binary.BigEndian.Uint32(add)
	temp_add += value
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, temp_add)
	return buf.Bytes()
}

func SubOperation(minuend []byte, value uint32) []byte {
	temp_minuend := binary.BigEndian.Uint32(minuend)
	if temp_minuend >= value {
		temp_minuend -= value
	}
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, temp_minuend)
	return buf.Bytes()
}
