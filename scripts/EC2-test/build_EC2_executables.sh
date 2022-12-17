source ~/.bashrc
GOSRC=$GOPATH/src
ROOT=$GOSRC/github.com/AFukun/haechi

mkdir -p build

go build -o build/ahlbc $ROOT/cmd/ahl/beacon
go build -o build/ahlshard $ROOT/cmd/ahl/shard
go build -o build/ahllatency $ROOT/cmd/latency-client/ahl
go build -o build/ahluser $ROOT/cmd/user-client/ahl

go build -o build/byshard $ROOT/cmd/byshard/coordinator
go build -o build/byshardlatency $ROOT/cmd/latency-client/byshard
go build -o build/bysharduser $ROOT/cmd/user-client/byshard

go build -o build/haechibc $ROOT/cmd/haechi/beacon
go build -o build/haechishard $ROOT/cmd/haechi/shard
go build -o build/haechilatency $ROOT/cmd/latency-client/haechi
go build -o build/haechiuser $ROOT/cmd/user-client/haechi

chmod +x build/*
