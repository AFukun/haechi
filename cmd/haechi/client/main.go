package main

import (
	"net/http"
)

func main() {
	/*send an cross-shard tx to shard 1 to call haechi contract*/
	request1 := "127.0.0.1:20057/broadcast_tx_commit?tx=\"fromid=0,shardid=0,type=1,from=ABCD,to=DCBA,value=20,data=DATA,nonce=1\""
	http.Get("http://" + request1)

	/*send an cross-shard tx to shard 2 to call haechi contract*/
	request2 := "127.0.0.1:21057/broadcast_tx_commit?tx=\"fromid=1,shardid=1,type=1,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=1\""
	http.Get("http://" + request2)
}
