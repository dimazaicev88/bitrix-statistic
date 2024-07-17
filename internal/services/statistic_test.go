package services

import (
	"bitrix-statistic/internal/config"
	"testing"
)

var storageConfig = config.ServerEnvConfig{
	ClickHouseHost:     "localhost",
	ClickHouseUser:     "bitrix",
	ClickHousePassword: "123",
	ClickHouseDbName:   "bitrix",
}

func TestGuestModel_Add(t *testing.T) {
	//chClient, _ := storage.NewClickHouseClient(storageConfig)
	//defer chClient.Close()
	//req := require.New(t)
	//err := NewStatistic(chClient).Add(entity.StatData{
	//	PHPSessionId:      "",
	//	GuestHash:         "44c2870053b0a6378f5db40c96406f00",
	//	Url:               "http://localhost/catalog/dresses/dress-fashionista-on-a-walk/",
	//	Referer:           "",
	//	Ip:                "127.0.0.1",
	//	UserAgent:         "Mozila",
	//	UserId:            0,
	//	HttpXForwardedFor: "",
	//	IsError404:        0,
	//	SiteId:            "s1",
	//})
	//req.NoError(err)
}
