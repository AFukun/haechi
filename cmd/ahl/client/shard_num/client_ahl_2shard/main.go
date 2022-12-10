package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"sync"
)

// outport
// 20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057

func main() {
	//go create_request()
	tx_num := 10
	cross_rate := float32(0.9)
	shard_num := 2
	for true {
		// create the same number of txs for each shard, with the same cross shard rate
		go send_request(20057, 10057, tx_num, 0, shard_num, cross_rate)
		go send_request(21057, 10057, tx_num, 1, shard_num, cross_rate)
		// go send_request(22057, 10057, tx_num, 2, shard_num, cross_rate)
		// go send_request(23057, 10057, tx_num, 3, shard_num, cross_rate)
		// go send_request(24057, 10057, tx_num, 4, shard_num, cross_rate)
		// go send_request(25057, 10057, tx_num, 5, shard_num, cross_rate)
		// go send_request(26057, 10057, tx_num, 6, shard_num, cross_rate)
		// go send_request(27057, 10057, tx_num, 7, shard_num, cross_rate)
		// go send_request(28057, 10057, tx_num, 8, shard_num, cross_rate)
		// go send_request(29057, 10057, tx_num, 9, shard_num, cross_rate)
		// go send_request(30057, 10057, tx_num, 10, shard_num, cross_rate)
		// go send_request(31057, 10057, tx_num, 11, shard_num, cross_rate)
		// go send_request(32057, 10057, tx_num, 12, shard_num, cross_rate)
		// go send_request(33057, 10057, tx_num, 13, shard_num, cross_rate)
		// go send_request(34057, 10057, tx_num, 14, shard_num, cross_rate)
		// go send_request(35057, 10057, tx_num, 15, shard_num, cross_rate)
		// time.Sleep(20 * time.Millisecond)
	}
	// time.Sleep(60*time.Second)
}

/*
	outport -> leader rpc port
	tx_num -> batch of txs create together
	fromid -> tx from which shard
	shard_num -> total num of shards
	cross_rate -> cross tx rate
	toid is generated randomly [0, shardNum)
*/
func create_request(shardport int, outport int, txNum int, fromid int, shard_num int, cross_rate float32) {
	send_request(shardport, outport, txNum, fromid, shard_num, cross_rate)
	// var wg sync.WaitGroup
	// wg.Add(txNum)
	// for i := 0; i < txNum; i++ {
	// 	if i < int(cross_rate*float32(txNum)) {
	// 		go http_get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v,txid=%v\"", outport, fromid, get_rand(int64(shard_num)), 1, get_rand(math.MaxInt64), get_rand(2000)), &wg)
	// 	} else {
	// 		go http_get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v,txid=%v\"", shardport, fromid, get_rand(int64(shard_num)), 0, get_rand(math.MaxInt64), i), &wg)
	// 	}
	// }
	// wg.Wait()
}

func http_get(request string, wg *sync.WaitGroup) {
	http.Get(request)
	wg.Done()
}

//string:=strconv.FormatInt(int64,10)
func get_rand(upperBond int64) string {
	maxInt := new(big.Int).SetInt64(upperBond)
	i, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		fmt.Printf("Can't generate random value: %v, %v", i, err)
	}
	outputRand := fmt.Sprintf("%v", i)
	return outputRand
}

func send_request(shardport int, outport int, txNum int, fromid int, shard_num int, cross_rate float32) {
	ctx_num := int(float32(txNum) * cross_rate)
	for i := 0; i < int(txNum-ctx_num); i++ {
		http.Get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v,txid=%v\"", shardport, fromid, get_rand(int64(shard_num)), 0, get_rand(math.MaxInt64), i))
	}
	for i := 0; i < ctx_num; i++ {
		http.Get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v,txid=%v\"", outport, fromid, get_rand(int64(shard_num)), 1, get_rand(math.MaxInt64), get_rand(20000)))
	}

}
