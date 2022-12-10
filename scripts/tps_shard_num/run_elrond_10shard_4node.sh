GOSRC=$GOPATH/src
TEST_SCENE="byshard"
TM_HOME="$HOME/.haechibyshard"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=60


rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r test/byshard_10shard_4node/* $TM_HOME
echo "configs generated"

pkill -9 byshard

./build/byshard -home $TM_HOME/shard0/node0 -leader "true" -shards 10 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard0node0.log &
./build/byshard -home $TM_HOME/shard0/node1 -leader "false" -shards 10 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard0node1.log &
./build/byshard -home $TM_HOME/shard0/node2 -leader "false" -shards 10 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard0node2.log &
./build/byshard -home $TM_HOME/shard0/node3 -leader "false" -shards 10 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard0node3.log &
./build/byshard -home $TM_HOME/shard1/node0 -leader "true" -shards 10 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard1node0.log &
./build/byshard -home $TM_HOME/shard1/node1 -leader "false" -shards 10 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard1node1.log &
./build/byshard -home $TM_HOME/shard1/node2 -leader "false" -shards 10 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard1node2.log &
./build/byshard -home $TM_HOME/shard1/node3 -leader "false" -shards 10 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard1node3.log &
./build/byshard -home $TM_HOME/shard2/node0 -leader "true" -shards 10 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard2node0.log &
./build/byshard -home $TM_HOME/shard2/node1 -leader "false" -shards 10 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard2node1.log &
./build/byshard -home $TM_HOME/shard2/node2 -leader "false" -shards 10 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard2node2.log &
./build/byshard -home $TM_HOME/shard2/node3 -leader "false" -shards 10 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard2node3.log &
./build/byshard -home $TM_HOME/shard3/node0 -leader "true" -shards 10 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard3node0.log &
./build/byshard -home $TM_HOME/shard3/node1 -leader "false" -shards 10 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard3node1.log &
./build/byshard -home $TM_HOME/shard3/node2 -leader "false" -shards 10 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard3node2.log &
./build/byshard -home $TM_HOME/shard3/node3 -leader "false" -shards 10 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard3node3.log &
./build/byshard -home $TM_HOME/shard4/node0 -leader "true" -shards 10 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard4node0.log &
./build/byshard -home $TM_HOME/shard4/node1 -leader "false" -shards 10 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard4node1.log &
./build/byshard -home $TM_HOME/shard4/node2 -leader "false" -shards 10 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard4node2.log &
./build/byshard -home $TM_HOME/shard4/node3 -leader "false" -shards 10 -shardid 4 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard4node3.log &
./build/byshard -home $TM_HOME/shard5/node0 -leader "true" -shards 10 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard5node0.log &
./build/byshard -home $TM_HOME/shard5/node1 -leader "false" -shards 10 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard5node1.log &
./build/byshard -home $TM_HOME/shard5/node2 -leader "false" -shards 10 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard5node2.log &
./build/byshard -home $TM_HOME/shard5/node3 -leader "false" -shards 10 -shardid 5 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard5node3.log &
./build/byshard -home $TM_HOME/shard6/node0 -leader "true" -shards 10 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard6node0.log &
./build/byshard -home $TM_HOME/shard6/node1 -leader "false" -shards 10 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard6node1.log &
./build/byshard -home $TM_HOME/shard6/node2 -leader "false" -shards 10 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard6node2.log &
./build/byshard -home $TM_HOME/shard6/node3 -leader "false" -shards 10 -shardid 6 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard6node3.log &
./build/byshard -home $TM_HOME/shard7/node0 -leader "true" -shards 10 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard7node0.log &
./build/byshard -home $TM_HOME/shard7/node1 -leader "false" -shards 10 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard7node1.log &
./build/byshard -home $TM_HOME/shard7/node2 -leader "false" -shards 10 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard7node2.log &
./build/byshard -home $TM_HOME/shard7/node3 -leader "false" -shards 10 -shardid 7 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard7node3.log &
./build/byshard -home $TM_HOME/shard8/node0 -leader "true" -shards 10 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard8node0.log &
./build/byshard -home $TM_HOME/shard8/node1 -leader "false" -shards 10 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard8node1.log &
./build/byshard -home $TM_HOME/shard8/node2 -leader "false" -shards 10 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard8node2.log &
./build/byshard -home $TM_HOME/shard8/node3 -leader "false" -shards 10 -shardid 8 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard8node3.log &
./build/byshard -home $TM_HOME/shard9/node0 -leader "true" -shards 10 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard9node0.log &
./build/byshard -home $TM_HOME/shard9/node1 -leader "false" -shards 10 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard9node1.log &
./build/byshard -home $TM_HOME/shard9/node2 -leader "false" -shards 10 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard9node2.log &
./build/byshard -home $TM_HOME/shard9/node3 -leader "false" -shards 10 -shardid 9 -inport 10057 -outport "20057,21057,22057,23057,24057,25057,26057,27057,28057,29057" &> $LOG_DIR/shard9node3.log &

echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 byshard
echo "all done"
