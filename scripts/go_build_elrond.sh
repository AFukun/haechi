#!/bin/bash
mkdir -p build
go build -o build/elrond github.com/AFukun/haechi/cmd/elrond/coordinator
chmod +x build/*
