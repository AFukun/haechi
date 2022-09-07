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
	//
	// fmt.Println("sending, current time is: " + time.Now().String())
	// send cross-shard tx
	go http.Get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=ABCD,to=EFGH,value=10,data=NONE,nonce=%v\"", 21057, 1, 1, 1, get_rand(math.MaxInt64)))
	go http.Get(fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v\"", 20057, 0, 1, 1, get_rand(math.MaxInt64)))
	// fmt.Println(resp.Status)
	// fmt.Println("Received, current time is: " + time.Now().String())
	time.Sleep(30 * time.Second)
}

// string:=strconv.FormatInt(int64,10)
func get_rand(upperBond int64) string {
	maxInt := new(big.Int).SetInt64(upperBond)
	i, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		fmt.Printf("Can't generate random value: %v, %v", i, err)
	}
	outputRand := fmt.Sprintf("%v", i)
	return outputRand
}
