package main

import (
	"github.com/AFukun/haechi/common"
	"github.com/AFukun/haechi/core/types"
	"github.com/AFukun/haechi/log"
	"github.com/AFukun/haechi/tools"
	"sync"
)

func main() {
	FROM, _ := common.HexStringToAddress("1a2b3c4d5e")
	TO, _ := common.HexStringToAddress("1b2c3d4e5f")
	IP1 := "127.0.0.1:21057"

	var wg sync.WaitGroup
	wg.Add(10)
	logger := log.New("client")
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			tx := types.AhlTx{
				From:      FROM,
				To:        TO,
				FromShard: 1,
				ToStard:   2,
				Value:     1,
				Data:      1,
				Nonce:     uint32(i),
			}
			status, err := tools.SendTxString(IP1, tx.EncodeToBase64String())
			if err != nil {
				logger.Error("error at sending tx", "err", err, "tx", tx.HashString())
			}
			logger.Info(status, "ip", IP1, "tx", tx.HashString())
		}(i)
	}

	wg.Wait()
}
