package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

var (
	Name    = "Go Cli App"
	Version = "1.0.0"
)

func main() {
	c := cli.NewCLI(Name, Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (command cli.Command, err error) {
			return &FooCommand{}, err
		},
		"bar": barCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

func barCommandFactory() (command cli.Command, err error) {

	mapConfig := map[string]string{
		"name1": "value1",
		"name2": "value2",
	}
	stringConfig := "value3"

	return &BarCommand{
		mapConfig,
		stringConfig,
	}, nil
}
