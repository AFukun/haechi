GOSRC=$GOPATH/src
TEST_SCENE="ahl"
TM_HOME="/tmp/$TEST_SCENE"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=30

rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r configs/testnet/* $TM_HOME
echo "configs generated"

pkill -9 $TEST_SCENE
./build/$TEST_SCENE -home $TM_HOME/node0 -node-type "validator" &> $LOG_DIR/node0.log &
./build/$TEST_SCENE -home $TM_HOME/node1 -node-type "validator" &> $LOG_DIR/node1.log &
./build/$TEST_SCENE -home $TM_HOME/node2 -node-type "validator" &> $LOG_DIR/node2.log &
./build/$TEST_SCENE -home $TM_HOME/node3 -node-type "validator" &> $LOG_DIR/node3.log &
echo "ahl launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 TEST_SCENE

echo "all done"
