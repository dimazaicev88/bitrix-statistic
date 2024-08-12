package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"time"
)

//TODO сделать рефакторинг.

type PathService struct {
	allModels        *models.Models
	ctx              context.Context
	pathCacheService *PathCacheService
	pathAdvService   *PathAdvService
}

func NewPath(
	ctx context.Context,
	allModels *models.Models,
	pathCacheService *PathCacheService,
	pathAdvService *PathAdvService,
) *PathService {
	return &PathService{
		ctx:              ctx,
		allModels:        allModels,
		pathCacheService: pathCacheService,
		pathAdvService:   pathAdvService,
	}
}

func (ps PathService) SavePath(siteId, sessionUuid, currentUrl, referer string, isError404 bool, advReferer entitydb.AdvReferer) error {

	if currentUrl == referer {
		return nil
	}

	lastPath, err := ps.pathCacheService.FindLastBySessionUuid(sessionUuid)
	if err != nil {
		return err
	}

	if lastPath.PathLastPage == currentUrl {
		return nil
	}

	var countAbnormal uint32
	if len(referer) == 0 {
		if lastPath != (entitydb.PathCache{}) {
			countAbnormal++
		}
	}

	var pathCache entitydb.PathCache
	if referer != "" {
		pathCache, err = ps.pathCacheService.FindByReferer(sessionUuid, referer)
		if err != nil {
			return err
		}
	} else {
		pathCache, err = ps.pathCacheService.FindBySession(sessionUuid)
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

	err = ps.allModels.PathCache.Add(entitydb.PathCache{
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

	path, err := ps.allModels.Path.FindByPathId(currentPathId, time.Now().Local().Format("2006-01-02"))
	if err != nil {
		return err
	}

	pageHash := utils.Crc32(pathCache.PathLastPage)
	if path == (entitydb.Path{}) {
		err = ps.allModels.Path.Add(entitydb.Path{
			PathId:          currentPathId,
			ParentPathId:    pathCache.PathId,
			Counter:         1,
			CounterAbnormal: countAbnormal,
			CounterFullPath: 1,
			Pages:           currentPathPages,
			FirstPage:       firstPage,
			FirstPageSiteId: firstPageSiteId,
			FirstPage404:    firstPage404,
			PrevPage:        pathCache.PathLastPage,
			PrevPageHash:    pageHash,
			LastPage:        currentUrl,
			LastPage404:     isError404,
			LastPageSiteId:  siteId,
			LastPageHash:    pageHash,
			Sign:            1,
			Version:         1,
		})
		if err != nil {
			return err
		}
	} else {
		newPath := path
		newPath.Counter += 1
		newPath.CounterFullPath += 1
		newPath.CounterAbnormal += 1

		err = ps.allModels.Path.Update(path, newPath)
		if err != nil {
			return err
		}
	}

	if pathCache.IsLastPage {
		previewPath, err := ps.allModels.Path.FindByPathId(pathCache.PathId, time.Now().Local().Format("2006-01-02"))
		if err != nil {
			return err
		}
		newPath := previewPath
		newPath.CounterFullPath -= 1
		err = ps.allModels.Path.Update(previewPath, newPath)
		if err != nil {
			return err
		}
		newPathCache := pathCache
		newPathCache.IsLastPage = false
		err = ps.allModels.PathCache.Update(pathCache, newPathCache)
		if err != nil {
			return err
		}
	}

	var advCounter uint32
	var advCounterBack uint32
	var advCounterFullPath uint32
	var advCounterFullPathBack uint32
	var advBack bool
	if len(advReferer.AdvUuid) > 0 && !advReferer.LastAdvBack {
		advCounter = 1
		advCounterBack = 0
		advCounterFullPath = 1
		advCounterFullPathBack = 0
		advBack = false
	} else if len(advReferer.AdvUuid) > 0 && advReferer.LastAdvBack {
		advCounter = 0
		advCounterBack = 1
		advCounterFullPath = 0
		advCounterFullPathBack = 1
		advBack = true
	} else {
		return nil
	}

	pathAdv, err := ps.pathAdvService.FindByPathId(currentPathId, time.Now().Local().Format("2006-01-02"))
	if err != nil {
		return err
	}
	if pathAdv == (entitydb.PathAdv{}) {
		err = ps.pathAdvService.Add(entitydb.PathAdv{
			AdvUuid:             advReferer.AdvUuid,
			PathId:              currentPathId,
			Counter:             advCounter,
			CounterBack:         advCounterBack,
			CounterFullPath:     advCounterFullPath,
			CounterFullPathBack: advCounterFullPathBack,
			Steps:               currentPathSteps,
			Sign:                1,
			Version:             1,
		})
		if err != nil {
			return err
		}
	} else {
		newPath := pathAdv
		newPath.Counter += advCounter
		newPath.CounterBack += advCounterBack
		newPath.CounterBack += advCounterBack
		newPath.CounterFullPath += advCounterFullPath
		newPath.CounterFullPathBack += advCounterFullPathBack
		ps.pathAdvService.Update(pathAdv, newPath)
	}

	return nil
}
