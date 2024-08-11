package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
)

type PathService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewPath(ctx context.Context, allModels *models.Models) *PathService {
	return &PathService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (ps PathService) SavePath(sessionUuid, currentUrl, referer string) error {

	if currentUrl == referer {
		return nil
	}

	countAbnormal := 0
	lastPath, err := ps.FindLastBySessionUuid(sessionUuid)
	if err != nil {
		return err
	}

	if lastPath.PathLastPage == currentUrl {
		return nil
	}

	if len(referer) == 0 {
		if lastPath != (entitydb.PathCache{}) {
			countAbnormal++
		}
	}

	var pathCache entitydb.PathCache
	if referer != "" {
		pathCache, err = ps.FindByReferer(sessionUuid, referer)
		if err != nil {
			return err
		}
	} else {
		pathCache, err = ps.FindBySession(sessionUuid)
	}

	ps.allModels.Path.AddCache(entitydb.PathCache{
		SessionUuid:         sessionUuid,
		PathUuid:            "",
		PathPages:           "",
		PathFirstPage:       "",
		PathFirstPage404:    false,
		PathFirstPageSiteId: "",
		PathLastPage:        "",
		PathLastPage404:     false,
		PathLastPageSiteId:  "",
		PathSteps:           0,
		IsLastPage:          false,
	})

	return nil
}

func (ps PathService) FindLastBySessionUuid(uuid string) (entitydb.PathCache, error) {
	return ps.allModels.Path.FindLastBySessionUuid(uuid)
}

func (ps PathService) FindByReferer(uuid string, referer string) (entitydb.PathCache, error) {
	return ps.allModels.Path.FindByReferer(uuid, referer)
}

func (ps PathService) FindBySession(uuid string) (entitydb.PathCache, error) {
	return ps.allModels.Path.FindBySession(uuid)
}
