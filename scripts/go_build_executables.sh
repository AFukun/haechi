GOSRC=$GOPATH/src
ROOT=$GOSRC/github.com/AFukun/haechi

mkdir -p build
go build -o build/example $ROOT/cmd/example
go build -o build/elrond $ROOT/cmd/elrond/coordinator
go build -o build/haechibc $ROOT/cmd/haechi/beacon
go build -o build/haechishard $ROOT/cmd/haechi/shard
chmod +x build/*
