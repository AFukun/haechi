package main

import (
	"github.com/AFukun/haechi/common"
	"github.com/AFukun/haechi/core/types"
	"github.com/AFukun/haechi/tools"
	"log"
	"sync"
)

func main() {
	from, _ := common.HexStringToAddress("1a2b3c4d5e")
	to, _ := common.HexStringToAddress("1b2c3d4e5f")

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			tx := types.AhlTx{
				From:      from,
				To:        to,
				FromShard: 1,
				ToStard:   1,
				Value:     1,
				Data:      1,
				Nonce:     uint32(i),
			}
			status, err := tools.SendTxString("127.0.0.1:21057", tx.EncodeToString())
			if err != nil {
				log.Fatalln("tx number:", i+1, err)
			}
			log.Println("tx number:", i+1, status)
		}(i)
	}

	wg.Wait()
}
