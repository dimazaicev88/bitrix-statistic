package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/repository"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatistic(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("/home/zajtsev/projects/bitrix-statistic/.env")
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}

	req := require.New(t)
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	statisticService := NewStatistic(NewGuest(repository.NewGuest(chClient)), NewHit(repository.NewHit(chClient)))

	t.Run("Add", func(t *testing.T) {
		utils.TruncateTable(ctx, "hits", chClient)

		userData := dto.UserData{
			PHPSessionId:      uuid.New().String(),
			GuestHash:         utils.GetMD5Hash("test_guestHash"),
			Url:               "test_Url",
			Referer:           "test_Referer",
			Ip:                "127.0.0.1",
			UserAgent:         "test_UserAgent",
			UserId:            0,
			HttpXForwardedFor: "127.0.0.10",
			IsError404:        true,
			SiteId:            "om",
			Lang:              "ru",
			Event1:            "ev1",
			Event2:            "ev2",
			Event3:            "ev3",
			Method:            "get",
			Cookies:           "test_cookies",
		}
		err = statisticService.Add(ctx, userData, true)
		req.NoError(err)

		var singleHit models.Hit
		var allDbHits []models.Hit
		rows, _ := chClient.Query(context.Background(), "SELECT * from hits")

		for rows.Next() {
			var dbHit models.Hit
			err = rows.ScanStruct(&dbHit)
			req.Nil(err)
			allDbHits = append(allDbHits, dbHit)
		}
		req.Equal(1, len(allDbHits))
		singleHit = allDbHits[0]

		req.Equal(userData.PHPSessionId, singleHit.PhpSessionId)
		req.Equal(userData.Event1, singleHit.Event1)
		req.Equal(userData.Event2, singleHit.Event2)
		req.Equal(userData.GuestHash, singleHit.GuestHash)
		req.Equal(true, singleHit.IsNewGuest)
		req.Equal(userData.UserId, singleHit.UserId)
		req.Equal(userData.Url, singleHit.Url)
		req.Equal(userData.Referer, singleHit.Referer)
		req.Equal(userData.IsError404, singleHit.Url404)
		req.Equal(userData.Ip, singleHit.Ip)
		req.Equal(userData.Method, singleHit.Method)
		req.Equal(userData.Cookies, singleHit.Cookies)
		req.Equal(userData.UserAgent, singleHit.UserAgent)
		req.Equal(userData.SiteId, singleHit.SiteId)
	})
}
