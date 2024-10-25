package utils

import "strings"

func StringConcat(values ...string) string {
	var builder strings.Builder
	for _, item := range values {
		builder.WriteString(item)
	}
	return builder.String()
}

func IsInt(value any) bool {
	_, ok := value.(int)
	return ok
}
