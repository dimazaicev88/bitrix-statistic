package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
)

type OptionModel struct {
	storage storage.Storage
}

func NewOptionModel(storage storage.Storage) OptionModel {
	return OptionModel{storage: storage}
}

func (o OptionModel) Add(options []entity.Option) error {
	return nil
}

func (o OptionModel) Get(name string) string {
	return ""
}
