#!/bin/bash
TEST_SCENE="ahl"
DURATION=30

TM_HOME="/tmp/$TEST_SCENE"
WORKSPACE="$(go env GOPATH)/src/github.com/AFukun/haechi"
EXEC="$WORKSPACE/build/$TEST_SCENE"
CURRENT_DATE=`date +"%Y-%m-%d-%H-%M"`
LOG_DIR="$WORKSPACE/tmplog/$TEST_SCENE-$CURRENT_DATE"

rm -rf $TM_HOME/*
mkdir -p $TM_HOME

pkill -9 $TEST_SCENE
pkill -9 ${TEST_SCENE}_client
for i in 1
do
    SHARD="shard$i"
    if [ i == 0 ];
    then
        NODE_TYPE="validator"
    else
        NODE_TYPE="validator"
    fi

    mkdir -p $LOG_DIR/$SHARD
    mkdir -p $TM_HOME/$SHARD
    cp -r configs/ahl/$SHARD/* $TM_HOME/$SHARD

    for j in 0 1 2 3
    do
        NODE_HOME="$TM_HOME/$SHARD/node$j"
        LOG_FILE="$LOG_DIR/$SHARD/node$j.log"
        $EXEC -home $NODE_HOME -node-type $NODE_TYPE &> $LOG_FILE &
    done
done
echo "ahl launched"

sleep 3
${EXEC}_client &> $LOG_DIR/$SHARD/client.log &
echo "client launched"

echo "running for ${DURATION}s..."
sleep $DURATION
pkill -9 $TEST_SCENE
pkill -9 ${TEST_SCENE}_client

echo "all done"
