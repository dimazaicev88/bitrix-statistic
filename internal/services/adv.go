package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"net/url"
	"regexp"
)

//TODO добавить авто создание рекламной компании

type AdvServices struct {
	allModels     *models.Models
	ctx           context.Context
	optionService *OptionService
}

func NewAdv(ctx context.Context, allModels *models.Models) *AdvServices {
	return &AdvServices{
		ctx:       ctx,
		allModels: allModels,
	}
}

// GetAdv Получить рекламную компанию
func (as AdvServices) GetAdv(fullUrl string) (entitydb.AdvReferer, error) {
	var totalListUuidAdv []string

	urlValues, err := url.Parse(fullUrl)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	urlWithoutScheme, err := url.JoinPath(urlValues.Host, urlValues.Path)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	advUuidsPageTo, err := as.allModels.AdvModel.FindAdvUuidByByPage(urlWithoutScheme, "TO") //Поиск рекламных компаний по условию Куда пришли [%_]:	(полные адреса страниц вашего сайта	разделенные новой строкой)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, advUuidsPageTo...)

	advUuidsSearcher, err := as.allModels.AdvModel.FindByByDomainSearcher(utils.StringConcat(urlValues.Scheme, urlValues.Host))
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, advUuidsSearcher...)

	advUuidsPageFrom, err := as.allModels.AdvModel.FindAdvUuidByByPage("FROM", fullUrl) //Откуда пришли [%_]: (ссылающиеся страницы,	разделенные новой строкой)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}
	totalListUuidAdv = append(totalListUuidAdv, advUuidsPageFrom...)

	byReferer, err := as.allModels.AdvModel.FindByReferer(urlValues.Query().Get("referer1"), urlValues.Query().Get("referer2")) //Поиск по referrer
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, byReferer...)

	referer, err := as.allModels.AdvModel.FindRefererByListAdv(totalListUuidAdv)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}

	//	if am.optionModel.Find("ADV_NA") == "Y" {
	//		Na1 := am.optionModel.Find("AVD_NA_REFERER1")
	//		Na2 := am.optionModel.Find("AVD_NA_REFERER2")
	//		if (Na1 != "" || Na2 != "") && referer1 == Na1 && referer2 == Na2 {
	//			na = "Y"
	//		}
	//
	//	}
	//

	return referer, nil
}

func (as AdvServices) FindByUuid(advUuid string) (entitydb.Adv, error) {
	return as.allModels.AdvModel.FindByUuid(advUuid)
}

// DeleteByUuid Удаление рекламной компании по uuid
func (as AdvServices) DeleteByUuid(uuid string) error {
	return as.allModels.AdvModel.DeleteByUuid(uuid)
}

// AutoCreateAdv Автоматическое создание рекламной компании
func (as AdvServices) AutoCreateAdv(referer1, referer2, siteId string) error {

	referrers, err := as.allModels.AdvModel.FindByReferer(referer1, referer2)
	if err != nil {
		return err
	}

	if len(referrers) == 0 {
		if as.optionService.IsAdvAutoCreate(siteId) {
			var refererValid bool
			if as.optionService.IsRefererCheck(siteId) {
				refererValid, err = regexp.MatchString("/^([0-9A-Za-z_:;.,-])*$/", referer1)
				if err != nil {
					return err
				}
				if refererValid {
					refererValid, err = regexp.MatchString("/^([0-9A-Za-z_:;.,-])*$/", referer2)
				}
				if err != nil {
					return err
				}
			} else {
				refererValid = true
			}

			if refererValid {
				err := as.allModels.AdvModel.AddAdv(referer1, referer2)
				if err != nil {
					return nil
				}
			}
		}
	}
	return nil
}
