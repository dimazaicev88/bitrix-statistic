package utils

import (
	"bitrix-statistic/internal/entityjson"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func GetGuestMd5(statData entityjson.StatData) string {
	var strBuilder strings.Builder
	strBuilder.WriteString(statData.UserAgent)
	strBuilder.WriteString(statData.Ip)
	strBuilder.WriteString(statData.HttpXForwardedFor)
	sum := GetMD5Hash(strBuilder.String())
	return sum
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
