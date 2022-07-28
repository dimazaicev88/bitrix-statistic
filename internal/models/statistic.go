package models

import "bitrix-statistic/internal/storage"

type StatisticModel struct {
	storage storage.Storage
}

func NewStatisticModel(storage storage.Storage) StatisticModel {
	return StatisticModel{
		storage: storage,
	}
}
