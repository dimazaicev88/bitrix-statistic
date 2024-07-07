package utils

import "strings"

func StringConcat(values ...string) string {
	var builder strings.Builder
	for _, item := range values {
		builder.WriteString(item)
	}
	return builder.String()
}

func IsInt(value interface{}) bool {
	_, ok := value.(int)
	return ok
}
