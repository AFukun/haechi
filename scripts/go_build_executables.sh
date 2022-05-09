ROOT=$GOSRC/github.com/AFukun/haechi

mkdir -p build
go build -o build/haechi $ROOT/cmd/haechi
chmod +x build/*
