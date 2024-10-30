package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
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
func (ss *SearcherService) IsSearcher(userAgent string) (bool, error) {
	if len(userAgent) == 0 {
		return false, nil
	}
	searcher, err := ss.allModels.Searcher.FindSearcherByUserAgent(userAgent)
	if err != nil {
		return false, err
	}
	return searcher != entitydb.Searcher{}, nil
}

func (ss *SearcherService) AddHitSearcher(data dto.UserData) error {
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

func (ss *SearcherService) FindSearcherParams(host string) (entitydb.SearcherParams, error) {
	return ss.allModels.Searcher.FindSearcherParamsByHost(host)
}

func (ss *SearcherService) AddPhraseList(list entitydb.PhraseList) error {
	return ss.allModels.Searcher.AddPhraseList(list)
}

func (ss *SearcherService) AddSearcherPhraseStat(searcherPhraseStat entitydb.SearcherPhraseStat) error {
	return ss.allModels.Searcher.AddSearcherPhraseStat(searcherPhraseStat)
}

func (ss *SearcherService) Find(filter filters.Filter) (any, error) {
	return ss.allModels.Searcher.Find(filter)
}

func (ss *SearcherService) FindDomainList(filter filters.Filter) (any, error) {
	return nil, nil
}

func (ss *SearcherService) FindDynamicList(filter filters.Filter) (any, error) {
	return nil, nil
}
