package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
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

func (ss SessionService) Filter(filter filters.Filter) ([]entitydb.Session, error) {
	return nil, nil
}

func (ss SessionService) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	return ss.allModels.Session.FindByPHPSessionId(phpSessionId)
}

func (ss SessionService) FindAll(skip uint32, limit uint32) ([]entitydb.Session, error) {
	return ss.allModels.Session.FindAll(skip, limit)
}

func (ss SessionService) ConvertToJSONSession(dbSession entitydb.Session) entityjson.Session {
	return entityjson.Session{
		Uuid:      dbSession.Uuid,
		GuestUuid: dbSession.GuestUuid,
		//NewGuest:     dbSession.IsNewGuest,
		//UserId:       dbSession.UserId,
		//UserAuth:     dbSession.IsUserAuth,
		//Events:       dbSession.Events,
		//Hits:         dbSession.Hits,
		//Favorites:    dbSession.Favorites,
		//UserAgent:    dbSession.UserAgent,
		//DateStat:     dbSession.DateStat,
		//DateFirst:    dbSession.DateFirst,
		//DateLast:     dbSession.DateLast,
		//IpFirst:      dbSession.IpFirst,
		//IpLast:       dbSession.IpLast,
		//FirstHitUuid: dbSession.FirstHitUuid,
		//LastHitUuid:  dbSession.LastHitUuid,
		//PhpSessionId: dbSession.PhpSessionId,
		//AdvUuid:      dbSession.AdvUuid,
		//AdvBack:      dbSession.AdvBack,
		//Referer1:     dbSession.Referer1,
		//Referer2:     dbSession.Referer2,
		//Referer3:     dbSession.Referer3,
		//StopListUuid: dbSession.StopListUuid.String(),
		//CountryId:    dbSession.CountryId,
		//FirstSiteId:  dbSession.FirstSiteId,
		//LastSiteId:   dbSession.LastSiteId,
		//CityId:       dbSession.CityId,
		//UrlFrom:      dbSession.UrlFrom,
		//UrlTo:        dbSession.UrlTo,
		//UrlTo404:     dbSession.UrlTo404,
		//UrlLast:      dbSession.UrlLast,
		//UrlLast404:   dbSession.UrlLast404,
	}
}

func (ss SessionService) ConvertToJSONListSession(dbSessions []entitydb.Session) []entityjson.Session {
	var sessions []entityjson.Session

	for _, dbHit := range dbSessions {
		sessions = append(sessions, ss.ConvertToJSONSession(dbHit))
	}
	return sessions
}
