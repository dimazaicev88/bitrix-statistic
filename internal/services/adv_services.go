package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"net/url"
)

//TODO добавить авто создание рекламной компании

type AdvServices struct {
	advModel models.AdvModel
}

func NewAdvServices(advModel models.AdvModel) *AdvServices {
	return &AdvServices{advModel: advModel}
}

// AutoCreateAdv Автоматическое создание рекламной компании
func (as AdvServices) AutoCreateAdv(fullUrl string) error {
	return nil
}

func (as AdvServices) ParseAdv(fullUrl string) (entity.AdvReferer, error) {
	parse, err := url.Parse(fullUrl)
	if err != nil {
		return entity.AdvReferer{}, err
	}
	urlQuery := parse.Query()

	return entity.AdvReferer{
		Referer1: urlQuery.Get("referer1"),
		Referer2: urlQuery.Get("referer2"),
		Referer3: urlQuery.Get("referer3"),
	}, err
}
