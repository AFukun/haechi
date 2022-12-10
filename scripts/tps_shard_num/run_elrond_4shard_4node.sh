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

cp -r test/byshard_4shard_4node/* $TM_HOME
echo "configs generated"

pkill -9 byshard

./build/byshard -home $TM_HOME/shard0/node0 -leader "true" -shards 4 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard0node0.log &
./build/byshard -home $TM_HOME/shard0/node1 -leader "false" -shards 4 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard0node1.log &
./build/byshard -home $TM_HOME/shard0/node2 -leader "false" -shards 4 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard0node2.log &
./build/byshard -home $TM_HOME/shard0/node3 -leader "false" -shards 4 -shardid 0 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard0node3.log &
./build/byshard -home $TM_HOME/shard1/node0 -leader "true" -shards 4 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard1node0.log &
./build/byshard -home $TM_HOME/shard1/node1 -leader "false" -shards 4 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard1node1.log &
./build/byshard -home $TM_HOME/shard1/node2 -leader "false" -shards 4 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard1node2.log &
./build/byshard -home $TM_HOME/shard1/node3 -leader "false" -shards 4 -shardid 1 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard1node3.log &
./build/byshard -home $TM_HOME/shard2/node0 -leader "true" -shards 4 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard2node0.log &
./build/byshard -home $TM_HOME/shard2/node1 -leader "false" -shards 4 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard2node1.log &
./build/byshard -home $TM_HOME/shard2/node2 -leader "false" -shards 4 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard2node2.log &
./build/byshard -home $TM_HOME/shard2/node3 -leader "false" -shards 4 -shardid 2 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard2node3.log &
./build/byshard -home $TM_HOME/shard3/node0 -leader "true" -shards 4 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard3node0.log &
./build/byshard -home $TM_HOME/shard3/node1 -leader "false" -shards 4 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard3node1.log &
./build/byshard -home $TM_HOME/shard3/node2 -leader "false" -shards 4 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard3node2.log &
./build/byshard -home $TM_HOME/shard3/node3 -leader "false" -shards 4 -shardid 3 -inport 10057 -outport "20057,21057,22057,23057" &> $LOG_DIR/shard3node3.log &

echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 byshard
echo "all done"
