package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
)

type SearcherService struct {
	ctx       context.Context
	allModels *models.Models
}

func NewSearcher(ctx context.Context, allModels *models.Models) *SearcherService {
	return &SearcherService{
		ctx:       ctx,
		allModels: allModels,
	}
}

// IsSearcher Проверка по user agent является ли guest поисковиком
func (ss SearcherService) IsSearcher(userAgent string) (bool, error) {
	if len(userAgent) == 0 {
		return false, nil
	}
	searcher, err := ss.allModels.Searcher.FindSearcherByUserAgent(userAgent)
	if err != nil {
		return false, err
	}
	return searcher != entitydb.Searcher{}, nil
}

func (ss SearcherService) AddHitSearcher(data entityjson.StatData) error {
	searcher, err := ss.allModels.Searcher.FindSearcherByUserAgent(data.UserAgent)
	if err != nil {
		return err
	}

	if searcher == (entitydb.Searcher{}) {
		return nil
	}

	if err = ss.allModels.Searcher.AddHitSearcher(searcher.Uuid, data); err != nil {
		return err
	}

	if err = ss.allModels.Searcher.AddSearcherDayHits(searcher.Uuid); err != nil {
		return err
	}

	return nil
}
