package utils

import (
	"bitrix-statistic/internal/entityjson"
	"github.com/codingsince1985/checksum"
	"strconv"
	"strings"
)

func StrToInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}

func GetGuestMd5(statData entityjson.StatData) (string, error) {
	var strBuilder strings.Builder
	strBuilder.WriteString(statData.UserAgent)
	strBuilder.WriteString(statData.Ip)
	strBuilder.WriteString(statData.HttpXForwardedFor)
	sum, err := checksum.MD5sum(strBuilder.String())
	return sum, err
}
