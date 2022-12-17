source ~/.bashrc
GOSRC=$GOPATH/src
ROOT=$GOSRC/github.com/AFukun/haechi

mkdir -p build

go build -o build/ahlbc $ROOT/cmd/ahl/beacon
go build -o build/ahlshard $ROOT/cmd/ahl/shard
go build -o build/ahlclient $ROOT/cmd/latency-client/ahl
go build -o build/ahlclient $ROOT/cmd/user-client/ahl

go build -o build/byshard $ROOT/cmd/byshard/coordinator
go build -o build/byshardclient $ROOT/cmd/latency-client/byshard
go build -o build/byshardclient $ROOT/cmd/user-client/byshard

go build -o build/haechibc $ROOT/cmd/haechi/beacon
go build -o build/haechishard $ROOT/cmd/haechi/shard
go build -o build/haechiclient $ROOT/cmd/latency-client/haechi
go build -o build/haechiclient $ROOT/cmd/user-client/haechi

chmod +x build/*
