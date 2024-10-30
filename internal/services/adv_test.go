package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
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
	advServices := NewAdv(context.Background(), allModels)
	advServices.SetHitService(NewHit(context.Background(), allModels))
	advServices.SetOptionService(NewOption(context.Background(), allModels))

	//TODO добавить проверку метода Find
	t.Run("Указано 'Куда пришли'", func(t *testing.T) {
		utils.TruncateAllTables(chClient)
		var advUuid string
		chClient.Exec(context.Background(), "insert into adv (uuid, referer1, referer2, cost, events_view, description, priority) VALUES (generateUUIDv7(), 'rt_1', 'rt_2', 100.5, '', '', 1)")
		row := chClient.QueryRow(context.Background(), "select uuid from adv")
		row.Scan(&advUuid)
		chClient.Exec(context.Background(), "insert into adv_page(uuid, adv_uuid, page)\nVALUES (generateUUIDv7(),?,'localhost/catalog')", advUuid)

		referer, err := advServices.GetAdv(dto.UserData{
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

	t.Run("GetAdv. Find adv", func(t *testing.T) {
		utils.TruncateTable("searcher_hit", chClient)
		advReferer, err := advServices.GetAdv(dto.UserData{
			PHPSessionId: "dqweqasdadasd",
			GuestUuid:    uuid.New(),
		},
		)
		req.NoError(err)
		req.NotEqual(advReferer, entitydb.AdvCompany{})
	})

}
