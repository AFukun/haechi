HAECHI_ROOT="$HOME/.haechi"

mkdir -p $HAECHI_ROOT
tendermint testnet  --o $HAECHI_ROOT/shard0 --p2p-port 26600 --starting-ip-address 127.0.0.1
