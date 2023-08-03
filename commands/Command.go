package commands

import (
	"panicmode/tools"
)

func RunCommand(arguments []string) {
	command, err := CreateCommand(arguments)
	if err != nil {
		tools.Logger.Panicw("NOT_ACCEPTABLE_COMMAND_ERROR", "error", err.Error())
	}
	err = command.Run()
	if err != nil {
		tools.Logger.Panicw("EXECUTE_COMMAND_ERROR", "error", err.Error())
	}
	tools.Logger.Info("Command Ran Successfully")
}
