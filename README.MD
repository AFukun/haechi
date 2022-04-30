# Haechi testing descriptions

## 1. Running a test chain with multiple validators
### (1) Using tendermint to create node files 
execute
`tendermint testnet --v [val_num] --o [cfg_path] --populate-persistent-peers --starting-ip-address 127.0.0.1`

This will create **[val_num]** node files (including config, data) in path **[cfg_path]**

*Example: Use `tendermint testnet --v 2 --o $HAECHIROOT/shard0 --populate-persistent-peers --starting-ip-address 127.0.0.1`* will create file *HAECHIROOT/shard0/node0* and *HAECHIROOT/shard0/node1*, along with the config file and data file. The genesis file as well as connection information among nodes will be configured automatically.

### (2) Modify config file to enable communications for nodes
Modify parameters *Ports* in file *config.toml* to avoid conflicts, refer to file *$HAECHIROOT/shard0*

### (3) Run haechi with the command
execute `go run haechi --home=[node file]` to respectively start {node0, node1, ..., node[val_num-1]}