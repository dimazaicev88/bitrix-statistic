package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
)

type OptionModel struct {
	storage *storage.MysqlStorage
}

func NewOptionModel(storage *storage.MysqlStorage) OptionModel {
	return OptionModel{storage: storage}
}

func (o OptionModel) Add(options []entity.Option) error {
	return nil
}

func (o OptionModel) GetOption(s string) string {
	return ""
}
