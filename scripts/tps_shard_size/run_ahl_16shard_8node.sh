GOSRC=$GOPATH/src
TEST_SCENE="ahl"
TM_HOME="$HOME/.haechiAhl"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=60


rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r test/ahl_16shard_8node/* $TM_HOME
echo "configs generated"

pkill -9 ahlbc
pkill -9 ahlshard

./build/ahlshard -home $TM_HOME/shard0/node0 -leader "true" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node0.log &
./build/ahlshard -home $TM_HOME/shard0/node1 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node1.log &
./build/ahlshard -home $TM_HOME/shard0/node2 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node2.log &
./build/ahlshard -home $TM_HOME/shard0/node3 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node3.log &
./build/ahlshard -home $TM_HOME/shard0/node4 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node4.log &
./build/ahlshard -home $TM_HOME/shard0/node5 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node5.log &
./build/ahlshard -home $TM_HOME/shard0/node6 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node6.log &
./build/ahlshard -home $TM_HOME/shard0/node7 -leader "false" -shards 16 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard0node7.log &
./build/ahlshard -home $TM_HOME/shard1/node0 -leader "true" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node0.log &
./build/ahlshard -home $TM_HOME/shard1/node1 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node1.log &
./build/ahlshard -home $TM_HOME/shard1/node2 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node2.log &
./build/ahlshard -home $TM_HOME/shard1/node3 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node3.log &
./build/ahlshard -home $TM_HOME/shard1/node4 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node4.log &
./build/ahlshard -home $TM_HOME/shard1/node5 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node5.log &
./build/ahlshard -home $TM_HOME/shard1/node6 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node6.log &
./build/ahlshard -home $TM_HOME/shard1/node7 -leader "false" -shards 16 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard1node7.log &
./build/ahlshard -home $TM_HOME/shard2/node0 -leader "true" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node0.log &
./build/ahlshard -home $TM_HOME/shard2/node1 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node1.log &
./build/ahlshard -home $TM_HOME/shard2/node2 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node2.log &
./build/ahlshard -home $TM_HOME/shard2/node3 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node3.log &
./build/ahlshard -home $TM_HOME/shard2/node4 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node4.log &
./build/ahlshard -home $TM_HOME/shard2/node5 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node5.log &
./build/ahlshard -home $TM_HOME/shard2/node6 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node6.log &
./build/ahlshard -home $TM_HOME/shard2/node7 -leader "false" -shards 16 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard2node7.log &
./build/ahlshard -home $TM_HOME/shard3/node0 -leader "true" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node0.log &
./build/ahlshard -home $TM_HOME/shard3/node1 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node1.log &
./build/ahlshard -home $TM_HOME/shard3/node2 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node2.log &
./build/ahlshard -home $TM_HOME/shard3/node3 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node3.log &
./build/ahlshard -home $TM_HOME/shard3/node4 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node4.log &
./build/ahlshard -home $TM_HOME/shard3/node5 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node5.log &
./build/ahlshard -home $TM_HOME/shard3/node6 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node6.log &
./build/ahlshard -home $TM_HOME/shard3/node7 -leader "false" -shards 16 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard3node7.log &
./build/ahlshard -home $TM_HOME/shard4/node0 -leader "true" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node0.log &
./build/ahlshard -home $TM_HOME/shard4/node1 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node1.log &
./build/ahlshard -home $TM_HOME/shard4/node2 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node2.log &
./build/ahlshard -home $TM_HOME/shard4/node3 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node3.log &
./build/ahlshard -home $TM_HOME/shard4/node4 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node4.log &
./build/ahlshard -home $TM_HOME/shard4/node5 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node5.log &
./build/ahlshard -home $TM_HOME/shard4/node6 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node6.log &
./build/ahlshard -home $TM_HOME/shard4/node7 -leader "false" -shards 16 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard4node7.log &
./build/ahlshard -home $TM_HOME/shard5/node0 -leader "true" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node0.log &
./build/ahlshard -home $TM_HOME/shard5/node1 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node1.log &
./build/ahlshard -home $TM_HOME/shard5/node2 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node2.log &
./build/ahlshard -home $TM_HOME/shard5/node3 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node3.log &
./build/ahlshard -home $TM_HOME/shard5/node4 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node4.log &
./build/ahlshard -home $TM_HOME/shard5/node5 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node5.log &
./build/ahlshard -home $TM_HOME/shard5/node6 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node6.log &
./build/ahlshard -home $TM_HOME/shard5/node7 -leader "false" -shards 16 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard5node7.log &
./build/ahlshard -home $TM_HOME/shard6/node0 -leader "true" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node0.log &
./build/ahlshard -home $TM_HOME/shard6/node1 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node1.log &
./build/ahlshard -home $TM_HOME/shard6/node2 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node2.log &
./build/ahlshard -home $TM_HOME/shard6/node3 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node3.log &
./build/ahlshard -home $TM_HOME/shard6/node4 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node4.log &
./build/ahlshard -home $TM_HOME/shard6/node5 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node5.log &
./build/ahlshard -home $TM_HOME/shard6/node6 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node6.log &
./build/ahlshard -home $TM_HOME/shard6/node7 -leader "false" -shards 16 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard6node7.log &
./build/ahlshard -home $TM_HOME/shard7/node0 -leader "true" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node0.log &
./build/ahlshard -home $TM_HOME/shard7/node1 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node1.log &
./build/ahlshard -home $TM_HOME/shard7/node2 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node2.log &
./build/ahlshard -home $TM_HOME/shard7/node3 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node3.log &
./build/ahlshard -home $TM_HOME/shard7/node4 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node4.log &
./build/ahlshard -home $TM_HOME/shard7/node5 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node5.log &
./build/ahlshard -home $TM_HOME/shard7/node6 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node6.log &
./build/ahlshard -home $TM_HOME/shard7/node7 -leader "false" -shards 16 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard7node7.log &
./build/ahlshard -home $TM_HOME/shard8/node0 -leader "true" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node0.log &
./build/ahlshard -home $TM_HOME/shard8/node1 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node1.log &
./build/ahlshard -home $TM_HOME/shard8/node2 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node2.log &
./build/ahlshard -home $TM_HOME/shard8/node3 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node3.log &
./build/ahlshard -home $TM_HOME/shard8/node4 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node4.log &
./build/ahlshard -home $TM_HOME/shard8/node5 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node5.log &
./build/ahlshard -home $TM_HOME/shard8/node6 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node6.log &
./build/ahlshard -home $TM_HOME/shard8/node7 -leader "false" -shards 16 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard8node7.log &
./build/ahlshard -home $TM_HOME/shard9/node0 -leader "true" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node0.log &
./build/ahlshard -home $TM_HOME/shard9/node1 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node1.log &
./build/ahlshard -home $TM_HOME/shard9/node2 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node2.log &
./build/ahlshard -home $TM_HOME/shard9/node3 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node3.log &
./build/ahlshard -home $TM_HOME/shard9/node4 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node4.log &
./build/ahlshard -home $TM_HOME/shard9/node5 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node5.log &
./build/ahlshard -home $TM_HOME/shard9/node6 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node6.log &
./build/ahlshard -home $TM_HOME/shard9/node7 -leader "false" -shards 16 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard9node7.log &
./build/ahlshard -home $TM_HOME/shard10/node0 -leader "true" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node0.log &
./build/ahlshard -home $TM_HOME/shard10/node1 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node1.log &
./build/ahlshard -home $TM_HOME/shard10/node2 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node2.log &
./build/ahlshard -home $TM_HOME/shard10/node3 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node3.log &
./build/ahlshard -home $TM_HOME/shard10/node4 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node4.log &
./build/ahlshard -home $TM_HOME/shard10/node5 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node5.log &
./build/ahlshard -home $TM_HOME/shard10/node6 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node6.log &
./build/ahlshard -home $TM_HOME/shard10/node7 -leader "false" -shards 16 -shardid 10 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard10node7.log &
./build/ahlshard -home $TM_HOME/shard11/node0 -leader "true" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node0.log &
./build/ahlshard -home $TM_HOME/shard11/node1 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node1.log &
./build/ahlshard -home $TM_HOME/shard11/node2 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node2.log &
./build/ahlshard -home $TM_HOME/shard11/node3 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node3.log &
./build/ahlshard -home $TM_HOME/shard11/node4 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node4.log &
./build/ahlshard -home $TM_HOME/shard11/node5 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node5.log &
./build/ahlshard -home $TM_HOME/shard11/node6 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node6.log &
./build/ahlshard -home $TM_HOME/shard11/node7 -leader "false" -shards 16 -shardid 11 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard11node7.log &
./build/ahlshard -home $TM_HOME/shard12/node0 -leader "true" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node0.log &
./build/ahlshard -home $TM_HOME/shard12/node1 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node1.log &
./build/ahlshard -home $TM_HOME/shard12/node2 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node2.log &
./build/ahlshard -home $TM_HOME/shard12/node3 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node3.log &
./build/ahlshard -home $TM_HOME/shard12/node4 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node4.log &
./build/ahlshard -home $TM_HOME/shard12/node5 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node5.log &
./build/ahlshard -home $TM_HOME/shard12/node6 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node6.log &
./build/ahlshard -home $TM_HOME/shard12/node7 -leader "false" -shards 16 -shardid 12 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard12node7.log &
./build/ahlshard -home $TM_HOME/shard13/node0 -leader "true" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node0.log &
./build/ahlshard -home $TM_HOME/shard13/node1 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node1.log &
./build/ahlshard -home $TM_HOME/shard13/node2 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node2.log &
./build/ahlshard -home $TM_HOME/shard13/node3 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node3.log &
./build/ahlshard -home $TM_HOME/shard13/node4 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node4.log &
./build/ahlshard -home $TM_HOME/shard13/node5 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node5.log &
./build/ahlshard -home $TM_HOME/shard13/node6 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node6.log &
./build/ahlshard -home $TM_HOME/shard13/node7 -leader "false" -shards 16 -shardid 13 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard13node7.log &
./build/ahlshard -home $TM_HOME/shard14/node0 -leader "true" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node0.log &
./build/ahlshard -home $TM_HOME/shard14/node1 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node1.log &
./build/ahlshard -home $TM_HOME/shard14/node2 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node2.log &
./build/ahlshard -home $TM_HOME/shard14/node3 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node3.log &
./build/ahlshard -home $TM_HOME/shard14/node4 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node4.log &
./build/ahlshard -home $TM_HOME/shard14/node5 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node5.log &
./build/ahlshard -home $TM_HOME/shard14/node6 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node6.log &
./build/ahlshard -home $TM_HOME/shard14/node7 -leader "false" -shards 16 -shardid 14 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard14node7.log &
./build/ahlshard -home $TM_HOME/shard15/node0 -leader "true" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node0.log &
./build/ahlshard -home $TM_HOME/shard15/node1 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node1.log &
./build/ahlshard -home $TM_HOME/shard15/node2 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node2.log &
./build/ahlshard -home $TM_HOME/shard15/node3 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node3.log &
./build/ahlshard -home $TM_HOME/shard15/node4 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node4.log &
./build/ahlshard -home $TM_HOME/shard15/node5 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node5.log &
./build/ahlshard -home $TM_HOME/shard15/node6 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node6.log &
./build/ahlshard -home $TM_HOME/shard15/node7 -leader "false" -shards 16 -shardid 15 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/shard15node7.log &
./build/ahlbc -home $TM_HOME/beacon/node0 -leader "true" -shards 16 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/beaconnode0.log &
./build/ahlbc -home $TM_HOME/beacon/node1 -leader "false" -shards 16 -inport 10157 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/beaconnode1.log &
./build/ahlbc -home $TM_HOME/beacon/node2 -leader "false" -shards 16 -inport 10257 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/beaconnode2.log &
./build/ahlbc -home $TM_HOME/beacon/node3 -leader "false" -shards 16 -inport 10357 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057,30057,31057,32057,33057,34057,35057" &> $LOG_DIR/beaconnode3.log &

echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 ahlbc
pkill -9 ahlshard
echo "all done"
