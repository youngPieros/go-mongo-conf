package main

import (
	"os"
	"panicmode/commands"
)

func main() {
	arguments := os.Args[1:]
	commands.RunCommand(arguments)
}
