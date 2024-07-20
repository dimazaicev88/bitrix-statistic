package services

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSearcherService_AllTests(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()
	req := require.New(t)
	searchService := NewSearcherService(context.Background(), chClient)

	t.Run("IsSearcher method", func(t *testing.T) {
		utils.TruncateAllTables(chClient)

		searcher, err := searchService.IsSearcher("FeedDemon")
		req.NoError(err)
		req.True(searcher)

		searcher, err = searchService.IsSearcher("WFQXZFeedDemon")
		req.NoError(err)
		req.False(searcher)

		searcher, err = searchService.IsSearcher("")
		req.NoError(err)
		req.False(searcher)
	})

	t.Run("AddSearcherHit user agent exists", func(t *testing.T) {
		utils.TruncateTable("searcher_hit", chClient)
		err := searchService.AddHitSearcher(entityjson.StatData{
			PHPSessionId:      "",
			GuestHash:         "",
			Url:               "https://test.local.com",
			Referer:           "",
			Ip:                "192.168.1.98",
			UserAgent:         "Abilon",
			UserId:            0,
			UserLogin:         "",
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "mg",
			Event1:            "",
			Event2:            "",
			IsUserAuth:        false,
		},
		)
		req.NoError(err)

		var searcher []entitydb.SearcherHitDb
		resultSql := `select uuid, date_hit, searcher_uuid, url, url_404, ip, user_agent, site_id from searcher_hit`
		err = chClient.Select(context.Background(), &searcher, resultSql)
		req.NoError(err)

		req.Equal(1, len(searcher))
		req.Equal(searcher[0].Url, "https://test.local.com")
		req.Equal(searcher[0].Ip, "192.168.1.98")
		req.Equal(searcher[0].UserAgent, "Abilon")
		req.Equal(searcher[0].SearcherId, "0190c649-2821-7886-bddf-723a522c66d2")
		req.Equal(searcher[0].SiteId, "mg")
	})

	t.Run("AddSearcherHit user agent not exists", func(t *testing.T) {
		utils.TruncateTable("searcher_hit", chClient)
		err := searchService.AddHitSearcher(entityjson.StatData{
			PHPSessionId:      "",
			GuestHash:         "",
			Url:               "https://test.local.com",
			Referer:           "",
			Ip:                "192.168.1.98",
			UserAgent:         "_#2!2$%",
			UserId:            0,
			UserLogin:         "",
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "mg",
			Event1:            "",
			Event2:            "",
			IsUserAuth:        false,
		},
		)
		req.NoError(err)

		var searcher []entitydb.SearcherHitDb
		resultSql := `select uuid, date_hit, searcher_uuid, url, url_404, ip, user_agent, site_id from searcher_hit`
		err = chClient.Select(context.Background(), &searcher, resultSql)
		req.NoError(err)

		req.Equal(0, len(searcher))
	})

	t.Run("AddSearcherHit check searcher day values", func(t *testing.T) {
		utils.TruncateTable("searcher_hit", chClient)
		searchService.AddHitSearcher(entityjson.StatData{
			PHPSessionId:      "",
			GuestHash:         "",
			Url:               "https://test.local.com",
			Referer:           "",
			Ip:                "192.168.1.98",
			UserAgent:         "Abilon",
			UserId:            0,
			UserLogin:         "",
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "mg",
			Event1:            "",
			Event2:            "",
			IsUserAuth:        false,
		},
		)

		searchService.AddHitSearcher(entityjson.StatData{
			PHPSessionId:      "",
			GuestHash:         "",
			Url:               "https://test.local.com",
			Referer:           "",
			Ip:                "192.168.1.98",
			UserAgent:         "Abilon",
			UserId:            0,
			UserLogin:         "",
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "mg",
			Event1:            "",
			Event2:            "",
			IsUserAuth:        false,
		},
		)
		searchService.AddHitSearcher(entityjson.StatData{
			PHPSessionId:      "",
			GuestHash:         "",
			Url:               "https://test.local.com",
			Referer:           "",
			Ip:                "192.168.1.98",
			UserAgent:         "Abilon",
			UserId:            0,
			UserLogin:         "",
			HttpXForwardedFor: "",
			IsError404:        false,
			SiteId:            "mg",
			Event1:            "",
			Event2:            "",
			IsUserAuth:        false,
		},
		)

		var searcher []entitydb.SearcherDayDb
		resultSql := `select date_stat, date_last, searcher_uuid, total_hits from searcher_day`
		err := chClient.Select(context.Background(), &searcher, resultSql)
		req.NoError(err)

		req.Equal(1, len(searcher))
		req.Equal(uint64(3), searcher[0].TotalHits)
		req.Equal("0190c649-2821-7886-bddf-723a522c66d2", searcher[0].SearcherUuid.String())
		req.Equal(time.Now().Format("2006-01-02"), searcher[0].DateStat.Format("2006-01-02"))
	})
}
