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
	return pas.allModels.PathAdv.FindByPathUuid(pathId, dateStat)
}

//func (pas PathAdvService) Update(oldValue entitydb.PathAdv, newValue entitydb.PathAdv) error {
//	return pas.allModels.PathAdv.Update(oldValue, newValue)
//}

func (pas PathAdvService) FindByPageAndAdvUuid(pathId int32, advUuid string) (entitydb.PathAdv, error) {
	return pas.allModels.PathAdv.FindByPageAndAdvUuid(pathId, advUuid)
}