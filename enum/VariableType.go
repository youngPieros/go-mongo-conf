package enum

type VariableType string

const (
	INTEGER VariableType = "int"
	STRING  VariableType = "string"
	FLOAT   VariableType = "float"
	BOOLEAN VariableType = "boolean"
	NONE    VariableType = "NONE"
)

func GetVariableTypeFrom(variableType string) VariableType {
	switch variableType {
	case string(INTEGER):
		return INTEGER
	case string(STRING):
		return STRING
	case string(FLOAT):
		return FLOAT
	case string(BOOLEAN):
		return BOOLEAN
	default:
		return NONE
	}
}
