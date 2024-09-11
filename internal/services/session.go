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

func (ss *SessionService) Add(stopListUuid, guestUuid, hitUuid uuid.UUID, existGuest bool, userData entityjson.UserData, adv entitydb.AdvCompany) (entitydb.Session, error) {

	switch {
	case guestUuid == uuid.Nil:
		return entitydb.Session{}, errors.New("guestUuid is empty")
	case hitUuid == uuid.Nil:
		return entitydb.Session{}, errors.New("hitUuid is empty")
	case userData == (entityjson.UserData{}):
		return entitydb.Session{}, errors.New("userData is empty")
	}

	var sessionDb entitydb.Session
	sessionDb.Uuid = uuid.New()
	sessionDb.GuestUuid = guestUuid
	sessionDb.PhpSessionId = userData.PHPSessionId
	sessionDb.IsNewGuest = existGuest
	sessionDb.UserId = userData.UserId
	sessionDb.Favorites = userData.IsFavorite
	sessionDb.UrlFrom = userData.Referer
	sessionDb.UrlTo = userData.Url
	sessionDb.UrlTo404 = userData.IsError404
	sessionDb.UrlLast = userData.Url
	sessionDb.UrlLast404 = userData.IsError404
	sessionDb.UserAgent = userData.UserAgent
	sessionDb.IsUserAuth = userData.IsUserAuth
	sessionDb.DateStat = time.Now()
	sessionDb.DateFirst = time.Now()
	sessionDb.DateLast = time.Now()
	sessionDb.IpFirst = userData.Ip
	sessionDb.IpLast = userData.Ip
	sessionDb.FirstHitUuid = hitUuid
	sessionDb.LastHitUuid = hitUuid
	sessionDb.AdvUuid = adv.AdvUuid
	sessionDb.AdvBack = adv.LastAdvBack
	sessionDb.Referer1 = adv.Referer1
	sessionDb.Referer2 = adv.Referer2
	sessionDb.Referer3 = adv.Referer3
	sessionDb.FirstSiteId = userData.SiteId
	sessionDb.LastSiteId = userData.SiteId
	sessionDb.StopListUuid = stopListUuid
	sessionDb.Sign = 1
	sessionDb.Version = 1

	err := ss.allModels.Session.Add(sessionDb)

	if err != nil {
		return entitydb.Session{}, err
	}

	return sessionDb, nil
}

func (ss *SessionService) IsExistsByPhpSession(phpSession string) bool {
	exists, err := ss.allModels.Session.ExistsByPhpSession(phpSession)
	if err != nil {
		return false
	}
	return exists
}

func (ss *SessionService) Filter(filter filters.Filter) ([]entitydb.Session, error) {
	return nil, nil
}

func (ss *SessionService) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	return ss.allModels.Session.FindByPHPSessionId(phpSessionId)
}

func (ss *SessionService) Update(oldSession entitydb.Session, newSession entitydb.Session) error {
	oldSession.Sign *= -1
	newSession.Sign *= 1
	newSession.Version += 1
	return ss.allModels.Session.Update(oldSession, newSession)
}
