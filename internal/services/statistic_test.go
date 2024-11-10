package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/repository"
	"bitrix-statistic/internal/storage"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

const pathToEnvFile = "/home/zajtsev/projects/bitrix-statistic/.env"

func TestStatistic(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("E:\\projects\\bitrix-statistic\\.env")
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}

	req := require.New(t)
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	statisticService := NewStatistic(NewGuest(repository.NewGuest(chClient)), NewHit(repository.NewHit(chClient)))

	t.Run("Add", func(t *testing.T) {
		userData := dto.UserData{
			PHPSessionId:      "",
			GuestHash:         "",
			Url:               "",
			Referer:           "",
			Ip:                "",
			UserAgent:         "",
			UserId:            0,
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "",
			Lang:              "",
			Event1:            "",
			Event2:            "",
			Event3:            "",
			Method:            "",
			Cookies:           "",
		}
		err = statisticService.Add(ctx, userData, true)
		req.NoError(err)

	})
}
