package tools

import "strconv"

func IsInteger(value string) bool {
	_, err := strconv.ParseInt(value, 0, 64)
	return err == nil
}

func IsFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func IsBoolean(value string) bool {
	if value == "true" || value == "True" || value == "TRUE" {
		return true
	}
	if value == "false" || value == "False" || value == "FALSE" {
		return true
	}
	return false
}
