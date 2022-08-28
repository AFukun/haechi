GOSRC=$GOPATH/src
TEST_SCENE="elrond"
TM_HOME="$HOME/.haechiElrond"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=10


rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r configs/elrond/* $TM_HOME
echo "configs generated"

pkill -9 elrond
./build/elrond -home $TM_HOME/shard0/node0 -leader "true" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node0.log &
./build/elrond -home $TM_HOME/shard1/node0 -leader "true" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node0.log &
./build/elrond -home $TM_HOME/shard0/node1 -leader "false" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node1.log &
./build/elrond -home $TM_HOME/shard1/node1 -leader "false" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node1.log &


echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 elrond
echo "all done"
