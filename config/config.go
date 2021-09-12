// Copyright (c) 2021 Alexandre Trottier (bonedaddy). All rights reserved.
// Use of this source code is governed by the Apache 2.0 License that can be found in
// the LICENSE file.

package config

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/bonedaddy/go-defi/testenv"
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	config "github.com/ipfs/go-ipfs-config"
	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
	"go.bobheadxi.dev/zapx/zapx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config provides the base configuration type
type Config struct {
	Database   `yaml:"database" envconfig:"optional"`
	Blockchain `yaml:"blockchain" envconfig:"optional"`
	Logger     `yaml:"logger"`
}

// Blockchain provides configuration for connection to a blockchain rpc provider
type Blockchain struct {
	// APIKey provides an API key to be used for authenticaiton with an rpc provider
	APIKey string `yaml:"api_key" envconfig:"optional"`
	// ProvderType indicates the type of rpc provider
	// current values are: infura, direct, simulated
	// when direct is provided we connected directly via the RPC parameter
	// when simulated is provided we create an in-memory blockchain
	ProviderType string `yaml:"provider_type"`
	// Network indicates the network type
	// current values are: mainnet, kovan, rinkeby, ropsten
	Network string `yaml:"network"`
	// RPC is the JSON-RPC provider url used when directly connecting to a node
	RPC string `yaml:"rpc" envconfig:"optional"`
	// Websockets when set to true will attempt to use websockets connection negotiation
	Websockets bool `yaml:"websockets" envconfig:"optional"`
	// Account provides configuration for transaction signing and gas price estimation
	Account `yaml:"account"`
}

// Account provides configuration over the account used to sign transactions
// from https://github.com/indexed-finance/circuit-breaker/blob/master/config/config.go
// copyright for this code can be found in the LICENSE file of indexed-finance/circuit-breaker
type Account struct {
	// Mode specifies the ethereum account mode to use
	// currently supported values: keyfile, privatekey
	Mode            string `yaml:"mode"`
	KeyFilePath     string `yaml:"key_file_path"`
	KeyFilePassword string `yaml:"key_file_password"`
	PrivateKey      string `yaml:"private_key"`
	GasPrice        `yaml:"gas_price"`
}

// GasPrice controls how we specify gas prices for sending transactions
// from https://github.com/indexed-finance/circuit-breaker/blob/master/config/config.go
// copyright for this code can be found in the LICENSE file of indexed-finance/circuit-breaker
type GasPrice struct {
	// reports the minimum amount of gwei we will spend
	// if the gas price oracle from go-ethereum reports less than this
	// we override the gas price to this
	MinimumGwei string `yaml:"minimum_gwei"`
	// Multiplier controls the gas price multiplier
	// whatever the minimum gwei, or the value reported by the gasprice oracle
	// we multiply it by the value specified here. It defaults to three and should only
	// be changed with care as decreasing the multiplier could result in failure to get
	// next block transaction inclusion
	Multiplier string `yaml:"multiplier"`
}

// Database provides configuration over our database connection
type Database struct {
	// Type specifies the database type to be used
	// currently supports: sqlite, postgres
	// when using sqlite type all other options except DBName and DBPath are ignore
	Type string `yaml:"type"`
	// Host specifies the database host address
	Host string `yaml:"host"  envconfig:"optional"`
	// Port specifies the port the database is exposed on
	Port string `yaml:"port"  envconfig:"optional"`
	// User specifies the username to use for authentication
	User string `yaml:"user"  envconfig:"optional"`
	// Pass specifies the password to use for authentication
	Pass string `yaml:"pass"  envconfig:"optional"`
	// DBName specifies the name of the database
	DBName string `yaml:"db_name"`
	// DBPath specifies the path to the database file (only applicable to sqlite)
	DBPath string `yaml:"db_path"  envconfig:"optional"`
	// SSLModeDisable specifies whether or not to disable ssl connection layer security
	SSLModeDisable bool `yaml:"ssl_mode_disable"  envconfig:"optional"`
}

// Logger provides configuration over zap logger
type Logger struct {
	// Path to store the configuration file
	Path string `yaml:"path" envconfig:"default=go-defi.log"`
	// Debug enables displaying of debug logs
	Debug bool `yaml:"debug" envconfig:"optional"`
	// Dev enables foramtting of logs to be more human readable
	Dev bool `yaml:"dev" envconfig:"optional"`
	// FileOnly disables stdout logging and will only display log output in `Path`
	FileOnly bool `yaml:"file_only" envconfig:"optional"`
	// Fields provides optional fields to use for logging metadata
	Fields map[string]interface{} `yaml:"fields" envconfig:"optional"`
}

