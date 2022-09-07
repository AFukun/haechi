package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"time"
)

// outport
// 20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057

func main() {
	// send intra-shard tx for testing security
	// go http.Get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=ABCD,to=EFGH,value=10,data=NONE,nonce=%v,txid=%v\"", 21057, 1, 1, 0, get_rand(math.MaxInt64), get_rand(20000)))
	fmt.Println("sending, current time is: " + time.Now().String())
	// send cross-shard tx for testing confirmation latency
	resp, _ := http.Get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v,txid=%v\"", 10057, 0, 1, 1, get_rand(math.MaxInt64), get_rand(20000)))
	fmt.Println(resp.Status)
	fmt.Println("Received, current time is: " + time.Now().String())
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
