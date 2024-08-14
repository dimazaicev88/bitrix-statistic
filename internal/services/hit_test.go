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
	"strings"
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
	hit, err := hitService.Add(true, sessionDb, advReferer, statData)
	req.Nil(err)

	var allDbHits []entitydb.Hit
	rows, _ := chClient.Query(context.Background(), "SELECT * from hit")

	for rows.Next() {
		var dbHit entitydb.Hit
		rows.ScanStruct(&dbHit)
		allDbHits = append(allDbHits, dbHit)
	}
	req.Equal(1, len(allDbHits))

	req.Equal(hit.Uuid.String(), allDbHits[0].Uuid.String())
	req.Equal(hit.SessionUuid.String(), allDbHits[0].SessionUuid.String())
	req.Equal(hit.AdvUuid.String(), allDbHits[0].AdvUuid.String())
	req.Equal(hit.GuestUuid.String(), allDbHits[0].GuestUuid.String())
	req.Equal(hit.IsNewGuest, allDbHits[0].IsNewGuest)
	req.Equal(hit.PhpSessionId, allDbHits[0].PhpSessionId)
	req.Equal(hit.UserId, allDbHits[0].UserId)
	req.Equal(hit.IsUserAuth, allDbHits[0].IsUserAuth)
	req.Equal(hit.Url, allDbHits[0].Url)
	req.Equal(hit.Url404, allDbHits[0].Url404)
	req.Equal(hit.UrlFrom, allDbHits[0].UrlFrom)
	req.Equal(hit.Method, allDbHits[0].Method)
	req.Equal(hit.Cookies, allDbHits[0].Cookies)
	req.Equal(hit.UserAgent, allDbHits[0].UserAgent)
	req.Equal(hit.StopListUuid, allDbHits[0].StopListUuid)
	req.Equal(strings.Trim(" ", hit.CountryId), strings.Trim(" ", allDbHits[0].CountryId))
	req.Equal(hit.CityUuid.String(), allDbHits[0].CityUuid.String())
	req.Equal(hit.SiteId, allDbHits[0].SiteId)
}

func TestHitService_Add_EmptyData(t *testing.T) {
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
	sessionDb := entitydb.Session{}

	advReferer := entitydb.AdvReferer{}

	statData := entityjson.StatData{}
	hit, err := hitService.Add(true, sessionDb, advReferer, statData)
	req.Nil(err)

	var allDbHits []entitydb.Hit
	rows, _ := chClient.Query(context.Background(), "SELECT * from hit")

	for rows.Next() {
		var dbHit entitydb.Hit
		rows.ScanStruct(&dbHit)
		allDbHits = append(allDbHits, dbHit)
	}
	req.Equal(1, len(allDbHits))

	req.Equal(hit.Uuid.String(), allDbHits[0].Uuid.String())
	req.Equal(hit.SessionUuid.String(), allDbHits[0].SessionUuid.String())
	req.Equal(hit.AdvUuid.String(), allDbHits[0].AdvUuid.String())
	req.Equal(hit.GuestUuid.String(), allDbHits[0].GuestUuid.String())
	req.Equal(hit.IsNewGuest, allDbHits[0].IsNewGuest)
	req.Equal(hit.PhpSessionId, allDbHits[0].PhpSessionId)
	req.Equal(hit.UserId, allDbHits[0].UserId)
	req.Equal(hit.IsUserAuth, allDbHits[0].IsUserAuth)
	req.Equal(hit.Url, allDbHits[0].Url)
	req.Equal(hit.Url404, allDbHits[0].Url404)
	req.Equal(hit.UrlFrom, allDbHits[0].UrlFrom)
	req.Equal(hit.Method, allDbHits[0].Method)
	req.Equal(hit.Cookies, allDbHits[0].Cookies)
	req.Equal(hit.UserAgent, allDbHits[0].UserAgent)
	req.Equal(hit.StopListUuid, allDbHits[0].StopListUuid)
	req.Equal(strings.Trim(" ", hit.CountryId), strings.Trim(" ", allDbHits[0].CountryId))
	req.Equal(hit.CityUuid.String(), allDbHits[0].CityUuid.String())
	req.Equal(hit.SiteId, allDbHits[0].SiteId)
}
