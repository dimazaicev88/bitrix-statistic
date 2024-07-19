package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearcherService_IsSearcher(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	utils.TruncateAllTables(chClient)
	defer chClient.Close()
	req := require.New(t)
	searchService := NewSearcherService(context.Background(), chClient)
	searcher, err := searchService.IsSearcher("FeedDemon")
	req.NoError(err)
	req.True(searcher)

	searcher, err = searchService.IsSearcher("WFQXZFeedDemon")
	req.NoError(err)
	req.False(searcher)

	searcher, err = searchService.IsSearcher("")
	req.NoError(err)
	req.False(searcher)
}

func TestSearcherService_AddSearcherStatData(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	//utils.TruncateAllTables(chClient)
	defer chClient.Close()
	req := require.New(t)
	searchService := NewSearcherService(context.Background(), chClient)
	err := searchService.AddStatData(entityjson.StatData{
		PHPSessionId:      "",
		GuestHash:         "",
		Url:               "",
		Referer:           "",
		Ip:                "",
		UserAgent:         "",
		UserId:            0,
		UserLogin:         "",
		HttpXForwardedFor: "",
		IsError404:        false,
		SiteId:            "",
		Event1:            "",
		Event2:            "",
		IsUserAuth:        false,
	},
	)
	req.NoError(err)
	req.True(searcher)

	searcher, err = searchService.IsSearcher("WFQXZFeedDemon")
	req.NoError(err)
	req.False(searcher)

	searcher, err = searchService.IsSearcher("")
	req.NoError(err)
	req.False(searcher)
}
