package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
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

func (ps PathService) SavePath(siteId, sessionUuid, currentUrl, referer string, isError404 bool) error {

	if currentUrl == referer {
		return nil
	}

	lastPath, err := ps.FindLastBySessionUuid(sessionUuid)
	if err != nil {
		return err
	}

	if lastPath.PathLastPage == currentUrl {
		return nil
	}

	countAbnormal := 0
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
		if err != nil {
			return err
		}
	}

	currentPathId := utils.Crc32(utils.StringConcat(currentUrl, string(pathCache.PathId)))
	tmpSiteId := ""

	if siteId != "" {
		tmpSiteId = utils.StringConcat("[", siteId, "]")
	}

	var currentPathPages string

	if isError404 {
		currentPathPages = utils.StringConcat(pathCache.PathPages, tmpSiteId, "ERROR_404:", currentUrl, "\n")
	} else {
		currentPathPages = utils.StringConcat(pathCache.PathPages, tmpSiteId, currentUrl, "\n")
	}

	currentPathSteps := pathCache.PathSteps + 1

	var firstPage404 bool
	var firstPage string
	var firstPageSiteId string
	if pathCache.PathFirstPage != "" {
		firstPage = pathCache.PathFirstPage
		firstPageSiteId = pathCache.PathFirstPageSiteId
		if pathCache.PathFirstPage404 {
			firstPage404 = true
		} else {
			firstPage404 = false
		}
	} else {
		firstPage = currentUrl
		firstPageSiteId = siteId
		firstPage404 = isError404
	}

	err = ps.allModels.Path.AddPathCache(entitydb.PathCache{
		SessionUuid:         sessionUuid,
		PathId:              currentPathId,
		PathPages:           currentPathPages,
		PathFirstPage:       firstPage,
		PathFirstPage404:    firstPage404,
		PathFirstPageSiteId: firstPageSiteId,
		PathLastPage:        currentUrl,
		PathLastPage404:     isError404,
		PathLastPageSiteId:  siteId,
		PathSteps:           currentPathSteps,
		IsLastPage:          true,
	})

	if err != nil {
		return err
	}

	ps.allModels.Path.AddPath()

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
