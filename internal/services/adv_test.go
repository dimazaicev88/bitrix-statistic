package services

import (
	"bitrix-statistic/internal/config"
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

//TODO добавить проверку установки referer3;
//TODO добавить проверку установки ADV NA;

func TestAdvServices_AutoCreateAdv(t *testing.T) {

}

func TestAdvServices_GetAdv(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	allModels := models.NewModels(context.Background(), chClient)
	hitService := NewHit(context.Background(), allModels)
	advServices := NewAdv(context.Background(), allModels, hitService)

	t.Run("Указано 'Куда пришли'", func(t *testing.T) {
		utils.TruncateAllTables(chClient)
		var advUuid string
		chClient.Exec(context.Background(), "insert into adv (uuid, referer1, referer2, cost, events_view, description, priority) VALUES (generateUUIDv7(), 'rt_1', 'rt_2', 100.5, '', '', 1)")
		row := chClient.QueryRow(context.Background(), "select uuid from adv")
		row.Scan(&advUuid)
		chClient.Exec(context.Background(), "insert into adv_page(uuid, adv_uuid, page)\nVALUES (generateUUIDv7(),?,'localhost/catalog')", advUuid)

		referer, err := advServices.GetAdv(entityjson.UserData{
			PHPSessionId:      "",
			GuestUuid:         uuid.UUID{},
			Url:               "http://localhost/catalog",
			Referer:           "",
			Ip:                "",
			UserAgent:         "",
			UserId:            0,
			UserLogin:         "",
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "",
			Event1:            "",
			Event2:            "",
			IsUserAuth:        false,
			Method:            "",
			Cookies:           "",
			IsFavorite:        false,
		})
		req.NoError(err)
		req.Equal(referer.Referer1, "rt_1")
		req.Equal(referer.Referer2, "rt_2")
	})

	//t.Run("AddSearcherHit user agent exists", func(t *testing.T) {
	//	utils.TruncateTable("searcher_hit", chClient)
	//	err := advServices.AddHitSearcher(entityjson.UserData{
	//		PHPSessionId:      "",
	//		GuestUuid:         "",
	//		Url:               "https://test.local.com",
	//		Referer:           "",
	//		Ip:                "192.168.1.98",
	//		UserAgent:         "Abilon",
	//		UserId:            0,
	//		UserLogin:         "",
	//		HttpXForwardedFor: "",
	//		IsError404:        false,
	//		SiteId:            "mg",
	//		Event1:            "",
	//		Event2:            "",
	//		IsUserAuth:        false,
	//	},
	//	)
	//	req.NoError(err)
	//
	//	var searcher []entitydb.SearcherHitDb
	//	resultSql := `select uuid, date_hit, searcher_uuid, url, url_404, ip, user_agent, site_id from searcher_hit`
	//	err = chClient.Select(context.Background(), &searcher, resultSql)
	//	req.NoError(err)
	//
	//	req.Equal(1, len(searcher))
	//	req.Equal(searcher[0].Url, "https://test.local.com")
	//	req.Equal(searcher[0].Ip, "192.168.1.98")
	//	req.Equal(searcher[0].UserAgent, "Abilon")
	//	req.Equal(searcher[0].SearcherId, "0190d4cb-825b-7512-8008-efd0c75f0fbc")
	//	req.Equal(searcher[0].SiteId, "mg")
	//})

}
