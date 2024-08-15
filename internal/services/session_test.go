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
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSessionService_Add(t *testing.T) {

}

// Добавление хита, но данные с сессией переданы пустые
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

	guest, err := sessionService.Add(entityjson.UserData{}, entitydb.AdvReferer{})
	req.NotNil(err)
	req.Equal("user data is empty", err.Error())
	req.Equal(guest, entitydb.Guest{})
}
