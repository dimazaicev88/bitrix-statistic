package utils

import (
	"bitrix-statistic/internal/entityjson"
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"strings"
)

func GetGuestMd5(statData entityjson.UserData) string {
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

func Crc32(str string) int32 {
	const IEEE = 0xedb88320
	var result int32
	crc32q := crc32.MakeTable(IEEE)
	result = int32(crc32.Checksum([]byte(str), crc32q))
	if result > 2147483647 {
		result = -(2147483647 - result + 1)
	}
	return result
}
