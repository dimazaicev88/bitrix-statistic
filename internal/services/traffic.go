package services

import (
	"bitrix-statistic/internal/models"
	"context"
)

type TrafficService struct {
	ctx       context.Context
	allModels *models.Models
}

func NewTraffic(ctx context.Context, allModels *models.Models) *TrafficService {
	return &TrafficService{
		ctx:       ctx,
		allModels: allModels,
	}
}
