package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
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

func (ss SessionService) Add(sessionUuid, guestUuid, hitUuid uuid.UUID, phpSessionId string) (entitydb.Session, error) {

	switch {
	case guestUuid == uuid.Nil:
		return entitydb.Session{}, errors.New("guestUuid is empty")
	case hitUuid == uuid.Nil:
		return entitydb.Session{}, errors.New("hitUuid is empty")
	}

	var sessionDb entitydb.Session
	sessionDb.Uuid = sessionUuid
	sessionDb.GuestUuid = guestUuid
	sessionDb.PhpSessionId = phpSessionId
	err := ss.allModels.Session.Add(sessionDb)

	if err != nil {
		return entitydb.Session{}, err
	}

	return sessionDb, nil
}

func (ss SessionService) IsExistsByPhpSession(phpSession string) bool {
	exists, err := ss.allModels.Session.ExistsByPhpSession(phpSession)
	if err != nil {
		return false
	}
	return exists
}

func (ss SessionService) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	return ss.allModels.Session.FindByPHPSessionId(phpSessionId)
}
