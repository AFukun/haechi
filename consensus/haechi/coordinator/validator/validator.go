package validator

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"sort"
	"strconv"
	"time"

	hctypes "github.com/AFukun/haechi/types"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	dbm "github.com/tendermint/tm-db"
)

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
	ShardCLMsgs         []ShardCrosslinkMsg
	ShardBlockInterval  []int64
	ShardNextTS         []int64
	ValidTSRange        [2]int64
	ShardBlockLastTS    []int64
	Leader              bool
	input_addr          hctypes.HaechiAddress
	output_shards_addrs []hctypes.HaechiAddress
	shard_num           uint8
	currentCCLs         hctypes.CrossShardCallLists
	min_next_TS         int64
	Start_Order         bool
}

func NewValidatorInterface(bcstate *BlockchainState, shard_num uint8, leader bool, in_addr hctypes.HaechiAddress, out_addrs []hctypes.HaechiAddress) *ValidatorInterface {
	var new_validator ValidatorInterface
	new_validator.BCState = bcstate
	new_validator.shard_num = shard_num
	new_validator.ShardCLMsgs = make([]ShardCrosslinkMsg, shard_num)
	for i := uint8(0); i < shard_num; i++ {
		new_validator.ShardCLMsgs[i].CL = aq.New()
	}
	new_validator.ShardBlockInterval = make([]int64, shard_num)
	new_validator.ShardNextTS = make([]int64, shard_num)
	new_validator.ShardBlockLastTS = make([]int64, shard_num)
	for i := uint8(0); i < shard_num; i++ {
		new_validator.ShardBlockInterval = append(new_validator.ShardBlockInterval, 0)
		new_validator.ShardNextTS = append(new_validator.ShardNextTS, 0)
		new_validator.ShardBlockLastTS = append(new_validator.ShardBlockLastTS, 0)
	}
	new_validator.ShardBlockInterval = new_validator.ShardBlockInterval[1:]
	new_validator.ShardNextTS = new_validator.ShardNextTS[1:]
	new_validator.ShardBlockLastTS = new_validator.ShardBlockLastTS[1:]
	new_validator.ValidTSRange[0] = time.Now().Unix()
	new_validator.ValidTSRange[1] = math.MaxInt64
	new_validator.Leader = leader
	new_validator.input_addr = in_addr
	new_validator.output_shards_addrs = make([]hctypes.HaechiAddress, shard_num)
	new_validator.currentCCLs = make(hctypes.CrossShardCallLists, shard_num)
	for i := uint8(0); i < shard_num; i++ {
		new_validator.output_shards_addrs[i].Ip = out_addrs[i].Ip
		new_validator.output_shards_addrs[i].Port = out_addrs[i].Port
		new_validator.currentCCLs[i].Call_txs = aq.New()
	}
	new_validator.min_next_TS = math.MaxInt64
	new_validator.Start_Order = true
	return &new_validator
}

func (nw *ValidatorInterface) GlobalOrdering() {
	for shard_id := uint(0); shard_id < uint(nw.shard_num); shard_id++ {
		ccl_size := nw.currentCCLs[shard_id].Call_txs.Size()
		// // tln("beacon shardid is: " + fmt.Sprint(shard_id))
		// // tln("beacon ccls txs number is: " + fmt.Sprint(ccl_size))
		if ccl_size == 0 {
			continue
		}
		temp_cls := make([]hctypes.CrossLink, ccl_size)
		for k := uint(0); k < uint(ccl_size); k++ {
			temp_tx, _ := nw.currentCCLs[shard_id].Call_txs.Dequeue()
			cl_tx := temp_tx.(hctypes.CrossLink)
			temp_cls[k] = cl_tx
		}
		sort.SliceStable(temp_cls, func(i, j int) bool {
			return temp_cls[i].Index < temp_cls[j].Index
		})
		sort.SliceStable(temp_cls, func(i, j int) bool {
			return temp_cls[i].Block_timestamp < temp_cls[j].Block_timestamp
		})
		nw.currentCCLs[shard_id].Call_txs.Clear()

		for j := uint(0); j < uint(len(temp_cls)); j++ {
			nw.currentCCLs[shard_id].Call_txs.Enqueue(temp_cls[j])
		}
	}
}

