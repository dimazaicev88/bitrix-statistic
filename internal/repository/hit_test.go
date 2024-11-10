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
	chClient := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	hitRepo := NewHit(chClient)

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
		Ip:           "test_Ip",
		Method:       "test_Method",
		Cookies:      "test_Cookies",
		UserAgent:    "test_UserAgent",
		SiteId:       "test_SiteId",
	}

	t.Run("Add", func(t *testing.T) {
		err = hitRepo.AddHit(ctx, hit)
		req.NoError(err)
		var dbHit models.Hit
		chClient.NewSelect().Model(&dbHit).Scan(ctx)

		req.Equal(hit.Uuid, dbHit.Uuid)
		req.Equal(hit.PhpSessionId, dbHit.PhpSessionId)
		req.Equal(hit.Event1, dbHit.Event1)
		req.Equal(hit.Event2, dbHit.Event2)
		req.Equal(hit.GuestHash, dbHit.GuestHash)
		req.Equal(hit.IsNewGuest, dbHit.IsNewGuest)
		req.Equal(hit.UserId, dbHit.UserId)
		req.Equal(hit.Url, dbHit.Url)
		req.Equal(hit.Referer, dbHit.Referer)
		req.Equal(hit.Url404, dbHit.Url404)
		req.Equal(hit.UrlFrom, dbHit.UrlFrom)
		req.Equal(hit.Ip, dbHit.Ip)
		req.Equal(hit.Method, dbHit.Method)
		req.Equal(hit.Cookies, dbHit.Cookies)
		req.Equal(hit.UserAgent, dbHit.UserAgent)
		req.Equal(hit.SiteId, dbHit.SiteId)
		req.Equal(dbHit.Uuid, dbHit.Uuid)

	})
}
