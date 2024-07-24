package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type PhraseService struct {
	phraseModel *models.Phrase
}

func NewPhraseService(ctx context.Context, chClient driver.Conn) *PhraseService {
	return &PhraseService{
		phraseModel: models.NewPhrase(ctx, chClient),
	}
}

func (s *PhraseService) Filter(filter filters.Filter) ([]entitydb.PhraseListDB, error) {
	return s.phraseModel.Filter(filter)
}
