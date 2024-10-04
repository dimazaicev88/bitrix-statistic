package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"net/url"
	"regexp"
)

//TODO добавить авто создание рекламной компании

type AdvServices struct {
	allModels     *models.Models
	ctx           context.Context
	optionService *OptionService
	hitService    *HitService
}

func NewAdv(ctx context.Context, allModels *models.Models) *AdvServices {
	return &AdvServices{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (as *AdvServices) SetHitService(hitService *HitService) {
	as.hitService = hitService
}

func (as *AdvServices) SetOptionService(optionService *OptionService) {
	as.optionService = optionService
}

// GetAdv Получить рекламную компанию
func (as *AdvServices) GetAdv(statData entityjson.UserData) (entitydb.AdvCompany, error) {
	var totalListUuidAdv []string

	previewHit, err := as.hitService.FindLastHitWithoutSession(statData.GuestUuid, statData.PHPSessionId)
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	adv, err := as.FindByUuid(previewHit.AdvUuid)
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	if adv != (entitydb.Adv{}) {
		return entitydb.AdvCompany{
			AdvUuid:  adv.Uuid,
			Referer1: adv.Referer1,
			Referer2: adv.Referer2,
			//Referer3:    adv.Referer3,
			LastAdvBack: true,
		}, nil
	}

	urlValues, err := url.Parse(statData.Url)
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	urlWithoutScheme, err := url.JoinPath(urlValues.Host, urlValues.Path)
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	advUuidsPageTo, err := as.allModels.AdvModel.FindAdvUuidByByPage(urlWithoutScheme, "TO") //Поиск рекламных компаний по условию Куда пришли [%_]:	(полные адреса страниц вашего сайта	разделенные новой строкой)
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, advUuidsPageTo...)

	advUuidsSearcher, err := as.allModels.AdvModel.FindByByDomainSearcher(utils.StringConcat(urlValues.Scheme, urlValues.Host))
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, advUuidsSearcher...)

	advUuidsPageFrom, err := as.allModels.AdvModel.FindAdvUuidByByPage("FROM", statData.Url) //Откуда пришли [%_]: (ссылающиеся страницы,	разделенные новой строкой)
	if err != nil {
		return entitydb.AdvCompany{}, err
	}
	totalListUuidAdv = append(totalListUuidAdv, advUuidsPageFrom...)

	byReferer, err := as.allModels.AdvModel.FindByReferer(urlValues.Query().Get("referer1"), urlValues.Query().Get("referer2")) //Поиск по referrer
	if err != nil {
		return entitydb.AdvCompany{}, err
	}

	totalListUuidAdv = append(totalListUuidAdv, byReferer...)

	referer, err := as.allModels.AdvModel.FindRefererByListAdv(totalListUuidAdv)
	if err != nil {
		return entitydb.AdvCompany{}, err
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

func (as *AdvServices) FindByUuid(advUuid uuid.UUID) (entitydb.Adv, error) {
	return as.allModels.AdvModel.FindByUuid(advUuid)
}

// DeleteByUuid Удаление рекламной компании по uuid
func (as *AdvServices) DeleteByUuid(advUuid uuid.UUID) error {
	return as.allModels.AdvModel.DeleteByUuid(advUuid)
}

// AutoCreateAdv Автоматическое создание рекламной компании
func (as *AdvServices) AutoCreateAdv(referer1, referer2 string) error {

	referrers, err := as.allModels.AdvModel.FindByReferer(referer1, referer2)
	if err != nil {
		return err
	}

	if len(referrers) == 0 {
		if as.optionService.AdvAutoCreate() {
			var refererValid bool
			if as.optionService.RefererCheck() {
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

func (as *AdvServices) IsExistsAdv(advUuid uuid.UUID) (bool, error) {
	return as.allModels.AdvModel.IsExistsAdv(advUuid)
}
