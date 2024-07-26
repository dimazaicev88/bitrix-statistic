package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"net/url"
)

//TODO добавить авто создание рекламной компании

type AdvServices struct {
	AdvModel *models.Adv
}

func NewAdv(ctx context.Context, chClient driver.Conn) *AdvServices {
	return &AdvServices{AdvModel: models.NewAdv(ctx, chClient, models.NewOption(ctx, chClient))}
}

// AutoCreateAdv Автоматическое создание рекламной компании
func (as AdvServices) AutoCreateAdv(fullUrl string) error {
	return nil
}

// GetAdv Получить рекламную компанию
func (as AdvServices) GetAdv(fullUrl string) (entitydb.AdvReferer, error) {
	var totalListUuidAdv []string
	urlValues, err := url.Parse(fullUrl)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}
	urlWithoutCheme, err := url.JoinPath(urlValues.Host, urlValues.Path)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	advUuidsPageTo, err := as.AdvModel.FindAdvUuidByByPage(urlWithoutCheme, "TO") //Поиск рекламных компаний по условию Куда пришли [%_]:	(полные адреса страниц вашего сайта	разделенные новой строкой)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, advUuidsPageTo...)

	advUuidsSearcher, err := as.AdvModel.FindByByDomainSearcher(utils.StringConcat(urlValues.Scheme, urlValues.Host))
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, advUuidsSearcher...)

	advUuidsPageFrom, err := as.AdvModel.FindAdvUuidByByPage("FROM", fullUrl) //Откуда пришли [%_]: (ссылающиеся страницы,	разделенные новой строкой)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}
	totalListUuidAdv = append(totalListUuidAdv, advUuidsPageFrom...)

	byReferer, err := as.AdvModel.FindByReferer(urlValues.Query().Get("referer1"), urlValues.Query().Get("referer2")) //Поиск по referrer
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, byReferer...)

	referer, err := as.AdvModel.FindRefererByListAdv(totalListUuidAdv)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	//	if am.optionModel.Get("ADV_NA") == "Y" {
	//		Na1 := am.optionModel.Get("AVD_NA_REFERER1")
	//		Na2 := am.optionModel.Get("AVD_NA_REFERER2")
	//		if (Na1 != "" || Na2 != "") && referer1 == Na1 && referer2 == Na2 {
	//			na = "Y"
	//		}
	//
	//	}
	//

	return referer, nil
}

func (as AdvServices) FindByUuid(advUuid string) (entitydb.Adv, error) {
	return as.AdvModel.FindByUuid(advUuid)
}

// DeleteByUuid Удаление рекламной компании по uuid
func (as AdvServices) DeleteByUuid(uuid string) error {
	return as.AdvModel.DeleteByUuid(uuid)
}
