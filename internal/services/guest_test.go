package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/repository"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// Добавление хита, но данные с сессией переданы пустые
func TestGuestService_Add_EmptyUserData(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	_ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	allModels := repository.NewGuest(chClient)
	guestService := NewGuest(allModels)
	utils.TruncateTable("guest", chClient)

	err := guestService.Add(context.Background(), models.Guest{
		GuestHash: utils.GetGuestMd5(dto.UserData{
			UserAgent:         "",
			HttpXForwardedFor: "",
			Ip:                "",
		}),
		DateAdd: time.Now(),
	})
	req.NotNil(err)
	req.Equal("user data is empty", err.Error())

}
