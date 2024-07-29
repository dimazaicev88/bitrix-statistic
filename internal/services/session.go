package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type SessionService struct {
	ctx       context.Context
	allModels *models.Models
}

func NewSession(ctx context.Context, allModels *models.Models) *SessionService {
	return &SessionService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (ss SessionService) Add(guestUuid, phpSessionId string) error {
	err := ss.allModels.Session.Add(entitydb.Session{
		GuestUuid:    guestUuid,
		PhpSessionId: phpSessionId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ss SessionService) IsExistsSession(phpSession string) bool {
	count, err := ss.allModels.Session.ExistsByPhpSession(phpSession)
	if err != nil {
		return false
	}
	return count > 0
}

func (ss SessionService) UpdateSession(data entityjson.StatData) error {
	err := ss.allModels.Session.Update(data)
	if err != nil {
		return err
	}
	return nil
}

func (ss SessionService) Filter(filter filters.Filter) ([]entitydb.SessionStat, error) {
	return nil, nil
}
