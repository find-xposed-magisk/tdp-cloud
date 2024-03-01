package subset

import (
	"flag"

	"tdp-cloud/cmd/parse"
	"tdp-cloud/service"
)

func serverFlag() *FlagSet {

	var action string

	command := &FlagSet{
		FlagSet: flag.NewFlagSet("server", flag.ExitOnError),
		Comment: "TDP Cloud Server Management",
		Execute: func() {
			serverExec(action)
		},
	}

	command.StringVar(&action, "s", "", "management server service")
	command.StringVar(&parse.YamlFile, "c", "config.yml", "config file path")

	return command

}

func serverExec(act string) {

	c := &parse.ServerConfig{}
	c.Read()

	if act == "" || act == "start" {
		c.Save()
	}

	service.Control("server", act)

}
