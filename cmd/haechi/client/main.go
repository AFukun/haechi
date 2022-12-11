package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

var shardNum, batchSize, beaconPort, concurrentNum, reqDuration uint

// var requestRate int64
var crossRate float64
var shardPorts, beaconIp, shardIps string

func init() {
	flag.UintVar(&shardNum, "shards", 2, "the number of shards")
	flag.UintVar(&batchSize, "batch", 10, "the batch size of one request")
	flag.UintVar(&concurrentNum, "parallel", 100, "concurrent number for sending requests")
	flag.UintVar(&reqDuration, "duration", 120, "duration of sending request")
	// flag.Int64Var(&requestRate, "rate", 100, "the request rate controlled by sleeping time, ms")
	flag.Float64Var(&crossRate, "ratio", 0.8, "the ratio of cross-shard txs")

	flag.UintVar(&beaconPort, "beaconport", 10057, "beacon chain port")
	flag.StringVar(&shardPorts, "shardports", "20057,21057", "shards chain port")
	flag.StringVar(&beaconIp, "beaconip", "127.0.0.1", "beacon chain ip")
	flag.StringVar(&shardIps, "shardips", "127.0.0.1,127.0.0.1", "shards chain ip")
}

func main() {
	flag.Parse()
	// initialize ip
	shard_ports_temp := []byte(shardPorts)
	shard_ports := bytes.Split(shard_ports_temp, []byte(","))
	shard_ips_temp := []byte(shardIps)
	shard_ips := bytes.Split(shard_ips_temp, []byte(","))
	var ports_value64 []uint64
	for _, shard_port := range shard_ports {
		temp_port, _ := strconv.ParseUint(string(shard_port), 10, 64)
		ports_value64 = append(ports_value64, temp_port)
	}
	for p := 0; p < int(concurrentNum); p++ {
		for i, _ := range shard_ports {
			go send_request(uint(ports_value64[i]), string(shard_ips[i]), beaconPort, beaconIp, batchSize, uint(i), shardNum, crossRate)
		}
		// time.Sleep(time.Duration(requestRate) * time.Millisecond)
	}
	time.Sleep(time.Duration(reqDuration) * time.Second)
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

func send_request(s_port uint, s_ip string, b_port uint, b_ip string, tx_num uint, from_id uint, s_num uint, cross_rate float64) {
	for {
		ctx_num := uint(float64(tx_num) * cross_rate)
		for i := uint(0); i < tx_num-ctx_num; i++ {
			http.Get(fmt.Sprintf("http://%v:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=ABCD,to=EFGH,value=10,data=NONE,nonce=%v\"", s_ip, s_port, from_id, get_rand(int64(s_num)), 0, get_rand(math.MaxInt32)))
		}
		for i := uint(0); i < ctx_num; i++ {
			http.Get(fmt.Sprintf("http://%v:%v/broadcast_tx_commit?tx=\"fromid=%v,toid=%v,type=%v,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=%v\"", s_ip, s_port, from_id, get_rand(int64(s_num)), 1, get_rand(math.MaxInt32)))
		}
	}
}
