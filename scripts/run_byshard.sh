source ~/.bashrc
GOSRC=$GOPATH/src
TEST_SCENE="byshard"
TM_HOME="$HOME/.haechibyshard"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=60

IN_IP="127.0.0.1"
OUT_IPS="127.0.0.2,127.0.0.3"

rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r configs/byshard/* $TM_HOME
echo "configs generated"

pkill -9 byshard
./build/byshard -home $TM_HOME/shard0/node0 -leader "true" -shards 2 -shardid 0 -beaconport 10057 -shardports "20057,21057" -beaconip $IN_IP -shardips $OUT_IPS &> $LOG_DIR/shard0node0.log &
./build/byshard -home $TM_HOME/shard1/node0 -leader "true" -shards 2 -shardid 1 -beaconport 10057 -shardports "20057,21057" -beaconip $IN_IP -shardips $OUT_IPS &> $LOG_DIR/shard1node0.log &
./build/byshard -home $TM_HOME/shard0/node1 -leader "false" -shards 2 -shardid 0 -beaconport 10057 -shardports "20057,21057" -beaconip $IN_IP -shardips $OUT_IPS &> $LOG_DIR/shard0node1.log &
./build/byshard -home $TM_HOME/shard1/node1 -leader "false" -shards 2 -shardid 1 -beaconport 10057 -shardports "20057,21057" -beaconip $IN_IP -shardips $OUT_IPS &> $LOG_DIR/shard1node1.log &


echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 byshard
echo "all done"
