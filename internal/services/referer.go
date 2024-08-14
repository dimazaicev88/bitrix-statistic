package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/google/uuid"
	"net/url"
	"time"
)

type RefererService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewReferer(ctx context.Context, allModels *models.Models) *RefererService {
	return &RefererService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (rs RefererService) Find(filter filters.Filter) ([]entitydb.Referer, error) {
	return rs.allModels.Referer.Find(filter)
}

func (rs RefererService) Add(referer string) (string, error) {
	if referer == "" {
		return "", nil
	}
	return rs.allModels.Referer.Add(referer)
}

func (rs RefererService) AddToRefererList(advUuid, sessionUuid uuid.UUID, idReferer string, parsedUrl *url.URL, statData entityjson.UserData) (entitydb.RefererList, error) {
	refererList := entitydb.RefererList{
		Uuid:        uuid.New(),
		RefererId:   idReferer,
		DateHit:     time.Time{},
		Protocol:    parsedUrl.Scheme,
		SiteName:    parsedUrl.Hostname(),
		UrlFrom:     statData.Referer,
		UrlTo:       statData.Url,
		UrlTo404:    statData.IsError404,
		SessionUuid: sessionUuid,
		AdvUuid:     advUuid,
		SiteId:      statData.SiteId,
	}

	err := rs.allModels.Referer.AddToRefererList(refererList)
	if err != nil {
		return entitydb.RefererList{}, err
	}

	return refererList, err
}
