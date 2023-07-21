package main

import (
	"flag"
	"fmt"
)

var (
	flags = flag.NewFlagSet("test_install", flag.ExitOnError)
)

func main() {
	println("Hello world!")
	flags.Usage = usage
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix   = `Usage: test_install`
	usageCommands = `Commands: test_install`
)
