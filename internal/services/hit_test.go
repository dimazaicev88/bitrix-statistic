package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

// Добавление хита
func TestHitService_Add(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	utils.TruncateTable("hit", chClient)
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
		StopListUuid: uuid.New(),
		CountryId:    "",
		FirstSiteId:  "",
		LastSiteId:   "",
		CityId:       "",
		Sign:         1,
		Version:      1,
	}

	advReferer := entitydb.AdvCompany{
		AdvUuid:     uuid.New(),
		Referer1:    "ref1",
		Referer2:    "ref2",
		Referer3:    "ref3",
		LastAdvBack: false,
	}

	statData := dto.UserData{
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

// Добавление хита, но данные с сессией переданы пустые
func TestHitService_Add_EmptySessRefererAndStatData(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	utils.TruncateTable("hit", chClient)

	hit, err := hitService.Add(true, entitydb.Session{}, entitydb.AdvCompany{}, dto.UserData{})
	req.NotNil(err)
	req.Equal("session is empty", err.Error())
	req.Equal(hit, entitydb.Hit{})
}

// Добавление хита, но данные с данными о госте переданы пустые
func TestHitService_Add_EmptyStatData(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	utils.TruncateTable("hit", chClient)
	sessionDb := entitydb.Session{Uuid: uuid.New(),
		GuestUuid: uuid.New(),
	}

	hit, err := hitService.Add(true, sessionDb, entitydb.AdvCompany{}, dto.UserData{})
	req.NotNil(err)
	req.Equal("stat data is empty", err.Error())
	req.Equal(hit, entitydb.Hit{})
}

func TestHitService_FindLastHitWithoutSession(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	utils.TruncateTable("hit", chClient)

	guestUuid := uuid.New()

	err := chClient.Exec(context.Background(),
		`INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url,
														  url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
			   VALUES ('0191509a-eca7-760b-b46f-664cbfb5fb03', generateUUIDv7(), generateUUIDv7(), now(), 'ses1', @guestUuid, true, 1, true,
														'localhost', false, '', '10.136.254.100', 'get', 'cookie', 'user_ag', generateUUIDv7(), 'ru', generateUUIDv7(),	's1')`, clickhouse.Named("guestUuid", guestUuid))

	req.Nil(err)
	err = chClient.Exec(context.Background(), `INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url,
														  url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
			   VALUES ('0191508b-3d75-7048-8733-bb3a24cd6ed1', generateUUIDv7(), generateUUIDv7(), now(), 'ses1', @guestUuid, true, 1, true,
														'localhost', false, '', '10.136.254.100', 'get', 'cookie', 'user_ag', generateUUIDv7(), 'ru', generateUUIDv7(), 's1')`, clickhouse.Named("guestUuid", guestUuid))

	req.Nil(err)
	err = chClient.Exec(context.Background(), `INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url,
														  url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
			   VALUES ('0191508d-ec3b-7277-a776-69b528e8c327', generateUUIDv7(), generateUUIDv7(), now(), 'ses1', @guestUuid, true, 1, true,
														'localhost', false, '', '10.136.254.100', 'get', 'cookie', 'user_ag', generateUUIDv7(), 'ru', generateUUIDv7(), 's1')`, clickhouse.Named("guestUuid", guestUuid))

	req.Nil(err)
	err = chClient.Exec(context.Background(), `INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url,
														  url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
			   VALUES ('0191508f-ef1c-7b30-9009-4241f1e272a8', generateUUIDv7(), generateUUIDv7(), now(), 'ses1', @guestUuid, true, 1, true,
														'localhost', false, '', '10.136.254.100', 'get', 'cookie', 'user_ag', generateUUIDv7(), 'ru', generateUUIDv7(),	's1')`, clickhouse.Named("guestUuid", guestUuid))

	req.Nil(err)
	err = chClient.Exec(context.Background(), `INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url,
														  url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
			   VALUES ('0191508e-d648-7369-946e-3248ccddeab9', generateUUIDv7(), generateUUIDv7(), now(), 'ses2', @guestUuid, true, 1, true,
														'localhost', false, '', '10.136.254.100', 'get', 'cookie', 'user_ag', generateUUIDv7(), 'ru', generateUUIDv7(),'s1')`, clickhouse.Named("guestUuid", guestUuid))
	req.Nil(err)

	hit, err := hitService.FindLastHitWithoutSession(guestUuid, "ses2")
	req.Nil(err)
	req.Equal("0191508f-ef1c-7b30-9009-4241f1e272a8", hit.Uuid.String())

}

func TestHitService_FindByUuid(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	utils.TruncateTable("hit", chClient)

	hitUuid := uuid.New()

	err := chClient.Exec(context.Background(),
		`INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url,
														  url_404, url_from, ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
			   VALUES (@uuid, generateUUIDv7(), generateUUIDv7(), now(), 'ses1', generateUUIDv7(), true, 1, true,
														'localhost', false, '', '10.136.254.100', 'get', 'cookie', 'user_ag', generateUUIDv7(), 'ru', generateUUIDv7(),	's1');`, clickhouse.Named("uuid", hitUuid))
	req.Nil(err)

	hit, err := hitService.FindByUuid(hitUuid)
	req.Nil(err)
	req.Equal(hitUuid, hit.Uuid)
}
