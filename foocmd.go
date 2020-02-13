package main

import (
	"flag"
	"fmt"
	"strings"
)

type FooCommand struct {
	stringValue string
	intValue    int
	boolValue   bool
}

func (c *FooCommand) Help() string {
	helpText := `
Usage: foo [--<optional parameters>=value]
  Simple command that prints argument
Options:
  --string-value=<Just string value>
  --int-value=<Just int value>
  --bool-value=<Just bool value>
`
	return strings.TrimSpace(helpText)
}

func (c *FooCommand) Run(args []string) int {

	flags := flag.NewFlagSet("foo commands", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(c.Help()) }

	flags.StringVar(&c.stringValue, "string-value", "string", "Just string value")
	flags.IntVar(&c.intValue, "int-value", 60, "Just int value")
	flags.BoolVar(&c.boolValue, "bool-value", true, "Just bool value")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	fmt.Println("stringValue: ", c.stringValue)
	fmt.Println("intValue: ", c.intValue)
	fmt.Println("boolValue: ", c.boolValue)

	return 0
}

func (c *FooCommand) Synopsis() string {
	return "Simple command that prints argument"
}
