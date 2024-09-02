package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPathAdvService_Add_EmptyData(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	pathAdvService := NewPathAdvService(context.Background(), models.NewModels(context.Background(), chClient))

	err := pathAdvService.Add(entitydb.PathAdv{})
	req.NotNil(err)
	req.Equal("path adv is empty", err.Error())
}

func TestPathAdvService_Add(t *testing.T) {
	req := require.New(t)
	if err := godotenv.Load(pathToEnvFile); err != nil {
		req.Fail("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	allModels := models.NewModels(context.Background(), chClient)
	pathAdvService := NewPathAdvService(context.Background(), allModels)

	utils.TruncateTable("path_adv", chClient)

	pathAdv := entitydb.PathAdv{
		AdvUuid:             uuid.New(),
		PathId:              1,
		Counter:             2,
		CounterBack:         3,
		CounterFullPath:     4,
		CounterFullPathBack: 5,
		Steps:               6,
	}

	err := pathAdvService.Add(pathAdv)
	req.Nil(err)

	var allDbPathAdv []entitydb.PathAdv
	rows, _ := chClient.Query(context.Background(), "SELECT * from path_adv")

	for rows.Next() {
		var dbHit entitydb.PathAdv
		rows.ScanStruct(&dbHit)
		allDbPathAdv = append(allDbPathAdv, dbHit)
	}
	req.Equal(1, len(allDbPathAdv))

	req.Equal(pathAdv.AdvUuid, allDbPathAdv[0].AdvUuid)
	req.Equal(pathAdv.PathId, allDbPathAdv[0].PathId)
	req.Equal(pathAdv.Counter, allDbPathAdv[0].Counter)
	req.Equal(pathAdv.CounterBack, allDbPathAdv[0].CounterBack)
	req.Equal(pathAdv.CounterFullPath, allDbPathAdv[0].CounterFullPath)
	req.Equal(pathAdv.CounterFullPathBack, allDbPathAdv[0].CounterFullPathBack)
	req.Equal(pathAdv.Steps, allDbPathAdv[0].Steps)
	req.Equal(int8(1), allDbPathAdv[0].Sign)
	req.Equal(uint32(1), allDbPathAdv[0].Version)
}
