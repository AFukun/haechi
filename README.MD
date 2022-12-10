# Haechi

## build

```bash
$ make build
```

## run testnet

```bash
$ make run_testnet
```

# Experiment_1

## heachi: 16shard4node

调整cmd/haechi/client_haechi_16shard中main.go的cross_rate，从[0.1 ... 0.9]

```zsh
make build
make run_haechi_16shard_4node_test
```

在build目录下输入

```
./haechiclient_haechi_16shard_4node
```

开始输入交易

## ahl: 16shard4node

调整cmd/ahl/client_ahl_16shard中main.go的cross_rate，从[0.1 ... 0.9]

```zsh
make build
make run_ahl_16shard_4node_test
```

在build目录下输入

```zsh
./ahlclient_ahl_16shard_4node
```

开始输入交易

## erlond: 16shard4node

调整cmd/erlond/client_erlond_16shard中main.go的cross_rate，从[0.1 ... 0.9]

每次调整cross_rate 都需要重新make build 重新编译main.go

```zsh
make build
make run_erlond_16shard_4node_test
```

在build目录下输入

```zsh
./erlondclient_byshard_16shard_4node
```

开始输入交易

# Experiment_2

## heachi

调整cmd/haechi/client_haechi_%dshard中main.go的cross_rate为0.9

%d对应shard的数目，如client_haechi_2shard, client_haechi_4shard

```zsh
make build
# 不同的shard数目，对应不同的%d，从2到16
make run_haechi_%dshard_4ode_test
```

在build目录下输入

```zsh
# 不同的shard数目，对应不同的%d，从2到16
./haechiclient_haechi_%dshard_4node
```

开始输入交易

## ahl

调整cmd/ahl/client_ahl_%dshard中main.go的cross_rate为0.9

%d对应shard的数目，如client_ahl_2shard, client_ahl_4shard

```zsh
make build
# 不同的shard数目，对应不同的%d，从2到16
make run_ahl_%dshard_4ode_test
```

转到build目录下输入

```zsh
# 不同的shard数目，对应不同的%d，从2到16
./ahlclient_ahl_%dshard_4node
```

开始输入交易

## erlond

调整cmd/erlond/client_byshard_%dshard中main.go的cross_rate为0.9

%d对应shard的数目，如client_byshard_2shard, client_byshard_4shard

```zsh
make build
# 不同的shard数目，对应不同的%d，从2到16
make run_byshard_%dshard_4ode_test
```

在build目录下输入

```zsh
# 不同的shard数目，对应不同的%d，从2到16
./byshardclient_erlond_%dshard_4node
```

开始输入交易

# Experiment_3

## heachi

调整cmd/haechi/client_haechi_16shard中main.go的cross_rate为0.9

```zsh
make build
# 不同的node数目，对应不同的%d，从4到20
make run_haechi_16shard_%dnode_test
```

在build目录下输入

```zsh
# 不同的node数目，对应不同的%d，从4到20
./haechiclient_haechi_16shard_%dnode
```

开始输入交易

## ahl

调整cmd/ahl/client_ahl_16shard中main.go的cross_rate为0.9

```zsh
make build
# 不同的node数目，对应不同的%d，从4到20
make run_ahl_16shard_4ode_test
```

转到build目录下输入

```zsh
# 不同的node数目，对应不同的%d，从4到20
./ahlclient_ahl_16shard_%dnode
```

开始输入交易

## erlond

调整cmd/erlond/client_byshard_16shard中main.go的cross_rate为0.9

```zsh
make build
# 不同的node数目，对应不同的%d，从4到20
make run_byshard_16shard_%dnode_test
```

在build目录下输入

```zsh
# 不同的node数目，对应不同的%d，从4到20
./byshardclient_erlond_16shard_%dnode
```

开始输入交易

# Experiment_4

## heachi

调整cmd/haechi/client_haechi_16shard中main.go的cross_rate为0.9

```zsh
make build
# 不同的beacon数目，对应不同的%d，从4到20
make run_haechi_16shard_4node_%dbeacon
```

在build目录下输入

```zsh
# 不同的beacon数目，对应不同的%d，从4到20
./haechiclient_haechi_16shard_4node_%dbeacon
```

开始输入交易