package repository

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestHit(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("E:\\projects\\bitrix-statistic\\.env")
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}

	req := require.New(t)
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	hitRepo := NewHit(chClient)

	t.Run("Add", func(t *testing.T) {
		utils.TruncateTable(ctx, "hits", chClient)
		hit := models.Hit{
			Uuid:         uuid.New(),
			PhpSessionId: "test_PhpSessionId",
			Event1:       "test_Event1",
			Event2:       "test_Event2",
			DateHit:      time.Now(),
			GuestHash:    utils.GetMD5Hash("test_GuestHash"),
			IsNewGuest:   true,
			UserId:       110,
			Url:          "test_Url",
			Referer:      "test_Referer",
			Url404:       true,
			UrlFrom:      "test_UrlFrom",
			Ip:           "127.0.0.1",
			Method:       "test_Method",
			Cookies:      "test_Cookies",
			UserAgent:    "test_UserAgent",
			SiteId:       "OM",
		}

		err = hitRepo.AddHit(ctx, hit, true)
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

		req.Equal(hit.Uuid, singleHit.Uuid)
		req.Equal(hit.PhpSessionId, singleHit.PhpSessionId)
		req.Equal(hit.Event1, singleHit.Event1)
		req.Equal(hit.Event2, singleHit.Event2)
		req.Equal(hit.GuestHash, singleHit.GuestHash)
		req.Equal(hit.IsNewGuest, singleHit.IsNewGuest)
		req.Equal(hit.UserId, singleHit.UserId)
		req.Equal(hit.Url, singleHit.Url)
		req.Equal(hit.Referer, singleHit.Referer)
		req.Equal(hit.Url404, singleHit.Url404)
		req.Equal(hit.UrlFrom, singleHit.UrlFrom)
		req.Equal(hit.Ip, singleHit.Ip)
		req.Equal(hit.Method, singleHit.Method)
		req.Equal(hit.Cookies, singleHit.Cookies)
		req.Equal(hit.UserAgent, singleHit.UserAgent)
		req.Equal(hit.SiteId, singleHit.SiteId)
		req.Equal(singleHit.Uuid, singleHit.Uuid)

	})
}
