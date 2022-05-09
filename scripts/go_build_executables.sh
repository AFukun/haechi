ROOT=$GOSRC/github.com/AFukun/haechi

mkdir -p build
go build -o build/example $ROOT/cmd/example
chmod +x build/*
