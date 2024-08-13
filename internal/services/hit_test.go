package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestHitService_Add(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)

	if err := chClient.Exec(context.Background(), "TRUNCATE hit"); err != nil {
		logrus.Fatal(err)
	}
	existsGuest := true
	guestUuid := uuid.New()
	sessionDb := entitydb.Session{
		Uuid:         uuid.New(),
		GuestUuid:    guestUuid,
		IsNewGuest:   true,
		UserId:       10,
		IsUserAuth:   true,
		Events:       3,
		Hits:         5,
		Favorites:    true,
		UrlFrom:      "https://google.com/",
		UrlTo:        "http://localhost/",
		UrlTo404:     false,
		UrlLast:      "",
		UrlLast404:   false,
		UserAgent:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		DateStat:     time.Now(),
		DateFirst:    time.Now(),
		DateLast:     time.Now().Add(5 + time.Hour),
		IpFirst:      "10.136.254.102",
		IpLast:       "10.136.252.152",
		FirstHitUuid: uuid.New(),
		LastHitUuid:  uuid.New(),
		PhpSessionId: "b59c67bf196a4758191e42f76670ceba",
		AdvUuid:      uuid.New(),
		AdvBack:      false,
		Referer1:     "ref1",
		Referer2:     "ref2",
		Referer3:     "ref3",
		StopListUuid: uuid.New().String(),
		CountryId:    "",
		FirstSiteId:  "",
		LastSiteId:   "",
		CityId:       "",
		Sign:         1,
		Version:      1,
	}

	advReferer := entitydb.AdvReferer{
		AdvUuid:     uuid.New(),
		Referer1:    "ref1",
		Referer2:    "ref2",
		Referer3:    "ref3",
		LastAdvBack: false,
	}

	statData := entityjson.StatData{
		PHPSessionId:      "b59c67bf196a4758191e42f76670ceba",
		GuestUuid:         guestUuid,
		Url:               "ttp://localhost/",
		Referer:           "https://google.com/",
		Ip:                "10.136.252.152",
		UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		UserId:            10,
		UserLogin:         "admin",
		HttpXForwardedFor: "",
		IsError404:        false,
		SiteId:            "s1",
		Event1:            "",
		Event2:            "",
		IsUserAuth:        true,
		Method:            "GET",
		Cookies:           "_ga=GA1.2.131634531.1723559453; _gid=GA1.2.1901415454.1723559453; _ga_PWTK27XVWP=GS1.1.1723559453.1.1.1723559817.0.0.0",
		IsFavorite:        false,
	}
	hit, err := hitService.Add(existsGuest, sessionDb, advReferer, statData)
	req.Nil(err)

	req.Equal("", hit.Uuid)
	req.Equal("", hit.SessionUuid)
	req.Equal("", hit.AdvUuid)
	req.Equal("", hit.DateHit)
	req.Equal("", hit.GuestUuid)
	req.Equal("", hit.IsNewGuest)
	req.Equal("", hit.UserId)
	req.Equal("", hit.IsUserAuth)
	req.Equal("", hit.Url)
	req.Equal("", hit.Url404)
	req.Equal("", hit.UrlFrom)
	req.Equal("", hit.Method)
	req.Equal("", hit.Cookies)
	req.Equal("", hit.UserAgent)
	req.Equal("", hit.StopListUuid)
	req.Equal("", hit.CountryId)
	req.Equal("", hit.CityUuid)
	req.Equal("", hit.SiteId)
}
