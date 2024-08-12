package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
)

type PathAdvService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewPathAdvService(ctx context.Context, allModels *models.Models) *PathAdvService {
	return &PathAdvService{
		allModels: allModels,
		ctx:       ctx,
	}
}

func (pas PathAdvService) Add(pathAdv entitydb.PathAdv) error {
	return pas.allModels.PathAdv.Add(pathAdv)
}

func (pas PathAdvService) FindByPathId(pathId int32, dateStat string) (entitydb.PathAdv, error) {
	return pas.allModels.PathAdv.FindByPathId(pathId, dateStat)
}

func (pas PathAdvService) Update(oldValue entitydb.PathAdv, newValue entitydb.PathAdv) {
	pas.allModels.PathAdv.Update(oldValue, newValue)
}
