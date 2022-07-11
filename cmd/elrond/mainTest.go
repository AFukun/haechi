package main

import (
	"fmt"

	haechiNode "github.com/AFukun/haechi/consensus/haechi/shard/validator"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
)

func main() {
	aq_test := aq.New()
	tx_json := haechiNode.TransactionType{
		Tx_type: 2,
		From:    []byte("ABCD"),
		To:      []byte("ABCD"),
		Value:   5,
		Data:    []byte("ABCD"),
	}
	fmt.Println("tx json is: %v", tx_json)
	aq_test.Enqueue(tx_json)
	aq_ele, _ := aq_test.Dequeue()
	tx_from := aq_ele.(haechiNode.TransactionType)
	fmt.Println("tx from is: %s", string(tx_from.From))
	fmt.Println("tx to is: %s", tx_from.To)
}
