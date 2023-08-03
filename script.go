package main

import (
	"github.com/youngPieros/go-mongo-conf/commands"
	"os"
)

func main() {
	arguments := os.Args[1:]
	commands.RunCommand(arguments)
}
