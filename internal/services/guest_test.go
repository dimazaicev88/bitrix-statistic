package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

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

	hit, err := guestService.Add(entityjson.UserData{}, entitydb.AdvReferer{})
	req.Nil(err)
	req.Equal("session is empty", err.Error())
	req.Equal(hit, entitydb.Hit{})
}
