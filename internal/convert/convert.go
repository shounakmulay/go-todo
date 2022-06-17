package convert

import (
	errorutl "go-todo/internal/error"
	"strconv"
)

func StringToInt(s string) int {
	integer, err := strconv.Atoi(s)
	if err != nil {
		errorutl.Log(err)
		return 0
	}
	return integer
}

func StringToBool(s string) bool {
	boolean, err := strconv.ParseBool(s)
	if err != nil {
		errorutl.Log(err)
		return false
	}
	return boolean
}
