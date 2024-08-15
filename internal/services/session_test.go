package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSessionService_Add(t *testing.T) {

}

// Добавление хита, но данные с сессией переданы пустые
func TestGuestSessionService_Add_EmptyUserData(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	utils.TruncateTable("session", chClient)

	allModels := models.NewModels(context.Background(), chClient)
	sessionService := NewSession(context.Background(), allModels)

	session, err := sessionService.Add(uuid.Nil, uuid.New(), uuid.New(), false, entityjson.UserData{}, entitydb.AdvReferer{})
	req.NotNil(err)
	req.Equal("statData is empty", err.Error())
	req.Equal(session, entitydb.Session{})
}

func TestGuestSessionService_Add_EmptyGuestUuid(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	utils.TruncateTable("session", chClient)

	allModels := models.NewModels(context.Background(), chClient)
	sessionService := NewSession(context.Background(), allModels)
	hitUuid := uuid.New()
	session, err := sessionService.Add(uuid.Nil, uuid.UUID{}, hitUuid, false, entityjson.UserData{}, entitydb.AdvReferer{})
	req.NotNil(err)
	req.Equal("guestUuid is empty", err.Error())
	req.Equal(session, entitydb.Session{})
}

func TestGuestSessionService_Add_EmptyHitUuid(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	utils.TruncateTable("session", chClient)

	allModels := models.NewModels(context.Background(), chClient)
	sessionService := NewSession(context.Background(), allModels)
	session, err := sessionService.Add(uuid.Nil, uuid.New(), uuid.UUID{}, false, entityjson.UserData{}, entitydb.AdvReferer{})
	req.NotNil(err)
	req.Equal("hitUuid is empty", err.Error())
	req.Equal(session, entitydb.Session{})
}

func TestGuestSessionService_Add(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	utils.TruncateTable("session", chClient)

	allModels := models.NewModels(context.Background(), chClient)
	sessionService := NewSession(context.Background(), allModels)
	guestUuid := uuid.New()
	hitUuid := uuid.New()
	stopListUuid := uuid.New()
	userData := entityjson.UserData{
		PHPSessionId:      "php-ses",
		GuestUuid:         guestUuid,
		Url:               "http://localhost",
		Referer:           "https://www.google.com",
		Ip:                "0.0.0.0",
		UserAgent:         "user-agent-chrome",
		UserId:            10,
		UserLogin:         "admin",
		HttpXForwardedFor: "",
		IsError404:        false,
		SiteId:            "s1",
		Event1:            "",
		Event2:            "",
		IsUserAuth:        true,
		Method:            "GET",
		Cookies:           "cookie-value",
		IsFavorite:        true,
	}
	advReferer := entitydb.AdvReferer{}
	session, err := sessionService.Add(stopListUuid, guestUuid, hitUuid, false, userData, advReferer)
	req.Nil(err)

	var allDbSessions []entitydb.Session
	rows, _ := chClient.Query(context.Background(), "SELECT * from session")

	for rows.Next() {
		var dbSession entitydb.Session
		err = rows.ScanStruct(&dbSession)
		req.Nil(err)
		allDbSessions = append(allDbSessions, dbSession)
	}
	req.Equal(1, len(allDbSessions))

	req.Equal(1, allDbSessions[0].Uuid)
	req.Equal(1, allDbSessions[0].GuestUuid)
	req.Equal(1, allDbSessions[0].IsNewGuest)
	req.Equal(1, allDbSessions[0].IsUserAuth)
	req.Equal(1, allDbSessions[0].Favorites)
	req.Equal(1, allDbSessions[0].UrlFrom)
	req.Equal(1, allDbSessions[0].UrlTo404)
	req.Equal(1, allDbSessions[0].UrlLast)
	req.Equal(1, allDbSessions[0].UrlLast404)

	req.Equal(1, allDbSessions[0].UserAgent)
	req.Equal(1, allDbSessions[0].DateStat)
	req.Equal(1, allDbSessions[0].DateFirst)
	req.Equal(1, allDbSessions[0].DateLast)
	req.Equal(1, allDbSessions[0].IpFirst)
	req.Equal(1, allDbSessions[0].IpLast)
	req.Equal(1, allDbSessions[0].FirstHitUuid)
	req.Equal(1, allDbSessions[0].LastHitUuid)
	req.Equal(1, allDbSessions[0].PhpSessionId)
	req.Equal(1, allDbSessions[0].AdvUuid)
	req.Equal(1, allDbSessions[0].AdvBack)
	req.Equal(1, allDbSessions[0].Referer1)
	req.Equal(1, allDbSessions[0].Referer2)
	req.Equal(1, allDbSessions[0].Referer3)
	req.Equal(1, allDbSessions[0].StopListUuid)
	req.Equal(1, allDbSessions[0].CountryId)

	req.Equal(1, allDbSessions[0].FirstSiteId)
	req.Equal(1, allDbSessions[0].LastSiteId)
	req.Equal(1, allDbSessions[0].CityId)
	req.Equal(1, allDbSessions[0].Sign)
	req.Equal(1, allDbSessions[0].Version)
}
