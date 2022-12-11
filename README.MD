# Haechi

## build

```bash
$ make build
```

## test

### node parameter description

| node parameters | description        | default               |
|-----------------|--------------------|-----------------------|
| n               | shard number       | 2                     |
| m               | shard size         | 2                     |
| p               | beacon chain port  | 10057                 |
| i               | beacon chain ip    | "127.0.0.1"           |
| s               | shard chain points | "20057,21057"         |
| x               | shard chain ips    | "127.0.0.1,127.0.0.1" |

### client parameter description

| client parameters | description            | default               |
|-------------------|------------------------|-----------------------|
| shards            | shard number           | 2                     |
| beaconport        | beacon chain port      | 10057                 |
| beaconip          | beacon chain ip        | "127.0.0.1"           |
| shardports        | shard chain points     | "20057,21057"         |
| shardips          | shard chain ips        | "127.0.0.2,127.0.0.3" |
| batch             | batch size per request | 10                    |
| rate              | sleeping time, ms      | 100                   |
| ratio             | cross-shard txs ratio  | 0.8                   |

1. ahl test

- run node

```bash
./scripts/run_ahl.sh -n 2 -m 2 -p 10057 -i "127.0.0.1" -s "20057,21057" -x "127.0.0.1,127.0.0.1"
```

- run client

```bash
./build/ahlclient -shards 2 -beaconport 10057 -beaconip "127.0.0.1" -shardports "20057,21057" -shardips "127.0.0.1,127.0.0.1" -batch 10 -rate 100 -ratio 0.8
```

2. byshard test

- run node

```bash
./scripts/run_byshard.sh -n 2 -m 2 -p 10057 -i "127.0.0.1" -s "20057,21057" -x "127.0.0.2,127.0.0.3"
```

- run client

```bash
./build/byshardclient -shards 2 -beaconport 10057 -beaconip "127.0.0.1" -shardports "20057,21057" -shardips "127.0.0.1,127.0.0.1" -batch 10 -rate 100 -ratio 0.8
```

3. haechi test

- run node

```bash
./scripts/run_haechi.sh -n 2 -m 2 -p 10057 -i "127.0.0.1" -s "20057,21057" -x "127.0.0.2,127.0.0.3"
```

- run client

```bash
./build/haechiclient -shards 2 -beaconport 10057 -beaconip "127.0.0.1" -shardports "20057,21057" -shardips "127.0.0.1,127.0.0.1" -batch 10 -rate 100 -ratio 0.8
```
