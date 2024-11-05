package services

import (
	"testing"
)

const pathToEnvFile = "/home/zajtsev/projects/bitrix-statistic/.env"

func TestGuestModel_Searcher(t *testing.T) {
	//if err := godotenv.Load(pathToEnvFile); err != nil {
	//	logrus.Fatal("Error loading .env file")
	//}
	//chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	////utils.TruncateAllTables(chClient)
	//defer chClient.Close()
	//req := require.New(t)
	//err := NewStatistic(context.Background(), repository.NewModels(context.Background(), chClient)).Add(dto.StatData{
	//	PHPSessionId:      "te2ctj3n1nt6c2ci5l0era5di2",
	//	GuestUuid:         "44c2870053b0a6378f5db40c96406f00",
	//	Url:               "http://localhost/catalog/dresses/dress-fashionista-on-a-walk/",
	//	Referer:           "",
	//	Ip:                "127.0.0.1",
	//	UserAgent:         "CLX Bot",
	//	UserId:            1,
	//	HttpXForwardedFor: "",
	//	IsError404:        false,
	//	SiteId:            "s1",
	//})
	//req.NoError(err)
}
