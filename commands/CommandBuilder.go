package commands

import (
	"errors"
	"fmt"
	"go-mongo-conf/enum"
	"go-mongo-conf/mongoconf"
)

func CreateCommand(arguments []string) (CommandInterface, error) {
	commandType, commandArguments := splitCommandArguments(arguments)
	switch commandType {
	case enum.Delete:
		return buildDeleteVariableCommand(commandArguments)
	case enum.Clean:
		return buildCleanTableCommand(commandArguments)
	case enum.Update:
		return buildUpdateVariableCommand(commandArguments)
	case enum.BadCommand:
		return nil, errors.New("bad command")
	default:
		return nil, errors.New("command not found")
	}
}

func splitCommandArguments(arguments []string) (enum.CommandType, []string) {
	if len(arguments) == 0 {
		return enum.BadCommand, []string{}
	}
	commandName := arguments[0]
	return enum.GetCommandTypeFrom(commandName), arguments[1:]
}

func buildDeleteVariableCommand(arguments []string) (CommandInterface, error) {
	if len(arguments) != 2 {
		return nil, errors.New("delete command need two arguments: [table] [variable]")
	}
	return CreateDeletePanicVariableCommand(arguments[0], arguments[1]), nil
}

func buildCleanTableCommand(arguments []string) (CommandInterface, error) {
	if len(arguments) != 1 {
		return nil, errors.New("clean command need one argument: [table]")
	}
	return CreateCleanTableCommand(arguments[0]), nil
}

func buildUpdateVariableCommand(arguments []string) (CommandInterface, error) {
	if len(arguments) != 4 {
		return nil, errors.New("delete command need four arguments: [table] [variable] [type]{one of int,string,float,boolean} [value]")
	}
	variableType := enum.GetVariableTypeFrom(arguments[2])
	if variableType == enum.NONE {
		return nil, errors.New(fmt.Sprintf("wrong variable type. <%s> is not a variable type. this argument should be one of int, string, float and boolean", arguments[2]))
	}
	if err := mongoconf.ValidateTypeOfValue(arguments[3], variableType); err != nil {
		return nil, err
	}
	return CreateUpdateVariableCommand(arguments[0], arguments[1], arguments[3], variableType), nil
}
