package services

import (
	"bitrix-statistic/internal/models"
	"context"
)

type StatDayService struct {
	ctx       context.Context
	allModels *models.Models
}

func NewStatDay(ctx context.Context, allModels *models.Models) *StatDayService {
	return &StatDayService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (sds StatDayService) Add() {
	sds.allModels.Day.Add()
}
