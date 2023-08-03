package commands

import (
	"panicmode/enum"
	"panicmode/mongo"
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
