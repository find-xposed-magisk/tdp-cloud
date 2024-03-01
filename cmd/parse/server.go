package parse

import (
	"os"

	"github.com/opentdp/go-helper/strutil"

	"tdp-cloud/cmd/args"
)

type ServerConfig struct {
	Dataset  *args.IDataset  `yaml:"dataset"`
	Database *args.IDatabase `yaml:"database"`
	Logger   *args.ILogger   `yaml:"logger"`
	Server   *args.IServer   `yaml:"server"`
}

func (c *ServerConfig) Read() error {

	c.Dataset = args.Dataset
	c.Database = args.Database
	c.Logger = args.Logger
	c.Server = args.Server

	debug := os.Getenv("TDP_DEBUG")
	args.Debug = debug == "1" || debug == "true"

	if err := readYaml(c); err != nil {
		return err
	}

	initDataset()
	initLogger()

	if c.Database.Type == "sqlite" {
		c.Database.Host = c.Dataset.Dir
	}

	if c.Server.JwtKey == "" {
		c.Server.JwtKey = strutil.Rand(32)
	}

	return nil

}

func (c *ServerConfig) Save() error {

	return saveYaml(c)

}
