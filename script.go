package main

import (
	"go-mongo-conf/commands"
	"os"
)

func main() {
	arguments := os.Args[1:]
	commands.RunCommand(arguments)
}