func (nw *ValidatorInterface) DeliverCallList(shard_id uint8) {
	tx_string := ""
	for i := uint(0); !nw.currentCCLs[shard_id].Call_txs.Empty(); i++ {
		temp_tx, _ := nw.currentCCLs[shard_id].Call_txs.Dequeue()
		cl_tx := temp_tx.(hctypes.CrossLink)
		tx_string += "fromid="
		tx_string += strconv.Itoa(int(cl_tx.From_shard))
		tx_string += ",toid="
		tx_string += strconv.Itoa(int(cl_tx.To_shard))
		tx_string += ",type=5"
		tx_string += ",from="
		tx_string += string(cl_tx.From)
		tx_string += ",to="
		tx_string += string(cl_tx.To)
		tx_string += ",value="
		tx_string += strconv.Itoa(int(cl_tx.Value))
		tx_string += ",data="
		tx_string += string(cl_tx.Data)
		tx_string += ",nonce="
		tx_string += strconv.Itoa(int(cl_tx.Nonce))
		tx_string += ">"
	}
	if tx_string != "" {
		receiver_addr := net.JoinHostPort(nw.output_shards_addrs[shard_id].Ip.String(), strconv.Itoa(int(nw.output_shards_addrs[shard_id].Port)))
		request := receiver_addr
		request += "/broadcast_tx_commit?tx=\""
		request += tx_string
		request += "\""
		http.Get("http://" + request)
		// _, err := http.Get("http://" + request)
		// if err != nil {
		// 	fmt.Println("Error: deliver execution tx error when request a curl")
		// }
	}
}

func (nw *ValidatorInterface) DeliverCallLists() {
	nw.FormCCLs()
	for i := uint8(0); i < nw.shard_num; i++ {
		nw.DeliverCallList(i)
	}
	if nw.StartOrder() {
		nw.Start_Order = true
	}

}

func (nw *ValidatorInterface) FormCCLs() {
	nw.UpdateTimestampRange()
	var cls_size int
	for i := uint(0); i < uint(nw.shard_num); i++ {
		if nw.ShardCLMsgs[i].CL.Empty() {
			continue
		}
		cls_size = nw.ShardCLMsgs[i].CL.Size()
		// // tln("beacon formccls, cls size is: " + fmt.Sprint(cls_size))
		for j := uint(0); j < uint(cls_size); j++ {
			cl_temp, _ := nw.ShardCLMsgs[i].CL.Dequeue()
			cl := cl_temp.(hctypes.CrossLink)
			if cl.Block_timestamp > nw.ValidTSRange[1] {
				// advanced cross link
				// // tln("This is an advanced ccl")
				nw.ShardCLMsgs[i].CL.Enqueue(cl)
			} else {
				nw.currentCCLs[cl.To_shard].Call_txs.Enqueue(cl)
			}
			// nw.currentCCLs[cl.To_shard].Call_txs.Enqueue(cl)
		}
	}
	nw.GlobalOrdering()
}

func (nw *ValidatorInterface) UpdateShardCrosslinkMsgs(request []byte) {
	shardid := CheckFromShardId(request)
	// MicroBench: test time difference of shard's CrossLinks
	temp_output := fmt.Sprintf("current block height is %v, receive CrossLink from %v, at time %v", nw.BCState.Height, shardid, time.Now())
	fmt.Println(temp_output)
	_, crosslinks := hctypes.RequestToCrossLinks(request)
	for _, crosslink := range crosslinks {
		nw.ShardCLMsgs[shardid].CL.Enqueue(crosslink)
	}
}

func (nw *ValidatorInterface) UpdateOrderParameters(request []byte) {
	blockts := CheckBlockTimestamp(request)
	shardid := CheckFromShardId(request)
	current_interval := blockts - nw.ShardBlockLastTS[shardid]
	nw.ShardBlockInterval[shardid] = current_interval
	nw.ShardBlockLastTS[shardid] = blockts

	nw.ShardNextTS[shardid] = blockts + current_interval
	if nw.ShardNextTS[shardid] < nw.min_next_TS {
		nw.min_next_TS = nw.ShardNextTS[shardid]
	}
}

func (nw *ValidatorInterface) UpdateTimestampRange() {
	nw.ValidTSRange[0] = nw.ValidTSRange[1]
	nw.ValidTSRange[1] = nw.min_next_TS

}

func (nw *ValidatorInterface) StartOrder() bool {
	start := true
	for _, ccl := range nw.ShardCLMsgs {
		if ccl.CL.Empty() {
			start = false
		}
	}
	return start
}

func CheckFromShardId(tx []byte) uint8 {
	var fromshardid uint8
	tx_elements := bytes.Split(tx, []byte(","))
	for _, tx_element := range tx_elements {
		kv := bytes.Split(tx_element, []byte("="))
		if string(kv[0]) == "fromid" {
			temp_type64, _ := strconv.ParseUint(string(kv[1]), 10, 64)
			fromshardid = uint8(temp_type64)
			break
		}
	}
	return fromshardid
}

func CheckBlockTimestamp(tx []byte) int64 {
	var bts int64
	txs := bytes.Split(tx, []byte(">"))
	tx_elements := bytes.Split(txs[0], []byte(","))
	for _, tx_element := range tx_elements {
		kv := bytes.Split(tx_element, []byte("="))
		if string(kv[0]) == "blocktimestamp" {
			temp_type64, _ := strconv.ParseInt(string(kv[1]), 10, 64)
			bts = temp_type64
			break
		}
	}

	return bts
}
