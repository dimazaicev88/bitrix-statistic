package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
)

type UserOnlineService struct {
	ctx       context.Context
	allModels *models.Models
}

func NewUserOnline(ctx context.Context, allModels *models.Models) *UserOnlineService {
	return &UserOnlineService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (uos UserOnlineService) FindAll(skip uint32, limit uint32) ([]entitydb.UserOnline, error) {
	return uos.allModels.UserOnline.FindAll(skip, limit)
}
