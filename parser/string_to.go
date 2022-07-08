package parser

import (
	"strconv"
)

// StrToInt convert the string to int
// default value is 0 once cannot convert the value
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}

// StrToInt64 convert string to int64
func StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return i
}

// StrToBool converts string to boolean
func StrToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}

	return b
}
