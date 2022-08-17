#!/bin/bash
mkdir -p build
go build -o build/example github.com/AFukun/haechi/cmd/example
chmod +x build/*
