package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
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

func (ss SessionService) Add(stopListUuid, guestUuid, hitUuid uuid.UUID, existGuest bool, statData entityjson.UserData, adv entitydb.AdvReferer) (entitydb.Session, error) {

	switch {
	case guestUuid == uuid.Nil:
		return entitydb.Session{}, errors.New("guestUuid is empty")
	case hitUuid == uuid.Nil:
		return entitydb.Session{}, errors.New("hitUuid is empty")
	case statData == (entityjson.UserData{}):
		return entitydb.Session{}, errors.New("statData is empty")
	}

	var sessionDb entitydb.Session
	sessionDb.Uuid = uuid.New()
	sessionDb.GuestUuid = guestUuid
	sessionDb.PhpSessionId = statData.PHPSessionId
	sessionDb.IsNewGuest = existGuest
	sessionDb.UserId = statData.UserId
	sessionDb.Favorites = statData.IsFavorite
	sessionDb.UrlFrom = statData.Referer
	sessionDb.UrlTo = statData.Url
	sessionDb.UrlTo404 = statData.IsError404
	sessionDb.UrlLast = statData.Url
	sessionDb.UrlLast404 = statData.IsError404
	sessionDb.UserAgent = statData.UserAgent
	sessionDb.DateStat = time.Now()
	sessionDb.DateFirst = time.Now()
	sessionDb.DateLast = time.Now()
	sessionDb.IpFirst = statData.Ip
	sessionDb.IpLast = statData.Ip
	sessionDb.FirstHitUuid = hitUuid
	sessionDb.LastHitUuid = hitUuid
	sessionDb.AdvUuid = adv.AdvUuid
	sessionDb.AdvBack = adv.LastAdvBack
	sessionDb.Referer1 = adv.Referer1
	sessionDb.Referer2 = adv.Referer2
	sessionDb.Referer3 = adv.Referer3
	sessionDb.FirstSiteId = statData.SiteId
	sessionDb.LastSiteId = statData.SiteId
	sessionDb.StopListUuid = stopListUuid

	err := ss.allModels.Session.Add(sessionDb)

	if err != nil {
		return entitydb.Session{}, err
	}

	return sessionDb, nil
}

func (ss SessionService) IsExistsSession(phpSession string) bool {
	count, err := ss.allModels.Session.ExistsByPhpSession(phpSession)
	if err != nil {
		return false
	}
	return count > 0
}

func (ss SessionService) Filter(filter filters.Filter) ([]entitydb.Session, error) {
	return nil, nil
}

func (ss SessionService) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	return ss.allModels.Session.FindByPHPSessionId(phpSessionId)
}

func (ss SessionService) Update(oldSession entitydb.Session, newSession entitydb.Session) error {
	return ss.allModels.Session.Update(oldSession, newSession)
}
