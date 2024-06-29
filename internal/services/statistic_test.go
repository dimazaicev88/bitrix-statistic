package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
	"github.com/stretchr/testify/require"
	"testing"
)

var storageConfig = config.ServerEnvConfig{
	StorageHost:     "localhost",
	StorageUser:     "bitrix",
	StoragePassword: "123",
	StorageDbName:   "bitrix",
	StoragePort:     "3306",
}

func TestGuestModel_Add(t *testing.T) {
	mysqlStorage := storage.NewMysqlStorage(storageConfig)
	defer mysqlStorage.Close()
	req := require.New(t)
	err := NewStatistic(mysqlStorage).Add(entity.StatData{
		PhpSessionId:      "",
		CookieToken:       "44c2870053b0a6378f5db40c96406f00",
		SessionToken:      "55c2870053b0a6378f5db40c96406f00",
		Url:               "http://localhost/catalog/dresses/dress-fashionista-on-a-walk/",
		Referer:           "",
		Ip:                "127.0.0.1",
		UserAgent:         "Mozila",
		UserId:            0,
		HttpXForwardedFor: "",
		Error404:          "N",
		SiteId:            "s1",
	})
	req.NoError(err)

}