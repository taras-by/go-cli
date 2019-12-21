package main

import (
	"fmt"
	"strings"
)

type FooCommand struct {
	SomeProperty string
}

func (c *FooCommand) Help() string {
	helpText := `Usage: ... `
	return strings.TrimSpace(helpText)
}

func (c *FooCommand) Run(args []string) int {

	fmt.Println("SomeProperty: ", c.SomeProperty)
	return 0
}

func (c *FooCommand) Synopsis() string {
	return "Just foo command"
}
