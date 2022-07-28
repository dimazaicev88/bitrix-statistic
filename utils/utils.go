package utils

import "strconv"

func StrToInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
