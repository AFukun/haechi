#!/bin/bash
mkdir -p build
go build -o build/ahl github.com/AFukun/haechi/cmd/ahl/node
chmod +x build/*
