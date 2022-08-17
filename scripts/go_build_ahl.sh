#!/bin/bash
mkdir -p build
go build -o build/ahl github.com/AFukun/haechi/cmd/ahl/node
go build -o build/ahl_client github.com/AFukun/haechi/cmd/ahl/client
chmod +x build/*
