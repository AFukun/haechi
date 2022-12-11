source ~/.bashrc
GOSRC=$GOPATH/src
TEST_SCENE="haechi"
TM_HOME="$HOME/.haechiOrder"
WORKSPACE="$GOSRC/github.com/AFukun/haechi"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"
DURATION=60

SHARD_NUM=2
SHARD_SIZE=2
BEACON_PORT=10057
BEACON_IP="127.0.0.1"
SHARD_PORTS="20057,21057"
SHARD_IPS="127.0.0.2,127.0.0.3"

while getopts ":n:m:p:i:s:x:" opt
do 
    case $opt in
    n) # shard number
        echo "shard number is $OPTARG"
        SHARD_NUM=$OPTARG
        ;;
    m) # shard number
        echo "shard size is $OPTARG"
        SHARD_SIZE=$OPTARG
        ;;  
    p) # beacon port
        echo "beaconport is $OPTARG"
        BEACON_PORT=$OPTARG
        ;;  
    i) # beacon ip
        echo "beaconip is $OPTARG"
        BEACON_IP=$OPTARG
        ;;  
    s) # shard ports
        echo "shardports is $OPTARG"
        SHARD_PORTS=$OPTARG
        ;;  
    x) # shard ips
        echo "shardips is $OPTARG"  
        SHARD_IPS=$OPTARG
        ;;  
    ?)  
        echo "unknown: $OPTARG"
        ;;
    esac
done

rm -rf $TM_HOME/*
mkdir -p $TM_HOME
mkdir -p $LOG_DIR

cp -r configs/haechi/* $TM_HOME
echo "configs generated"

pkill -9 haechibc
pkill -9 haechishard

# run beacon node
for ((i=0;i<$SHARD_SIZE;i++))
do
    if [ $i -eq 0 ]; then
        echo "running beacon leader"
        ./build/haechibc -home $TM_HOME/beacon/node$i -leader "true" -shards $SHARD_NUM -beaconport $BEACON_PORT -shardports $SHARD_PORTS -beaconip $BEACON_IP -shardips $SHARD_IPS &> $LOG_DIR/beaconnode$i.log &
    else
        echo "running beacon validator"
        ./build/haechibc -home $TM_HOME/beacon/node$i -leader "false" -shards $SHARD_NUM -beaconport $BEACON_PORT -shardports $SHARD_PORTS -beaconip $BEACON_IP -shardips $SHARD_IPS &> $LOG_DIR/beaconnode$i.log &
    fi
    sleep 1
done

# run shard node
for ((j=0;j<$SHARD_NUM;j++))
do
    for ((k=0;k<$SHARD_SIZE;k++))
    do
        if [ $k -eq 0 ]; then
            echo "running shard$j leader"
            ./build/haechishard -home $TM_HOME/shard$j/node$k -leader "true" -shards $SHARD_NUM -shardid $j -beaconport $BEACON_PORT -shardports $SHARD_PORTS -beaconip $BEACON_IP -shardips $SHARD_IPS &> $LOG_DIR/shard$j-node$k.log &
        else
            echo "running shard$j validator$k"
            ./build/haechishard -home $TM_HOME/shard$j/node$k -leader "false" -shards $SHARD_NUM -shardid $j -beaconport $BEACON_PORT -shardports $SHARD_PORTS -beaconip $BEACON_IP -shardips $SHARD_IPS &> $LOG_DIR/shard$j-node$k.log &
        fi
    sleep 1
    done
done

# ./build/haechishard -home $TM_HOME/shard0/node0 -leader "true" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node0.log &
# ./build/haechishard -home $TM_HOME/shard1/node0 -leader "true" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node0.log &
# ./build/haechishard -home $TM_HOME/shard0/node1 -leader "false" -shards 2 -shardid 0 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard0node1.log &
# ./build/haechishard -home $TM_HOME/shard1/node1 -leader "false" -shards 2 -shardid 1 -inport 10057 -outport "20057,21057" &> $LOG_DIR/shard1node1.log &
# ./build/haechibc -home $TM_HOME/beacon/node0 -leader "true" -shards 2 -inport 10057 -outport "20057,21057" &> $LOG_DIR/beaconnode0.log &
# ./build/haechibc -home $TM_HOME/beacon/node1 -leader "false" -shards 2 -inport 10157 -outport "20057,21057" &> $LOG_DIR/beaconnode1.log &


echo "testnet launched"
echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 haechibc
pkill -9 haechishard
echo "all done"
