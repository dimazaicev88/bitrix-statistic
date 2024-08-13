package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"time"
)

type PathService struct {
	allModels        *models.Models
	ctx              context.Context
	pathCacheService *PathCacheService
	pathAdvService   *PathAdvService
	optionService    *OptionService
	pageService      *PageService
}

func NewPath(
	ctx context.Context,
	allModels *models.Models,
	pathCacheService *PathCacheService,
	pathAdvService *PathAdvService,
	optionService *OptionService,
	pageService *PageService,
) *PathService {
	return &PathService{
		ctx:              ctx,
		allModels:        allModels,
		pathCacheService: pathCacheService,
		pathAdvService:   pathAdvService,
		optionService:    optionService,
		pageService:      pageService,
	}
}

// SavePath TODO рефакторинг.
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

	currentDay := time.Now().Local().Format("2006-01-02")

	path, err := ps.allModels.Path.FindByPathId(currentPathId, currentDay)
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

		if err = ps.allModels.Path.Update(path, newPath); err != nil {
			return err
		}
	}

	if pathCache.IsLastPage {
		previewPath, err := ps.allModels.Path.FindByPathId(pathCache.PathId, currentDay)
		if err != nil {
			return err
		}
		newPath := previewPath
		newPath.CounterFullPath -= 1

		if err = ps.allModels.Path.Update(previewPath, newPath); err != nil {
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

	pathAdv, err := ps.pathAdvService.FindByPathId(currentPathId, currentDay)
	if err != nil {
		return err
	}
	if pathAdv == (entitydb.PathAdv{}) {
		if currentPathSteps <= ps.optionService.MaxPathSteps(siteId) {
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
		}
	} else {
		newPathAdv := pathAdv
		newPathAdv.Counter += advCounter
		newPathAdv.CounterBack += advCounterBack
		newPathAdv.CounterBack += advCounterBack
		newPathAdv.CounterFullPath += advCounterFullPath
		newPathAdv.CounterFullPathBack += advCounterFullPathBack

		if err = ps.pathAdvService.Update(pathAdv, newPathAdv); err != nil {
			return err
		}
	}

	pathAdv, err = ps.pathAdvService.FindByPathId(currentPathId, currentDay)
	newPathAdv := pathAdv
	if pathCache.IsLastPage {
		if advBack {
			newPathAdv.CounterFullPath -= 1

			if err = ps.pathAdvService.Update(pathAdv, newPathAdv); err != nil {
				return err
			}
		} else {
			newPathAdv.CounterFullPathBack -= 1
		}
	}

	return nil
}

func (ps PathService) SaveVisits(siteId, sessionNew,
	currentDir, currentPage, lastDir, lastDirUuid,
	lastPage, lastPageUuid string, isError404 bool,
	adv entitydb.AdvReferer,
) error {

	if len(currentPage) == 0 && len(currentDir) == 0 {
		return nil
	}
	var currentDirUuid string
	var currentPageUuid string
	var exitDirCounter uint32
	var exitPageCounter uint32

	if len(lastDir) == 0 || lastDir != currentDir || len(lastPage) == 0 || lastPage != currentPage {

		pages, err := ps.pageService.FindByPageAndDir(currentPage, currentPage, utils.GetCurrentDate())
		if err != nil {
			return err
		}

		for _, page := range pages {
			if page.Dir {
				currentDirUuid = page.Uuid
			} else {
				currentPageUuid = page.Uuid
			}
		}
		if currentDirUuid != lastDirUuid {
			exitDirCounter = 1
		}
		if currentPageUuid != lastPageUuid {
			exitPageCounter = 1
		}
	} else {
		currentDirUuid = lastDirUuid
		currentPageUuid = lastPageUuid
	}

	if len(lastDirUuid) > 0 && exitDirCounter > 0 {
		pages, err := ps.pageService.FindByUuid(lastDirUuid)
		if err != nil {
			return err
		}
		newPages := pages
		newPages.ExitCounter -= 1

		if err = ps.pageService.Update(pages, newPages); err != nil {
			return err
		}

		if len(adv.AdvUuid) > 0 && adv.LastAdvBack {
			pathAdv, err := ps.pathAdvService.FindByPageAndAdvUuid(lastDirUuid, adv.AdvUuid)
			if err != nil {
				return err
			}

			pathAdvNew := pathAdv

			if adv.LastAdvBack {
				pathAdvNew. =
			}

			ps.pathAdvService.Update()
		}
	}

}
