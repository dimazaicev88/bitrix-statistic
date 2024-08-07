package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
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

func (ss SessionService) Add(guestUuid, hitUuid string, existGuest bool, statData entityjson.StatData, adv entitydb.AdvReferer) (string, error) {
	var sessionDb entitydb.Session
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

	sessionUuid, err := ss.allModels.Session.Add(sessionDb)

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
