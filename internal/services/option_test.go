package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOptionService_GetValues(t *testing.T) {

}

func TestOptionService_Add(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	optionService := NewOption(context.Background(), models.NewModels(context.Background(), chClient))
	err := optionService.Add(entitydb.Option{
		Name:        "test_name",
		Value:       10,
		Description: "test description",
		SiteId:      "ts",
	})
	req.NoError(err)
}

func TestOptionService_Set(t *testing.T) {

}
