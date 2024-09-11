package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/google/uuid"
)

type PathCacheService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewPathCacheService(ctx context.Context, allModels *models.Models) *PathCacheService {
	return &PathCacheService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (pcs *PathCacheService) FindLastBySessionUuid(uuid uuid.UUID) (entitydb.PathCache, error) {
	return pcs.allModels.PathCache.FindLastBySessionUuid(uuid)
}

func (pcs *PathCacheService) FindByReferer(uuid uuid.UUID, referer string) (entitydb.PathCache, error) {
	return pcs.allModels.PathCache.FindByReferer(uuid, referer)
}

func (pcs *PathCacheService) FindBySession(uuid uuid.UUID) (entitydb.PathCache, error) {
	return pcs.allModels.PathCache.FindBySession(uuid)
}
