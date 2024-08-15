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

// Добавление хита, но данные с сессией переданы пустые
func TestGuestService_Add_EmptyUserData(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	guestService := NewGuest(context.Background(), allModels, hitService, NewAdv(context.Background(), allModels, hitService))
	utils.TruncateTable("guest", chClient)

	guest, err := guestService.Add(entityjson.UserData{}, entitydb.AdvReferer{})
	req.NotNil(err)
	req.Equal("user data is empty", err.Error())
	req.Equal(guest, entitydb.Guest{})
}

// Добавление хита, но данные с сессией переданы пустые
func TestGuestService_Add(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	guestService := NewGuest(context.Background(), allModels, hitService, NewAdv(context.Background(), allModels, hitService))
	utils.TruncateTable("guest", chClient)

	guestUuid := uuid.New()
	advUuid := uuid.New()
	userData := entityjson.UserData{
		PHPSessionId:      "php-session-id-v1",
		GuestUuid:         guestUuid,
		Url:               "http://localhost",
		Referer:           "https://www.google.com",
		Ip:                "127.0.0.1",
		UserAgent:         "user-agent-firefox",
		UserId:            10,
		UserLogin:         "admin",
		HttpXForwardedFor: "",
		IsError404:        false,
		SiteId:            "s1",
		Event1:            "evt1",
		Event2:            "ev2",
		IsUserAuth:        true,
		Method:            "GET",
		Cookies:           "cookies-value",
		IsFavorite:        true,
	}
	advReferer := entitydb.AdvReferer{
		AdvUuid:     advUuid,
		Referer1:    "ref1",
		Referer2:    "ref2",
		Referer3:    "ref3",
		LastAdvBack: true,
	}
	guest, err := guestService.Add(userData, advReferer)
	req.Nil(err)

	var allDbGuests []entitydb.Guest
	rows, _ := chClient.Query(context.Background(), "SELECT * from guest")

	for rows.Next() {
		var dbGuest entitydb.Guest
		err = rows.ScanStruct(&dbGuest)
		req.Nil(err)
		allDbGuests = append(allDbGuests, dbGuest)
	}
	req.Equal(1, len(allDbGuests))

	req.Equal(guest.Uuid.String(), allDbGuests[0].Uuid.String())
	req.Equal(guest.Favorites, allDbGuests[0].Favorites)
	req.Equal(uint32(0), allDbGuests[0].Events)
	req.Equal(uint32(1), allDbGuests[0].Sessions)
	req.Equal(uint32(1), allDbGuests[0].Hits)
	req.Equal(false, allDbGuests[0].Repair)
	req.Equal(userData.Referer, allDbGuests[0].FirstUrlFrom)
	req.Equal(userData.Url, allDbGuests[0].FirstUrlTo)
	req.Equal(userData.IsError404, allDbGuests[0].FirstUrlTo404)
	req.Equal(userData.SiteId, allDbGuests[0].FirstSiteId)
	req.Equal(advUuid.String(), allDbGuests[0].FirstAdvUuid.String())
	req.Equal(advReferer.Referer1, allDbGuests[0].FirstReferer1)
	req.Equal(advReferer.Referer2, allDbGuests[0].FirstReferer2)
	req.Equal(advReferer.Referer3, allDbGuests[0].FirstReferer3)
	req.Equal(userData.UserId, allDbGuests[0].LastUserId)
	req.Equal(userData.IsUserAuth, allDbGuests[0].LastUserAuth)
	req.Equal(userData.Url, allDbGuests[0].LastUrlLast)
	req.Equal(userData.IsError404, allDbGuests[0].LastUrlLast404)
	req.Equal(userData.UserAgent, allDbGuests[0].LastUserAgent)
	req.Equal(userData.Ip, allDbGuests[0].LastIp)
	req.Equal(userData.Cookies, allDbGuests[0].LastCookie)
	req.Equal(advUuid, allDbGuests[0].LastAdvUUid)
	req.Equal(advReferer.LastAdvBack, allDbGuests[0].LastAdvBack)
	req.Equal(advReferer.Referer1, allDbGuests[0].LastReferer1)
	req.Equal(advReferer.Referer2, allDbGuests[0].LastReferer2)
	req.Equal(advReferer.Referer3, allDbGuests[0].LastReferer3)
	req.Equal(userData.SiteId, allDbGuests[0].LastSiteId)
	req.Equal(int8(1), allDbGuests[0].Sign)
	req.Equal(uint32(1), allDbGuests[0].Version)
	req.Equal(userData.PHPSessionId, allDbGuests[0].PhpSessionId)
}

func TestGuestService_FindByUuid(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	guestService := NewGuest(context.Background(), allModels, hitService, NewAdv(context.Background(), allModels, hitService))
	utils.TruncateTable("guest", chClient)

	guestUuid := uuid.New()
	userData := entityjson.UserData{
		PHPSessionId:      "php-session-id-v1",
		GuestUuid:         guestUuid,
		Url:               "http://localhost",
		Referer:           "https://www.google.com",
		Ip:                "127.0.0.1",
		UserAgent:         "user-agent-firefox",
		UserId:            10,
		UserLogin:         "admin",
		HttpXForwardedFor: "",
		IsError404:        false,
		SiteId:            "s1",
		Event1:            "evt1",
		Event2:            "ev2",
		IsUserAuth:        true,
		Method:            "GET",
		Cookies:           "cookies-value",
		IsFavorite:        true,
	}
	guest, err := guestService.Add(userData, entitydb.AdvReferer{})
	guestFind, err := guestService.FindByUuid(userData.GuestUuid)

	req.Nil(err)
	req.Equal(guest.Uuid.String(), guestFind.Uuid.String())
}

