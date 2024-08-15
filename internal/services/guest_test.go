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
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

// Добавление хита, но данные с сессией переданы пустые
func TestGuestService_Add_EmptyUserData(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	req := require.New(t)
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
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	req := require.New(t)
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
		if err != nil {
			req.Fail(err.Error())
		}
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
