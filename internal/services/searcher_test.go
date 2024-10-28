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
	"time"
)

func TestSearcherService_AllTests(t *testing.T) {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	chClient, _ := storage.NewClickHouseClient(config.GetServerConfig())
	defer chClient.Close()

	searchService := NewSearcher(context.Background(), models.NewModels(context.Background(), chClient))

	t.Run("IsSearcher method", func(t *testing.T) {
		req := require.New(t)
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
		req := require.New(t)
		utils.TruncateTable("searcher_hit", chClient)
		err := searchService.AddHitSearcher(dto.UserData{
			PHPSessionId:      "",
			GuestUuid:         uuid.New(),
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

		var searcher []entitydb.SearcherHit
		resultSql := `select uuid, date_hit, searcher_uuid, url, url_404, ip, user_agent, site_id from searcher_hit`
		err = chClient.Select(context.Background(), &searcher, resultSql)
		req.NoError(err)

		req.Equal(1, len(searcher))
		req.Equal(searcher[0].Url, "https://test.local.com")
		req.Equal(searcher[0].Ip, "192.168.1.98")
		req.Equal(searcher[0].UserAgent, "Abilon")
		req.Equal(searcher[0].SearcherId, "0190d4cb-825b-7512-8008-efd0c75f0fbc")
		req.Equal(searcher[0].SiteId, "mg")
	})

	t.Run("AddSearcherHit user agent not exists", func(t *testing.T) {
		req := require.New(t)
		utils.TruncateTable("searcher_hit", chClient)
		err := searchService.AddHitSearcher(dto.UserData{
			PHPSessionId:      "",
			GuestUuid:         uuid.New(),
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

		var searcher []entitydb.SearcherHit
		resultSql := `select uuid, date_hit, searcher_uuid, url, url_404, ip, user_agent, site_id from searcher_hit`
		err = chClient.Select(context.Background(), &searcher, resultSql)
		req.NoError(err)

		req.Equal(0, len(searcher))
	})

	t.Run("AddSearcherHit check searcher day values", func(t *testing.T) {
		req := require.New(t)
		utils.TruncateTable("searcher_hit", chClient)
		utils.TruncateTable("searcher_day_hits", chClient)
		searchService.AddHitSearcher(dto.UserData{
			PHPSessionId:      "",
			GuestUuid:         uuid.New(),
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

		searchService.AddHitSearcher(dto.UserData{
			PHPSessionId:      "",
			GuestUuid:         uuid.New(),
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
		searchService.AddHitSearcher(dto.UserData{
			PHPSessionId:      "",
			GuestUuid:         uuid.New(),
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
		chClient.Exec(context.Background(), "OPTIMIZE TABLE searcher_day_hits DEDUPLICATE;")
		var searcher []entitydb.SearcherDayHits
		resultSql := `select hit_day, searcher_uuid, total_hits from searcher_day_hits`
		err := chClient.Select(context.Background(), &searcher, resultSql)
		req.NoError(err)

		req.Equal(1, len(searcher))
		req.Equal(uint64(3), searcher[0].TotalHits)
		req.Equal("0190d4cb-825b-7512-8008-efd0c75f0fbc", searcher[0].SearcherUuid.String())
		req.Equal(time.Now().Format("2006-01-02"), searcher[0].DateStat.Format("2006-01-02"))
	})
}
