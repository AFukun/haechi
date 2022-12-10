source ~/.bashrc
GOSRC=$GOPATH/src
TEST_SCENE="example"
TM_HOME="$HOME/.example"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=30


rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r configs/testnet/* $TM_HOME
echo "configs generated"

pkill -9 example
./build/example -home $TM_HOME/node0 &> $LOG_DIR/node0.log &
./build/example -home $TM_HOME/node1 &> $LOG_DIR/node1.log &
./build/example -home $TM_HOME/node2 &> $LOG_DIR/node2.log &
./build/example -home $TM_HOME/node3 &> $LOG_DIR/node3.log &
echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 example
echo "all done"
