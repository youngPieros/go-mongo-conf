package mongoconf

import (
	"fmt"
	"strings"
)

type PanicTable struct {
	name      string
	variables map[string]*Variable
}

func CreatePanicTable(name string, vars []*Variable) *PanicTable {
	variables := make(map[string]*Variable)
	for _, variable := range vars {
		variables[variable.GetName()] = variable
	}
	return &PanicTable{name: name, variables: variables}
}

func (table *PanicTable) GetName() string {
	return table.name
}

func (table *PanicTable) String() string {
	var variables []string
	for _, variable := range table.variables {
		variables = append(variables, variable.String())
	}
	return fmt.Sprintf("[<PanicTable> name: <<%s>>, variables:{%s}]", table.name, strings.Join(variables, ","))
}

func (table *PanicTable) GetVariable(name string) *Variable {
	if variable, isAvailable := table.variables[name]; isAvailable {
		return variable
	}
	return nil
}
