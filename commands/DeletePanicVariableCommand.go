package commands

import (
	"go-mongo-conf/enum"
	"go-mongo-conf/mongo"
)

type DeletePanicVariableCommand struct {
	name     enum.CommandType
	table    string
	variable string
}

func CreateDeletePanicVariableCommand(table, variable string) *DeletePanicVariableCommand {
	return &DeletePanicVariableCommand{name: enum.Delete, table: table, variable: variable}
}

func (command *DeletePanicVariableCommand) Run() error {
	return mongo.GetDAO().DeleteVariable(command.table, command.variable)
}
