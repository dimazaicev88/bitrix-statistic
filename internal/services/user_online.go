package services

import (
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
