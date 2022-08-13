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

	elrondapp "github.com/AFukun/haechi/consensus/elrond/coordinator/abci"

	elrondnode "github.com/AFukun/haechi/consensus/elrond/coordinator/validator"
	abciclient "github.com/tendermint/tendermint/abci/client"
	cfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/types"
)

var homeDir, isLeader string
var localPort, remotePort uint

// var isLeader bool

func init() {
	flag.StringVar(&homeDir, "home", "", "Path to the tendermint config directory (if empty, uses $HOME/.tendermint)")
	flag.StringVar(&isLeader, "leader", "false", "Is it a leader (default: false)")
	flag.UintVar(&localPort, "inport", 12345, "local rpc port")
	flag.UintVar(&remotePort, "outport", 12346, "remote rpc port")
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
	// fmt.Println("path database is: " + dbPath)
	db := elrondnode.NewBlockchainState("leveldb", dbPath)
	// db, err := dbm.NewGoLevelDBWithOpts
	// db := dbm.NewMemDB()
	// db, err := badger.Open(badger.DefaultOptions(dbPath))
	// if err != nil {
	// 	log.Fatalf("Opening database: %v", err)
	// }
	defer func() {
		if err := db.Database.Close(); err != nil {
			log.Fatalf("Closing database: %v", err)
		}
	}()
	var validatorInterface *elrondnode.ValidatorInterface
	if isLeader == "true" {
		validatorInterface = elrondnode.NewValidatorInterface(db, true, []byte{127, 0, 0, 1}, uint16(localPort), []byte{127, 0, 0, 1}, uint16(remotePort))
	} else if isLeader == "false" {
		validatorInterface = elrondnode.NewValidatorInterface(db, false, []byte{127, 0, 0, 1}, uint16(localPort), []byte{127, 0, 0, 1}, uint16(remotePort))
	}

	app := elrondapp.NewElrondApplication(validatorInterface)
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
