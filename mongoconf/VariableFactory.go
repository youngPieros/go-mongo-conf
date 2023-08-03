package mongoconf

import (
	"github.com/youngPieros/go-mongo-conf/enum"
	"strconv"
)

func createVariable(name string, variableType enum.VariableType, defaultValue string) *Variable {
	return &Variable{name: name, variableType: variableType, defaultValue: defaultValue, value: defaultValue}
}

func FromInteger(name string, value int64) *Variable {
	return createVariable(name, enum.INTEGER, strconv.FormatInt(value, 10))
}

func FromBoolean(name string, value bool) *Variable {
	return createVariable(name, enum.BOOLEAN, strconv.FormatBool(value))
}

func FromString(name string, value string) *Variable {
	return createVariable(name, enum.STRING, value)
}

func FromFloat(name string, value float64) *Variable {
	return createVariable(name, enum.FLOAT, strconv.FormatFloat(value, 'f', -1, 64))
}
