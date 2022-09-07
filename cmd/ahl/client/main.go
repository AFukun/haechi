package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"net/http"
)

func main() {
	/*send an cross-shard tx to beacon shard to call haechi contract managed in shard 1*/
	/* curl -s '127.0.0.1:10057/broadcast_tx_commit?tx="fromid=0,toid=1,type=1,from=ABCD,to=DCBA,value=10,data=NONE,nonce=1,txid=0"' */
	for true {
		request1 := fmt.Sprintf("http://127.0.0.1:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v,txid=%v\"", 20057, 0, 0, 0, get_rand(math.MaxInt64), get_rand(2000))
		go http.Get(request1)
	}
	/*send an cross-shard tx to beacon shard to call haechi contract managed in shard 0*/
	/* curl -s '127.0.0.1:10057/broadcast_tx_commit?tx="fromid=1,toid=0,type=1,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=1,txid=1"' */
	// request2 := "127.0.0.1:10057/broadcast_tx_commit?tx=\"fromid=1,toid=0,type=1,from=EFGH,to=WXYZ,value=20,data=NONE,nonce=1,txid=1\""
	// http.Get("http://" + request2)
}

func get_rand(upperBond int64) string {
	maxInt := new(big.Int).SetInt64(upperBond)
	i, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		fmt.Printf("Can't generate random value: %v, %v", i, err)
	}
	outputRand := fmt.Sprintf("%v", i)
	return outputRand
}
