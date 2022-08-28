package main

import (
	"net/http"
)

func main() {
	/*send an cross-shard tx to beacon shard to call haechi contract managed in shard 1*/
	/* curl -s '127.0.0.1:10057/broadcast_tx_commit?tx="fromid=0,toid=1,type=1,from=ABCD,to=DCBA,value=10,data=NONE,nonce=1,txid=0"' */
	request1 := "127.0.0.1:10057/broadcast_tx_commit?tx=\"fromid=0,toid=1,type=1,from=ABCD,to=DCBA,value=20,data=NONE,nonce=1,txid=0\""
	http.Get("http://" + request1)

	/*send an cross-shard tx to beacon shard to call haechi contract managed in shard 0*/
	/* curl -s '127.0.0.1:10057/broadcast_tx_commit?tx="fromid=1,toid=0,type=1,from=EFGH,to=WXYZ,value=10,data=NONE,nonce=1,txid=1"' */
	request2 := "127.0.0.1:10057/broadcast_tx_commit?tx=\"fromid=1,toid=0,type=1,from=EFGH,to=WXYZ,value=20,data=NONE,nonce=1,txid=1\""
	http.Get("http://" + request2)
}
