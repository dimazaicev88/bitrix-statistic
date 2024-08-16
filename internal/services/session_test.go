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
	"strings"
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

// TODO добавить проверку что метод Add возвращает правильно сформированную структуру.
func TestGuestSessionService_Add(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Nil(err)
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
	advReferer := entitydb.AdvReferer{
		AdvUuid:     uuid.New(),
		Referer1:    "r1",
		Referer2:    "r2",
		Referer3:    "r3",
		LastAdvBack: true,
	}
	session, err := sessionService.Add(stopListUuid, guestUuid, hitUuid, false, userData, advReferer)
	req.Nil(err)

	var allDbSessions []entitydb.Session
	rows, err := chClient.Query(context.Background(), "SELECT * from session")
	req.Nil(err)

	for rows.Next() {
		var dbSession entitydb.Session
		err = rows.ScanStruct(&dbSession)
		req.Nil(err)
		allDbSessions = append(allDbSessions, dbSession)
	}
	req.Equal(1, len(allDbSessions))

	req.Equal(session.Uuid, allDbSessions[0].Uuid)
	req.Equal(guestUuid, allDbSessions[0].GuestUuid)
	req.Equal(false, allDbSessions[0].IsNewGuest)
	req.Equal(true, allDbSessions[0].IsUserAuth)
	req.Equal(true, allDbSessions[0].Favorites)
	req.Equal("https://www.google.com", allDbSessions[0].UrlFrom)
	req.Equal(false, allDbSessions[0].UrlTo404)
	req.Equal("http://localhost", allDbSessions[0].UrlLast)
	req.Equal(false, allDbSessions[0].UrlLast404)
	req.Equal("user-agent-chrome", allDbSessions[0].UserAgent)
	req.Equal("0.0.0.0", allDbSessions[0].IpFirst)
	req.Equal("0.0.0.0", allDbSessions[0].IpLast)
	req.Equal(hitUuid, allDbSessions[0].FirstHitUuid)
	req.Equal(hitUuid, allDbSessions[0].LastHitUuid)
	req.Equal("php-ses", allDbSessions[0].PhpSessionId)
	req.Equal(advReferer.AdvUuid, allDbSessions[0].AdvUuid)
	req.Equal(advReferer.LastAdvBack, allDbSessions[0].AdvBack)
	req.Equal(advReferer.Referer1, allDbSessions[0].Referer1)
	req.Equal(advReferer.Referer2, allDbSessions[0].Referer2)
	req.Equal(advReferer.Referer3, allDbSessions[0].Referer3)
	req.Equal(stopListUuid, allDbSessions[0].StopListUuid)
	req.Equal("", strings.Trim("", allDbSessions[0].CountryId))
	req.Equal("s1", allDbSessions[0].FirstSiteId)
	req.Equal("s1", allDbSessions[0].LastSiteId)
	req.Equal("", allDbSessions[0].CityId)
	req.Equal(int8(1), allDbSessions[0].Sign)
	req.Equal(uint32(1), allDbSessions[0].Version)
}

func TestGuestSessionService_Update(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Nil(err)
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
	advReferer := entitydb.AdvReferer{
		AdvUuid:     uuid.New(),
		Referer1:    "r1",
		Referer2:    "r2",
		Referer3:    "r3",
		LastAdvBack: true,
	}

	oldSession, err := sessionService.Add(stopListUuid, guestUuid, hitUuid, false, userData, advReferer)
	req.Nil(err)

	newSession := oldSession
	newSession.IsNewGuest = true
	newSession.UserId = 100
	newSession.IsUserAuth = false
	newSession.Events = 100
	newSession.Hits = 100
	newSession.Favorites = true
	newSession.UrlFrom = ""
	newSession.UrlTo = ""
	newSession.UrlTo404 = true
	newSession.UrlLast = ""
	newSession.UrlLast404 = true
	newSession.UserAgent = ""
	newSession.IpFirst = "127.0.0.1"
	newSession.IpLast = "127.0.0.95"
	newSession.FirstHitUuid = uuid.Nil
	newSession.LastHitUuid = uuid.Nil
	newSession.PhpSessionId = "php-ses-new"
	newSession.AdvUuid = uuid.Nil
	newSession.AdvBack = false
	newSession.Referer1 = "nr1"
	newSession.Referer2 = "nr2"
	newSession.Referer3 = "nr3"
	newSession.StopListUuid = uuid.Nil
	newSession.CountryId = ""
	newSession.FirstSiteId = ""
	newSession.LastSiteId = ""
	newSession.CityId = ""

	err = sessionService.Update(oldSession, newSession)
	req.Nil(err)

	var allDbSessions []entitydb.Session
	rows, err := chClient.Query(context.Background(), "SELECT * from session")
	req.Nil(err)

	for rows.Next() {
		var dbSession entitydb.Session
		err = rows.ScanStruct(&dbSession)
		req.Nil(err)
		allDbSessions = append(allDbSessions, dbSession)
	}

	req.Equal(newSession.Uuid, allDbSessions[0].Uuid)
	req.Equal(newSession, allDbSessions[0].GuestUuid)
	req.Equal(newSession, allDbSessions[0].IsNewGuest)
	req.Equal(newSession, allDbSessions[0].IsUserAuth)
	req.Equal(newSession, allDbSessions[0].Favorites)
	req.Equal(newSession, allDbSessions[0].UrlFrom)
	req.Equal(newSession, allDbSessions[0].UrlTo404)
	req.Equal(newSession, allDbSessions[0].UrlLast)
	req.Equal(newSession, allDbSessions[0].UrlLast404)
	req.Equal(newSession, allDbSessions[0].UserAgent)
	req.Equal(newSession, allDbSessions[0].IpFirst)
	req.Equal(newSession, allDbSessions[0].IpLast)
	req.Equal(newSession, allDbSessions[0].FirstHitUuid)
	req.Equal(newSession, allDbSessions[0].LastHitUuid)
	req.Equal(newSession, allDbSessions[0].PhpSessionId)
	req.Equal(newSession.AdvUuid, allDbSessions[0].AdvUuid)
	req.Equal(newSession.AdvBack, allDbSessions[0].AdvBack)
	req.Equal(newSession.Referer1, allDbSessions[0].Referer1)
	req.Equal(newSession.Referer2, allDbSessions[0].Referer2)
	req.Equal(advReferer.Referer3, allDbSessions[0].Referer3)
	req.Equal(newSession, allDbSessions[0].StopListUuid)
	req.Equal(newSession, strings.Trim("", allDbSessions[0].CountryId))
	req.Equal(newSession, allDbSessions[0].FirstSiteId)
	req.Equal(newSession, allDbSessions[0].LastSiteId)
	req.Equal(newSession, allDbSessions[0].CityId)
	req.Equal(int8(1), allDbSessions[0].Sign)
	req.Equal(uint32(2), allDbSessions[0].Version)
}
