GOSRC=$GOPATH/src
TEST_SCENE="haechi"
TM_HOME="$HOME/.haechiOrder"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=60


rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r test/haechi_2shard_4node/* $TM_HOME
echo "configs generated"

pkill -9 haechibc
pkill -9 haechishard
pkill -9 haechiclient_haechi_2shard

./build/haechishard -home $TM_HOME/shard0/node0 -leader "true" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node0.log &
./build/haechishard -home $TM_HOME/shard0/node1 -leader "false" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node1.log &
./build/haechishard -home $TM_HOME/shard0/node2 -leader "false" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node2.log &
./build/haechishard -home $TM_HOME/shard0/node3 -leader "false" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node3.log &
./build/haechishard -home $TM_HOME/shard1/node0 -leader "true" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node0.log &
./build/haechishard -home $TM_HOME/shard1/node1 -leader "false" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node1.log &
./build/haechishard -home $TM_HOME/shard1/node2 -leader "false" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node2.log &
./build/haechishard -home $TM_HOME/shard1/node3 -leader "false" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node3.log &
./build/haechibc -home $TM_HOME/beacon/node0 -leader "true" -shards 2 -inport 10057 -outport "20057,21057" &> $LOG_DIR/beaconnode0.log &
./build/haechibc -home $TM_HOME/beacon/node1 -leader "false" -shards 2 -inport 10157 -outport "20057,21057" &> $LOG_DIR/beaconnode1.log &
./build/haechibc -home $TM_HOME/beacon/node2 -leader "false" -shards 2 -inport 10257 -outport "20057,21057" &> $LOG_DIR/beaconnode2.log &
./build/haechibc -home $TM_HOME/beacon/node3 -leader "false" -shards 2 -inport 10357 -outport "20057,21057" &> $LOG_DIR/beaconnode3.log &

echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 haechibc
pkill -9 haechishard
pkill -9 haechiclient_haechi_2shard
echo "all done"
