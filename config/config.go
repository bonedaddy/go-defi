// Copyright (c) 2021 Alexandre Trottier (bonedaddy). All rights reserved.
// Use of this source code is governed by the Apache 2.0 License that can be found in
// the LICENSE file.

package config

import (
	"io/ioutil"
	"os"

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
	// current values are: infura, direct
	// when direct is provided we connected directly via the RPC parameter
	ProviderType string `yaml:"provider_type"`
	// Network indicates the network type
	// current values are: mainnet, kovan, rinkeby, ropsten
	Network string `yaml:"network"`
	// RPC is the JSON-RPC provider url used when directly connecting to a node
	RPC string `yaml:"rpc" envconfig:"optional"`
	// Websockets when set to true will attempt to use websockets connection negotiation
	Websockets bool `yaml:"websockets" envconfig:"optional"`
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
			ProviderType: "infura",
			Network:      "mainnet",
			RPC:          "http://localhost:8545",
			Websockets:   true,
		},
		Logger: Logger{
			Path:  "go-defi.log",
			Debug: true,
			Dev:   true,
		},
	}
)

// NewConfig generates a new config and stores at path
func NewConfig(path string) error {
	data, err := yaml.Marshal(ExampleConfig)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

// LoadConfig loads the configuration at path
// however if loading from path fails, we attempt to load
// from environment variabels
func LoadConfig(path string) (cfg *Config, err error) {
	func() {
		var r []byte
		r, err = ioutil.ReadFile(path)
		if err != nil {
			return
		}
		var _cfg Config
		if err := yaml.Unmarshal(r, &cfg); err != nil {
			return
		}
		cfg = &_cfg
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
