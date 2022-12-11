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

	abcchainapp "github.com/AFukun/haechi/consensus/ahl/coordinator/abci"
	abcchainnode "github.com/AFukun/haechi/consensus/ahl/coordinator/validator"
	hctypes "github.com/AFukun/haechi/types"
	abciclient "github.com/tendermint/tendermint/abci/client"
	cfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/types"
)

var homeDir, isLeader, shardPorts, beaconIp, shardIps string
var beaconPort, shardNum uint

// var isLeader bool

func init() {
	flag.StringVar(&homeDir, "home", "", "Path to the tendermint config directory (if empty, uses $HOME/.tendermint)")
	flag.StringVar(&isLeader, "leader", "false", "Is it a leader (default: false)")
	flag.UintVar(&shardNum, "shards", 2, "the number of shards")
	flag.UintVar(&beaconPort, "beaconport", 10057, "beacon chain port")
	flag.StringVar(&shardPorts, "shardports", "20057,21057", "shards chain port")
	flag.StringVar(&beaconIp, "beaconip", "127.0.0.1", "beacon chain ip")
	flag.StringVar(&shardIps, "shardips", "127.0.0.1,127.0.0.1", "shards chain ip")
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
	db := abcchainnode.NewBlockchainState("leveldb", dbPath)
	defer func() {
		if err := db.Database.Close(); err != nil {
			log.Fatalf("Closing database: %v", err)
		}
	}()
	var validatorInterface *abcchainnode.ValidatorInterface
	in_ip_temp := hctypes.HaechiAddress{
		Ip:   hctypes.BytesToIp([]byte(beaconIp)),
		Port: uint16(beaconPort),
	}
	out_ips_temps := make([]hctypes.HaechiAddress, shardNum)
	out_ports_temp := []byte(shardPorts)
	out_ports := bytes.Split(out_ports_temp, []byte(","))
	out_ips_temp := []byte(shardIps)
	out_ips := bytes.Split(out_ips_temp, []byte(","))
	for i, out_port := range out_ports {
		temp_value64, _ := strconv.ParseUint(string(out_port), 10, 64)
		out_ips_temps[i] = hctypes.HaechiAddress{
			Ip:   hctypes.BytesToIp(out_ips[i]),
			Port: uint16(temp_value64),
		}
	}
	// out_ips_temps = out_ips_temps[1:]
	if isLeader == "true" {
		validatorInterface = abcchainnode.NewValidatorInterface(db, uint8(shardNum), true, in_ip_temp, out_ips_temps)
	} else if isLeader == "false" {
		validatorInterface = abcchainnode.NewValidatorInterface(db, uint8(shardNum), false, in_ip_temp, out_ips_temps)
	}

	app := abcchainapp.NewAhlBeaconApplication(validatorInterface)
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
