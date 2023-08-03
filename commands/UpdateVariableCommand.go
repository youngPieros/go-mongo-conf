package commands

import (
	"panicmode/enum"
	"panicmode/mongo"
)

type UpdateVariableCommand struct {
	name         enum.VariableType
	table        string
	variable     string
	value        string
	variableType enum.VariableType
}

func CreateUpdateVariableCommand(table, variable, value string, variableType enum.VariableType) *UpdateVariableCommand {
	return &UpdateVariableCommand{table: table, variable: variable, value: value, variableType: variableType}
}

func (command *UpdateVariableCommand) Run() error {
	variables := mongo.GetDAO().Load(command.table)
	for i, variable := range variables {
		if variable.Name == command.variable {
			variables[i].Value = command.value
			return mongo.GetDAO().SaveTable(command.table, variables)
		}
	}
	variables = append(variables, mongo.PanicModeVariable{Name: command.variable, Value: command.value, Type: string(command.variableType)})
	return mongo.GetDAO().SaveTable(command.table, variables)
}
