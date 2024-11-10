package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/storage"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"testing"
)

// Добавление хита
func TestHitService_Add(t *testing.T) {
	//if err := godotenv.Load(pathToEnvFile); err != nil {
	//	logrus.Fatal("Error loading .env file")
	//}
	//
	//chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	//defer chClient.Close()
	//
	//req := require.New(t)
	//hitModel := repository.NewModels(context.Background(), chClient)
	//hitService := NewHit(context.Background(), hitModel)
	//utils.TruncateTable("hitModel", chClient)
	//guestUuid := uuid.New()
	//sessionDb := models.Session{
	//	Uuid:      uuid.New(),
	//	GuestUuid: guestUuid,
	//	//IsNewGuest:   true,
	//	//UserId:       10,
	//	//IsUserAuth:   true,
	//	//Events:       3,
	//	//Hits:         5,
	//	//Favorites:    true,
	//	//UrlFrom:      "https://google.com/",
	//	//UrlTo:        "http://localhost/",
	//	//UrlTo404:     false,
	//	//UrlLast:      "",
	//	//UrlLast404:   false,
	//	//UserAgent:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
	//	//DateStat:     time.Now(),
	//	//DateFirst:    time.Now(),
	//	//DateLast:     time.Now().Add(5 + time.Hour),
	//	//IpFirst:      "10.136.254.102",
	//	//IpLast:       "10.136.252.152",
	//	//FirstHitUuid: uuid.New(),
	//	//LastHitUuid:  uuid.New(),
	//	PhpSessionId: "b59c67bf196a4758191e42f76670ceba",
	//	//AdvUuid:      uuid.New(),
	//	//AdvBack:      false,
	//	//Referer1:     "ref1",
	//	//Referer2:     "ref2",
	//	//Referer3:     "ref3",
	//	//StopListUuid: uuid.New(),
	//	//CountryId:    "",
	//	//FirstSiteId:  "",
	//	//LastSiteId:   "",
	//	//CityId:       "",
	//	//Sign:         1,
	//	//Version:      1,
	//}
	//
	////advReferer := models.AdvCompany{
	////	AdvUuid:     uuid.New(),
	////	Referer1:    "ref1",
	////	Referer2:    "ref2",
	////	Referer3:    "ref3",
	////	LastAdvBack: false,
	////}
	//
	//statData := dto.UserData{
	//	PHPSessionId:      "b59c67bf196a4758191e42f76670ceba",
	//	GuestUuid:         guestUuid,
	//	Url:               "ttp://localhost/",
	//	Referer:           "https://google.com/",
	//	Ip:                "10.136.252.152",
	//	UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
	//	UserId:            10,
	//	UserLogin:         "admin",
	//	HttpXForwardedFor: "",
	//	IsError404:        false,
	//	SiteId:            "s1",
	//	Event1:            "",
	//	Event2:            "",
	//	IsUserAuth:        true,
	//	Method:            "GET",
	//	Cookies:           "_ga=GA1.2.131634531.1723559453; _gid=GA1.2.1901415454.1723559453; _ga_PWTK27XVWP=GS1.1.1723559453.1.1.1723559817.0.0.0",
	//	IsFavorite:        false,
	//}
	////hitModel, err := hitService.Add(true, sessionDb, advReferer, statData)
	////req.Nil(err)
	//
	//var allDbHits []models.Hit
	//rows, _ := chClient.Query(context.Background(), "SELECT * from hitModel")
	//
	//for rows.Next() {
	//	var dbHit models.Hit
	//	rows.ScanStruct(&dbHit)
	//	allDbHits = append(allDbHits, dbHit)
	//}
	//req.Equal(1, len(allDbHits))

	//req.Equal(hitModel.Uuid.String(), allDbHits[0].Uuid.String())
	//req.Equal(hitModel.SessionUuid.String(), allDbHits[0].SessionUuid.String())
	//req.Equal(hitModel.AdvUuid.String(), allDbHits[0].AdvUuid.String())
	//req.Equal(hitModel.GuestUuid.String(), allDbHits[0].GuestUuid.String())
	//req.Equal(hitModel.IsNewGuest, allDbHits[0].IsNewGuest)
	//req.Equal(hitModel.PhpSessionId, allDbHits[0].PhpSessionId)
	//req.Equal(hitModel.UserId, allDbHits[0].UserId)
	//req.Equal(hitModel.IsUserAuth, allDbHits[0].IsUserAuth)
	//req.Equal(hitModel.Url, allDbHits[0].Url)
	//req.Equal(hitModel.Url404, allDbHits[0].Url404)
	//req.Equal(hitModel.UrlFrom, allDbHits[0].UrlFrom)
	//req.Equal(hitModel.Method, allDbHits[0].Method)
	//req.Equal(hitModel.Cookies, allDbHits[0].Cookies)
	//req.Equal(hitModel.UserAgent, allDbHits[0].UserAgent)
	//req.Equal(hitModel.StopListUuid, allDbHits[0].StopListUuid)
	//req.Equal(strings.Trim(" ", hitModel.CountryId), strings.Trim(" ", allDbHits[0].CountryId))
	//req.Equal(hitModel.CityUuid.String(), allDbHits[0].CityUuid.String())
	//req.Equal(hitModel.SiteId, allDbHits[0].SiteId)
}

// Добавление хита, но данные с сессией переданы пустые
func TestHitService_Add_EmptySessRefererAndStatData(t *testing.T) {
	//if err := godotenv.Load(pathToEnvFile); err != nil {
	//	logrus.Fatal("Error loading .env file")
	//}
	//chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	//defer chClient.Close()
	//
	//req := require.New(t)
	//hitModel := repository.NewModels(context.Background(), chClient)
	//hitService := NewHit(context.Background(), hitModel)
	//utils.TruncateTable("hitModel", chClient)
	//
	//hitModel, err := hitService.Add(true, models.Session{}, models.AdvCompany{}, dto.UserData{})
	//req.NotNil(err)
	//req.Equal("session is empty", err.Error())
	//req.Equal(hitModel, models.Hit{})
}

// Добавление хита, но данные с данными о госте переданы пустые
func TestHitService_Add_EmptyStatData(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	_ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	//req := require.New(t)
	//hitModel := repository.NewModels(context.Background(), chClient)
	//hitService := NewHit(context.Background(), hitModel)
	//utils.TruncateTable("hitModel", chClient)
	//sessionDb := models.Session{Uuid: uuid.New(),
	//	GuestUuid: uuid.New(),
	//}

	//hitModel, err := hitService.Add(true, sessionDb, models.AdvCompany{}, dto.UserData{})
	//req.NotNil(err)
	//req.Equal("stat data is empty", err.Error())
	//req.Equal(hitModel, models.Hit{})
}
