package services

import (
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
