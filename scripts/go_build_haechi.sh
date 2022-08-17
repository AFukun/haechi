#!/bin/bash
mkdir -p build
go build -o build/haechibc github.com/AFukun/haechi/cmd/haechi/beacon
go build -o build/haechishard github.com/AFukun/haechi/cmd/haechi/shard
chmod +x build/*
