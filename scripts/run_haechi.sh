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

cp -r configs/haechi/* $TM_HOME
echo "configs generated"

pkill -9 haechibc
./build/haechibc -home $TM_HOME/beacon/node0 -leader "true" -shards 2 -inport 10057 -outport "20057, 21057" &> $LOG_DIR/beaconnode0.log &
./build/elrond -home $TM_HOME/beacon/node0 -leader "false" -shards 2 -inport 10157 -outport "20057, 21057" &> $LOG_DIR/beaconnode1.log &

pkill -9 haechishard
./build/haechishard -home $TM_HOME/shard0/node0 -leader "true" -inport 10057 -outport 20057 &> $LOG_DIR/shard0node0.log &
./build/haechishard -home $TM_HOME/shard1/node0 -leader "true" -inport 10057 -outport 21057 &> $LOG_DIR/shard1node0.log &
./build/haechishard -home $TM_HOME/shard0/node1 -leader "false" -inport 10057 -outport 20057 &> $LOG_DIR/shard0node1.log &
./build/haechishard -home $TM_HOME/shard1/node1 -leader "false" -inport 10057 -outport 21057 &> $LOG_DIR/shard1node1.log &


echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 haechibc
pkill -9 haechishard
echo "all done"
