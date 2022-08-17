package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AFukun/haechi/abci"
	"github.com/spf13/viper"
	abciclient "github.com/tendermint/tendermint/abci/client"
	cfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/types"
)

var tmHome string
var nodeType string
var ipListString string

func init() {
	flag.StringVar(&tmHome, "home", "", "Path to the tendermint config directory (if empty, uses $HOME/.tendermint)")
	flag.StringVar(&nodeType, "node-type", "", "must be \"coordinator\" or \"validator\"")
	flag.StringVar(&ipListString, "ip-list", "", "IP addresses for inter-shard communication (split by \",\")")
}

func main() {
	flag.Parse()
	if tmHome == "" {
		tmHome = os.ExpandEnv("$HOME/.tendermint")
	}
	ipList := strings.Split(ipListString, ",")

	var acc abciclient.Creator
	switch nodeType {
	case "":
		log.Fatalln("Node config: node type must be given")
	case "coordinator":
	case "validator":
		if len(ipList) > 1 {
			log.Fatalln("Node config: validator forward IP count must be less than 1")
		}
		app := abci.NewAhlValidatorApplication(ipList[0])
		acc = abciclient.NewLocalCreator(app)
	default:
		log.Fatalln("Node config: node type must be \"coordinator\" or \"validator\"")
	}

	config := cfg.DefaultValidatorConfig()
	config.SetRoot(tmHome)
	viper.SetConfigFile(fmt.Sprintf("%s/%s", tmHome, "config/config.toml"))
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
