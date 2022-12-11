source ~/.bashrc
GOSRC=$GOPATH/src
ROOT=$GOSRC/github.com/AFukun/haechi

mkdir -p build

go build -o build/ahlbc $ROOT/cmd/ahl/beacon
go build -o build/ahlshard $ROOT/cmd/ahl/shard
go build -o build/ahlclient $ROOT/cmd/ahl/client

go build -o build/byshard $ROOT/cmd/byshard/coordinator
go build -o build/byshardclient $ROOT/cmd/byshard/client

go build -o build/haechibc $ROOT/cmd/haechi/beacon
go build -o build/haechishard $ROOT/cmd/haechi/shard
go build -o build/haechiclient $ROOT/cmd/haechi/client
# go build -o build/haechiclient_haechi_2shard $ROOT/cmd/haechi/client_haechi_2shard
# go build -o build/haechiclient_haechi_4shard $ROOT/cmd/haechi/client_haechi_4shard
# go build -o build/haechiclient_haechi_6shard $ROOT/cmd/haechi/client_haechi_6shard
# go build -o build/haechiclient_haechi_8shard $ROOT/cmd/haechi/client_haechi_8shard
# go build -o build/haechiclient_haechi_10shard $ROOT/cmd/haechi/client_haechi_10shard
# go build -o build/haechiclient_haechi_14shard $ROOT/cmd/haechi/client_haechi_14shard
# go build -o build/haechiclient_haechi_16shard $ROOT/cmd/haechi/client_haechi_16shard

chmod +x build/*
