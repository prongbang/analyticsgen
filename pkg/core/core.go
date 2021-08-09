package core

import (
	"fmt"
	"reflect"
	"strings"
)

func FirstUpperCase(value string) string {
	if len(value) == 0 {
		return value
	}
	first := value[:1]
	other := value[1:]
	variable := strings.ToUpper(first) + other
	return variable
}

func FirstLowerCase(value string) string {
	if len(value) == 0 {
		return value
	}
	first := value[:1]
	other := value[1:]
	variable := strings.ToLower(first) + other
	return variable
}

func VariableCamel(value string) string {
	if len(value) == 0 {
		return value
	}
	first := value[:1]
	other := value[1:]
	variable := strings.ToLower(first) + other
	underscore := strings.Split(variable, "_")
	camelCase := ""
	for i, n := range underscore {
		if i == 0 {
			camelCase += FirstLowerCase(n)
		} else {
			camelCase += FirstUpperCase(n)
		}
	}
	return camelCase
}

func ToString(inf interface{}) string {
	return fmt.Sprintf("%v", inf)
}

func GetMapKeys(inf interface{}) []reflect.Value {
	return reflect.ValueOf(inf).MapKeys()
}