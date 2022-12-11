# Haechi

## build

```bash
$ make build
```

## test

| parameters | description        | default               |
|------------|--------------------|-----------------------|
| n          | shard number       | 2                     |
| m          | shard size         | 2                     |
| p          | beacon chain port  | 10057                 |
| i          | beacon chain ip    | "127.0.0.1"           |
| s          | shard chain points | "20057,21057"         |
| x          | shard chain ips    | "127.0.0.2,127.0.0.3" |

1. ahl test

```bash
./scripts/run_ahl.sh -n 2 -m 2 -p 10057 -i "127.0.0.1" -s "20057,21057" -x "127.0.0.2,127.0.0.3"
```

2. byshard test

```bash
./scripts/run_byshard.sh -n 2 -m 2 -p 10057 -i "127.0.0.1" -s "20057,21057" -x "127.0.0.2,127.0.0.3"
```

3. haechi test

```bash
./scripts/run_haechi.sh -n 2 -m 2 -p 10057 -i "127.0.0.1" -s "20057,21057" -x "127.0.0.2,127.0.0.3"
```
