package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/spf13/viper"

	hshardapp "github.com/AFukun/haechi/consensus/haechi/shard/abci"

	hshardnode "github.com/AFukun/haechi/consensus/haechi/shard/validator"
	abciclient "github.com/tendermint/tendermint/abci/client"
	cfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/types"
)

var homeDir, isLeader, remotePorts string
var localPort, shardNum, shardid uint

// var isLeader bool

func init() {
	flag.StringVar(&homeDir, "home", "", "Path to the tendermint config directory (if empty, uses $HOME/.tendermint)")
	flag.StringVar(&isLeader, "leader", "false", "Is it a leader (default: false)")
	flag.UintVar(&shardNum, "shards", 2, "the number of shards")
	flag.UintVar(&shardid, "shardid", 0, "shard id")
	flag.UintVar(&localPort, "inport", 10057, "beacon chain rpc port")
	flag.StringVar(&remotePorts, "outport", "20057,21057", "shards chain rpc port")
}

func main() {
	flag.Parse()
	if homeDir == "" {
		homeDir = os.ExpandEnv("$HOME/.tendermint")
	}
	config := cfg.DefaultValidatorConfig()

	config.SetRoot(homeDir)

	viper.SetConfigFile(fmt.Sprintf("%s/%s", homeDir, "config/config.toml"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Reading config: %v", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Decoding config: %v", err)
	}
	if err := config.ValidateBasic(); err != nil {
		log.Fatalf("Invalid configuration data: %v", err)
	}
	gf, err := types.GenesisDocFromFile(config.GenesisFile())
	if err != nil {
		log.Fatalf("Loading genesis document: %v", err)
	}

	dbPath := filepath.Join(homeDir, "leveldb")
	db := hshardnode.NewBlockchainState("leveldb", dbPath)
	defer func() {
		if err := db.Database.Close(); err != nil {
			log.Fatalf("Closing database: %v", err)
		}
	}()
	var validatorInterface *hshardnode.ValidatorInterface
	in_ip_temp := hshardnode.HaechiAddress{
		Ip:   []byte{127, 0, 0, 1},
		Port: uint16(localPort),
	}
	out_ips_temps := make([]hshardnode.HaechiAddress, shardNum)
	out_ports_temp := []byte(remotePorts)
	out_ports := bytes.Split(out_ports_temp, []byte(","))
	for i, out_port := range out_ports {
		temp_value64, _ := strconv.ParseUint(string(out_port), 10, 64)
		out_ips_temps[i] = hshardnode.HaechiAddress{
			Ip:   []byte{127, 0, 0, 1},
			Port: uint16(temp_value64),
		}
	}
	if isLeader == "true" {
		validatorInterface = hshardnode.NewValidatorInterface(db, uint8(shardNum), uint8(shardid), true, in_ip_temp, out_ips_temps)
	} else if isLeader == "false" {
		validatorInterface = hshardnode.NewValidatorInterface(db, uint8(shardNum), uint8(shardid), false, in_ip_temp, out_ips_temps)
	}

	app := hshardapp.NewHaechiShardApplication(validatorInterface)
	acc := abciclient.NewLocalCreator(app)

	logger := tmlog.MustNewDefaultLogger(tmlog.LogFormatPlain, tmlog.LogLevelInfo, false)
	node, err := nm.New(config, logger, acc, gf)
	if err != nil {
		log.Fatalf("Creating node: %v", err)
	}

	node.Start()
	defer func() {
		node.Stop()
		node.Wait()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
