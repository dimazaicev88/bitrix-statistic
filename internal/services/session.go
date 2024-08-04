package services

import (
	"bitrix-statistic/internal/entitydb"
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

func (ss SessionService) Add(guestUuid, phpSessionId string) (string, error) {
	sessionUuid, err := ss.allModels.Session.Add(entitydb.Session{
		//Uuid:         sessionUuid.String(),
		//GuestUuid:    guestUuid,
		//PhpSessionId: phpSessionId,
	})

	if err != nil {
		return "", err
	}

	return sessionUuid, nil
}

func (ss SessionService) IsExistsSession(phpSession string) bool {
	count, err := ss.allModels.Session.ExistsByPhpSession(phpSession)
	if err != nil {
		return false
	}
	return count > 0
}

//func (ss SessionService) UpdateSession(data entityjson.StatData) error {
//	err := ss.allModels.Session.Update(data)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (ss SessionService) Filter(filter filters.Filter) ([]entitydb.Session, error) {
	return nil, nil
}

func (ss SessionService) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	return ss.allModels.Session.FindByPHPSessionId(phpSessionId)
}

func (ss SessionService) Update(oldSession entitydb.Session, newSession entitydb.Session) error {
	return ss.allModels.Session.Update(oldSession, newSession)
}
