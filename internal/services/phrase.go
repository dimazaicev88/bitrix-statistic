package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type PhraseService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewPhraseService(ctx context.Context, allModels *models.Models) *PhraseService {
	return &PhraseService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (s *PhraseService) Filter(filter filters.Filter) ([]entitydb.PhraseList, error) {
	return s.allModels.Phrase.Filter(filter)
}