func TestGuestService_Update(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	guestService := NewGuest(context.Background(), allModels, hitService, NewAdv(context.Background(), allModels, hitService))
	utils.TruncateTable("guest", chClient)

	guestUuid := uuid.New()
	userData := entityjson.UserData{
		PHPSessionId:      "php-session-id-v1",
		GuestUuid:         guestUuid,
		Url:               "http://localhost",
		Referer:           "https://www.google.com",
		Ip:                "127.0.0.1",
		UserAgent:         "user-agent-firefox",
		UserId:            10,
		UserLogin:         "admin",
		HttpXForwardedFor: "",
		IsError404:        false,
		SiteId:            "s1",
		Event1:            "evt1",
		Event2:            "ev2",
		IsUserAuth:        true,
		Method:            "GET",
		Cookies:           "cookies-value",
		IsFavorite:        true,
	}
	guest, err := guestService.Add(userData, entitydb.AdvReferer{})
	newGuest := guest
	newGuest.Events = 1
	newGuest.Favorites = true
	newGuest.Sessions = 10
	newGuest.Hits = 2
	newGuest.Repair = true
	newGuest.PhpSessionId = "php-session-id-v2"
	newGuest.FirstSessionUuid = uuid.New()
	newGuest.FirstUrlFrom = "https://www.ozone.com"
	newGuest.FirstUrlTo = "https://www.bbb.com"
	newGuest.FirstUrl404 = true
	newGuest.FirstUrlTo404 = true
	newGuest.FirstSiteId = "s2"
	newGuest.FirstAdvUuid = uuid.New()
	newGuest.FirstReferer1 = "rrrr1"
	newGuest.FirstReferer2 = "rrrr2"
	newGuest.FirstReferer3 = "rrrr3"
	newGuest.LastSessionUuid = uuid.New()
	newGuest.LastUserId = 1002
	newGuest.LastUserAuth = true
	newGuest.LastUrlLast = "https://www.bbb.com"
	newGuest.LastUrlLast404 = true
	newGuest.LastUserAgent = "user-agent-chrome"
	newGuest.LastIp = "0.0.0.0"
	newGuest.LastCookie = "cookies-value-v2"
	newGuest.LastLanguage = "ru"
	newGuest.LastAdvUUid = uuid.New()
	newGuest.LastAdvBack = false
	newGuest.LastReferer1 = "r1"
	newGuest.LastReferer2 = "r2"
	newGuest.LastReferer3 = "r3"
	newGuest.LastSiteId = "s1"
	newGuest.LastCountryId = "c1"
	newGuest.LastCityId = "ct1"
	newGuest.LastCityInfo = "ci1"

	err = guestService.UpdateGuest(guest, newGuest)
	req.Nil(err)

	err = chClient.Exec(context.Background(), "OPTIMIZE TABLE guest DEDUPLICATE;")
	req.Nil(err)

	guestAfterUpdate, _ := guestService.FindByUuid(guest.Uuid)

	req.Equal(guestAfterUpdate.Uuid.String(), newGuest.Uuid.String())
	req.Equal(guestAfterUpdate.Favorites, newGuest.Favorites)
	req.Equal(guestAfterUpdate.Events, newGuest.Events)
	req.Equal(guestAfterUpdate.Sessions, newGuest.Sessions)
	req.Equal(guestAfterUpdate.Hits, newGuest.Hits)
	req.Equal(guestAfterUpdate.Repair, newGuest.Repair)
	req.Equal(guestAfterUpdate.FirstUrlFrom, newGuest.FirstUrlFrom)
	req.Equal(guestAfterUpdate.FirstUrlTo, newGuest.FirstUrlTo)
	req.Equal(guestAfterUpdate.FirstUrlTo404, newGuest.FirstUrlTo404)
	req.Equal(guestAfterUpdate.FirstSiteId, newGuest.FirstSiteId)
	req.Equal(guestAfterUpdate.FirstAdvUuid.String(), newGuest.FirstAdvUuid.String())
	req.Equal(guestAfterUpdate.FirstReferer1, newGuest.FirstReferer1)
	req.Equal(guestAfterUpdate.FirstReferer2, newGuest.FirstReferer2)
	req.Equal(guestAfterUpdate.FirstReferer3, newGuest.FirstReferer3)
	req.Equal(guestAfterUpdate.LastUserId, newGuest.LastUserId)
	req.Equal(guestAfterUpdate.LastUserAuth, newGuest.LastUserAuth)
	req.Equal(guestAfterUpdate.LastUrlLast, newGuest.LastUrlLast)
	req.Equal(guestAfterUpdate.LastUrlLast404, newGuest.LastUrlLast404)
	req.Equal(guestAfterUpdate.LastUserAgent, newGuest.LastUserAgent)
	req.Equal(guestAfterUpdate.LastIp, newGuest.LastIp)
	req.Equal(guestAfterUpdate.LastCookie, newGuest.LastCookie)
	req.Equal(guestAfterUpdate.LastAdvUUid, newGuest.LastAdvUUid)
	req.Equal(guestAfterUpdate.LastAdvBack, newGuest.LastAdvBack)
	req.Equal(guestAfterUpdate.LastReferer1, newGuest.LastReferer1)
	req.Equal(guestAfterUpdate.LastReferer2, newGuest.LastReferer2)
	req.Equal(guestAfterUpdate.LastReferer3, newGuest.LastReferer3)
	req.Equal(guestAfterUpdate.LastSiteId, newGuest.LastSiteId)
	req.Equal(guestAfterUpdate.Sign, newGuest.Sign)
	req.Equal(guestAfterUpdate.Version, newGuest.Version)
	req.Equal(guestAfterUpdate.PhpSessionId, newGuest.PhpSessionId)
}
