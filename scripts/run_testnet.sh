TM_HOME="$HOME/.tendermint"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
LOG_DIR="$TM_HOME/log"

rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r configs/testnet/* $TM_HOME

pkill -9 haechi
./build/haechi -home $TM_HOME/node0 &> $LOG_DIR/node0.log &
./build/haechi -home $TM_HOME/node1 &> $LOG_DIR/node1.log &
./build/haechi -home $TM_HOME/node2 &> $LOG_DIR/node2.log &
./build/haechi -home $TM_HOME/node3 &> $LOG_DIR/node3.log &
echo "testnet launched."
sleep 30
pkill -9 haechi
