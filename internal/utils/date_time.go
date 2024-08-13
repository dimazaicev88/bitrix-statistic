package utils

import "time"

func GetCurrentDate() string {
	return time.Now().Local().Format("2006-01-02")
}
