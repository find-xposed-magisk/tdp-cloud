package parse

import (
	"os"

	"tdp-cloud/cmd/args"
)

type WorkerConfig struct {
	Dataset *args.IDataset `yaml:"dataset"`
	Logger  *args.ILogger  `yaml:"logger"`
	Worker  *args.IWorker  `yaml:"worker"`
}

func (c *WorkerConfig) Read() error {

	c.Dataset = args.Dataset
	c.Logger = args.Logger
	c.Worker = args.Worker

	debug := os.Getenv("TDP_DEBUG")
	args.Debug = debug == "1" || debug == "true"

	if err := readYaml(c); err != nil {
		return err
	}

	initDataset()
	initLogger()

	return nil

}

func (c *WorkerConfig) Save() error {

	return saveYaml(c)

}
