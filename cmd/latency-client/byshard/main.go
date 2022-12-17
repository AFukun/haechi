package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"time"
	// "time"
)

// ./build/latency -shardport "20057" -shardip "18.188.221.188"

var shardPort, shardIp, beaconPort, beaconIp string

func init() {
	flag.StringVar(&shardPort, "shardport", "20057", "shards chain port")
	flag.StringVar(&shardIp, "shardip", "127.0.0.1", "shards chain ip")
	flag.StringVar(&beaconPort, "beaconport", "10057", "beacon chain port")
	flag.StringVar(&beaconIp, "beaconip", "127.0.0.1", "beacon chain ip")
}

func main() {
	flag.Parse()
	// intra-shard latency
	intra_start := time.Now()
	// fmt.Println("intra-shard start in:", intra_start)
	intra_req := fmt.Sprintf("http://%v:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=ABCD,to=EFGH,value=10,data=NONE,nonce=%v\"", shardIp, shardPort, 0, 0, 0, get_rand(math.MaxInt32))
	intra_res, _ := http.Get(intra_req)
	intra_elapsed := time.Since(intra_start)
	fmt.Println("intra-shard receive", intra_res)
	fmt.Println("Byshard: intra-shard confirmation latency is:", intra_elapsed)

	// cross-shard latency
	cross_start := time.Now()
	nonce := get_rand(math.MaxInt64)
	fmt.Println("Byshard: cross-shard tx start in:", cross_start, "nonce is", nonce)
	http.Get(fmt.Sprintf("http://%v:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=CROS,to=WXYZ,value=10,data=NONE,nonce=%v\"", shardIp, shardPort, 0, 1, 1, get_rand(math.MaxInt32)))
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
