package main

import (
	"fmt"
	"strings"
)

type BarCommand struct {
	mapConfig map[string]string
	stringConfig string
}

func (c *BarCommand) Help() string {
	helpText := `Usage: ... `
	return strings.TrimSpace(helpText)
}

func (c *BarCommand) Run(args []string) int {

	fmt.Println("configs: ", c.mapConfig, c.stringConfig)
	return 0
}

func (c *BarCommand) Synopsis() string {
	return "Command that prints configs"
}
