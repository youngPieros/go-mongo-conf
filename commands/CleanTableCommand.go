package commands

import (
	"github.com/youngPieros/go-mongo-conf/enum"
	"github.com/youngPieros/go-mongo-conf/mongo"
)

type CleanTableCommand struct {
	name  enum.CommandType
	table string
}

func CreateCleanTableCommand(table string) *CleanTableCommand {
	return &CleanTableCommand{name: enum.Clean, table: table}
}

func (command *CleanTableCommand) Run() error {
	return mongo.GetDAO().DeleteTable(command.table)
}
