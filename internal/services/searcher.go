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

func NewSearcher(ctx context.Context, chClient driver.Conn) *SearcherService {
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

func (ss SearcherService) AddHitSearcher(data entityjson.StatData) error {
	searcher, err := ss.SearcherModel.FindSearcherByUserAgent(data.UserAgent)
	if err != nil {
		return err
	}

	if searcher == (entitydb.SearcherDb{}) {
		return nil
	}

	if err = ss.SearcherModel.AddHitSearcher(searcher.Uuid, data); err != nil {
		return err
	}

	if err = ss.SearcherModel.AddSearcherDayHits(searcher.Uuid); err != nil {
		return err
	}

	return nil
}
