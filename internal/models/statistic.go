package models

import "bitrix-statistic/internal/storage"

type StatisticModel struct {
	storage *storage.MysqlStorage
}

func NewStatisticModel(storage *storage.MysqlStorage) StatisticModel {
	return StatisticModel{
		storage: storage,
	}
}
