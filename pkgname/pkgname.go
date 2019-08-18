package pkgname

import "strings"

func Name(typeName string) string {
	if typeName == "interface{}" {
		return "finterface"
	}

	if strings.HasPrefix(typeName, "*") {
		return "fp" + typeName[1:]
	}

	return "f" + typeName
}
