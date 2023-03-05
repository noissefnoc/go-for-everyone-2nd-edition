package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// Command for sub command interface
type Command interface {
	Synopsis() string
	Help() string
	Run(args []string) int
}

// Add for add sub command
type AddCommand struct {
	Debug bool
}

func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

func (c *AddCommand) Run(args []string) int {
	var debug bool

	flags := flag.NewFlagSet("add", flag.ContinueOnError)
	flags.BoolVar(&debug, "debug", false, "Run as DEBUG mode")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	fmt.Printf("add sub command option: %T\n", debug)

	return 0
}

func main() {
	c := cli.NewCLI("todo", "0.1.0")

	c.Args = os.Args[1:]

	var debug bool

	flags := flag.NewFlagSet("add", flag.ContinueOnError)
	flags.BoolVar(&debug, "debug", false, "Run as DEBUG mode")

	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{
				Debug: debug,
			}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to execute: %s\n", err.Error())
	}

	os.Exit(exitCode)
}
