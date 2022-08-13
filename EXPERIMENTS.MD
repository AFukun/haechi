# 1 根据实验创建配置文件
*文件命名格式可参考./config文件夹下的文件目录*

## (1) 生成配置文件，运行指令：
```
$ tendermint testnet --v 单个分片的节点数量 --o /Users/gudale/go/src/github.com/AFukun/haechi/configs/文件名 --populate-persistent-peers --starting-ip-address 127.0.0.1
```
## (2) 修改配置文件：
几个要注意的点：
- 端口号第一位确定分片/协调者身份；第二位为分片id（分片号从0开始）；第三位为节点id（从0开始）
- “persistent-peers”参数的端口号为其他节点的监听端口号，且ip地址需要统一改成“127.0.0.1”；
- “indexer”参数的默认值“null”需改为“kv”；
- “create-empty-blocks”参数改为“false”，表示不创建空块

# 2 编写测试脚本
在./scripts文件中实现运行节点的测试脚本，在Makefile文件中添加快速执行脚本的指令

# 3 实现客户端发送交易请求
在./cmd相应测试文件夹中编写main.go，在main.go中实现发交易请求功能

**注意：每个跨分片共识算法的交易格式是不一样的，具体可参考当前的“./cmd/算法/client/main.go”里面的指令噢**

# 4 性能计算
在./tmplog/文件名中查看日志，计算相关指标

# 5 一些有用的网站参考（欢迎补充）
- tendermint介绍：https://docs.tendermint.com/v0.35/