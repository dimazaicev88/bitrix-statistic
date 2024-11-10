package repository

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGuest(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("E:\\projects\\bitrix-statistic\\.env")
	if err != nil {
		logrus.Fatal("Error loading .env file", err.Error())
	}

	req := require.New(t)
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	guestRepo := NewGuest(chClient)

	t.Run("Add", func(t *testing.T) {
		utils.TruncateTable(ctx, "guests", chClient)
		guestHash := uuid.New().String()
		err = guestRepo.Add(ctx, models.Guest{GuestHash: guestHash, DateInsert: time.Now()})
		req.NoError(err)
		var guestDb models.Guest
		chClient.NewSelect().Model(&guestDb).Scan(ctx)
		req.Equal(guestHash, guestDb.GuestHash)
	})

	t.Run("Add", func(t *testing.T) {
		utils.TruncateTable(ctx, "guests", chClient)
		utils.TruncateTable(ctx, "guests", chClient)
		guestHash := uuid.New().String()

		chClient.NewInsert().Model(&models.Guest{GuestHash: guestHash, DateInsert: time.Now()}).Exec(ctx)
		guestDb, err := guestRepo.FindByHash(ctx, guestHash)
		req.NoError(err)
		req.Equal(guestHash, guestDb.GuestHash)
	})
}
