package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
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

var homeDir, isLeader string
var localPort, remotePort, shardid uint

// var isLeader bool

func init() {
	flag.StringVar(&homeDir, "home", "", "Path to the tendermint config directory (if empty, uses $HOME/.tendermint)")
	flag.StringVar(&isLeader, "leader", "false", "Is it a leader (default: false)")
	flag.UintVar(&localPort, "inport", 12345, "local rpc port")
	flag.UintVar(&remotePort, "outport", 10057, "beacon chain rpc port")
	flag.UintVar(&shardid, "shardid", 0, "shard id")
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
	if isLeader == "true" {
		validatorInterface = hshardnode.NewValidatorInterface(db, uint8(shardid), true, []byte{127, 0, 0, 1}, uint16(localPort), []byte{127, 0, 0, 1}, uint16(remotePort))
	} else if isLeader == "false" {
		validatorInterface = hshardnode.NewValidatorInterface(db, uint8(shardid), false, []byte{127, 0, 0, 1}, uint16(localPort), []byte{127, 0, 0, 1}, uint16(remotePort))
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
