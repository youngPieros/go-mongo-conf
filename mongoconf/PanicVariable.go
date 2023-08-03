package mongoconf

import (
	"errors"
	"fmt"
	"go-mongo-conf/enum"
	"go-mongo-conf/tools"
	"strconv"
)

type Variable struct {
	name         string
	variableType enum.VariableType
	value        string
	defaultValue string
}

func (variable *Variable) GetName() string {
	return variable.name
}

func (variable *Variable) GetDefaultValue() string {
	return variable.defaultValue
}

func (variable *Variable) GetVariableType() enum.VariableType {
	return variable.variableType
}

func (variable *Variable) GetInteger() int64 {
	intValue, _ := strconv.ParseInt(variable.value, 0, 64)
	return intValue
}

func (variable *Variable) GetString() string {
	return variable.value
}

func (variable *Variable) GetFloat() float64 {
	floatValue, _ := strconv.ParseFloat(variable.value, 64)
	return floatValue
}

func (variable *Variable) GetBoolean() bool {
	if variable.value == "true" || variable.value == "True" || variable.value == "TRUE" {
		return true
	}
	if variable.value == "false" || variable.value == "False" || variable.value == "FALSE" {
		return false
	}
	return false
}

func (variable *Variable) SetValue(value string) error {
	if variable.variableType == enum.INTEGER && !tools.IsInteger(value) {
		return errors.New(fmt.Sprintf("BAD_VALUE, <%s> is not <%s>", value, string(enum.INTEGER)))
	}
	if variable.variableType == enum.FLOAT && !tools.IsFloat(value) {
		return errors.New(fmt.Sprintf("BAD_VALUE, <%s> is not <%s>", value, string(enum.FLOAT)))
	}
	if variable.variableType == enum.BOOLEAN && !tools.IsBoolean(value) {
		return errors.New(fmt.Sprintf("BAD_VALUE, <%s> is not <%s>", value, string(enum.BOOLEAN)))
	}
	variable.value = value
	return nil
}

func ValidateTypeOfValue(value string, variableType enum.VariableType) error {
	if (variableType == enum.INTEGER && !tools.IsInteger(value)) ||
		(variableType == enum.FLOAT && !tools.IsFloat(value)) ||
		(variableType == enum.BOOLEAN && !tools.IsBoolean(value)) {
		return errors.New(fmt.Sprintf("<%s> is not an %s value", value, string(variableType)))
	}
	return nil
}

func (variable *Variable) String() string {
	return fmt.Sprintf("[<PanicModeVariable> <%s>, <%s>, <%s>[default: <%s>]]", variable.name, variable.variableType, variable.value, variable.defaultValue)
}
