package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type SearcherService struct {
	ctx           context.Context
	chClient      driver.Conn
	SearcherModel *models.Searcher
}

func NewSearcherService(ctx context.Context, chClient driver.Conn) *SearcherService {
	return &SearcherService{
		ctx:           ctx,
		chClient:      chClient,
		SearcherModel: models.NewSearcher(ctx, chClient),
	}
}

// IsSearcher Проверка по user agent является ли guest поисковиком
func (ss SearcherService) IsSearcher(userAgent string) (bool, error) {
	if len(userAgent) == 0 {
		return false, nil
	}
	searcher, err := ss.SearcherModel.FindSearcherByUserAgent(userAgent)
	if err != nil {
		return false, err
	}
	return searcher != entitydb.SearcherDb{}, nil
}

func (ss SearcherService) AddStatData(data entityjson.StatData) error {
	searcher, err := ss.SearcherModel.FindSearcherByUserAgent(data.UserAgent)
	if err != nil {
		return err
	}

	if ss.SearcherModel.ExistStatDayForSearcher(searcher.Uuid) {
		err = ss.SearcherModel.UpdateSearcherDay(searcher.Uuid)
		if err != nil {
			return err
		}
	} else {
		err = ss.SearcherModel.AddSearcherDay(searcher.Uuid)
		if err != nil {
			return err
		}
	}

	return nil
}
