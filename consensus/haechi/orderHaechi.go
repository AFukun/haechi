//-----------------------------------------------------------------------------
// the implement of Haechi cross-shard consensus
// where there is an extra shard, AHLCoordinator, for coordination
package haechi

import (
	"github.com/AFukun/haechi/common"
)

type HaechiCoordinator struct {
	shardIPs map[int]string
	listAdd  string // TCP socket address receiving cross-shard transactions
}

func (hc *HaechiCoordinator) OrderHaechiTx(cls []common.CrossLink) {

}