var (
	// ExampleConfig is primarily used to provide a template for generating the config file
	ExampleConfig = &Config{
		Database: Database{
			Type:           "sqlite",
			Host:           "localhost",
			Port:           "5432",
			User:           "postgres",
			Pass:           "password123",
			DBName:         "go-defi",
			DBPath:         "",
			SSLModeDisable: false,
		},
		Blockchain: Blockchain{
			APIKey:       "CHANGEME",
			ProviderType: "simulated",
			Network:      "mainnet",
			RPC:          "http://localhost:8545",
			Websockets:   true,
			Account: Account{
				Mode:            "privatekey",
				KeyFilePath:     "CHANGEME-PATH",
				KeyFilePassword: "CHANGEME-PASS",
				PrivateKey:      "CHANGEME-PK",
				GasPrice: GasPrice{
					MinimumGwei: ToWei("100.0", 9).String(), // 9 is the denomination of gwei
					Multiplier:  "3",
				},
			},
		},
		Logger: Logger{
			Path:  "go-defi.log",
			Debug: true,
			Dev:   true,
		},
	}
)

// NewConfig generates a new config and stores at path
// TODO(bonedaddy): generate a random private key
func NewConfig(path string) error {
	data, err := yaml.Marshal(ExampleConfig)
	if err != nil {
		return errors.Wrap(err, "marshal")
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

// LoadConfig loads the configuration at path
// however if loading from path fails, we attempt to load
// from environment variabels
func LoadConfig(path string) (cfg *Config, err error) {
	cfg = new(Config)
	func() {
		var r []byte
		r, err = ioutil.ReadFile(path)
		if err != nil {
			return
		}
		if err := yaml.Unmarshal(r, cfg); err != nil {
			return
		}
	}()
	if err != nil {
		cfg = &Config{}
		err = envconfig.Init(cfg)
		return
	}
	return
}

// ZapLogger parses the Logger configuration struct and returns a zap logger instance
func (c *Config) ZapLogger() (*zap.Logger, error) {
	var opts []zapx.Option
	if c.Logger.Debug {
		opts = append(opts, zapx.WithDebug(true))
	}
	if c.Logger.FileOnly {
		opts = append(opts, zapx.OnlyToFile())
	}
	if c.Logger.Fields != nil {
		opts = append(opts, zapx.WithFields(c.Logger.Fields))
	}
	return zapx.New(
		c.Logger.Path,
		c.Logger.Dev,
		opts...,
	)
}

// EthClient parses the Blockchain configuration struct and returns a blockchain interface
func (c *Config) EthClient(ctx context.Context) (utils.Blockchain, error) {
	var (
		ec  *ethclient.Client
		err error
	)
	switch c.Blockchain.ProviderType {
	case "simulated":
		return testenv.NewBlockchain(ctx)
	case "infura":
		transport := ""
		path := ""
		if c.Blockchain.Websockets {
			transport = "wss://"
			path = "/ws/v3"
		} else {
			transport = "http://"
			path = "/v3"
		}
		url := transport + c.Blockchain.Network + ".infura.io" + path + "/" + c.Blockchain.APIKey
		ec, err = ethclient.DialContext(ctx, url)
	case "direct":
		ec, err = ethclient.DialContext(ctx, c.Blockchain.RPC)
	}
	return ec, err
}

// Authorizer parses the Account configuration struct to return a transaction signer
// from https://github.com/indexed-finance/circuit-breaker/blob/master/cmd/services_run.go
// copyright for this code can be found in the LICENSE file of indexed-finance/circuit-breaker
func (c *Config) Authorizer(cfg *config.Config) (*utils.Authorizer, error) {
	switch c.Blockchain.Account.Mode {
	case "keyfile":
		return utils.NewAuthorizer(c.Blockchain.Account.KeyFilePath, c.Blockchain.Account.KeyFilePassword)
	case "privatekey":
		pk, err := crypto.HexToECDSA(c.Blockchain.Account.PrivateKey)
		if err != nil {
			return nil, err
		}
		return utils.NewAuthorizerFromPK(pk), nil
	default:
		return nil, errors.New("unsupported account mode, must be one of: keyfile, privatekey")
	}
}
