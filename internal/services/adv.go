package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
	"net/url"
	"regexp"
)

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
func (as *AdvServices) GetAdv(statData dto.UserData) (entitydb.AdvCompany, error) {
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
			AdvUuid:     adv.Uuid,
			Referer1:    adv.Referer1,
			Referer2:    adv.Referer2,
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

	return referer, nil
}

func (as *AdvServices) FindByUuid(advUuid uuid.UUID) (entitydb.Adv, error) {
	return as.allModels.AdvModel.FindByUuid(advUuid)
}

// Delete Удаление рекламной компании по uuid
func (as *AdvServices) Delete(advUuid uuid.UUID) error {
	return as.allModels.AdvModel.Delete(advUuid)
}

// AutoCreateAdv Автоматическое создание рекламной компании
func (as *AdvServices) AutoCreateAdv(referer1, referer2 string) (entitydb.Adv, error) {
	referrers, err := as.allModels.AdvModel.FindByReferer(referer1, referer2)
	if err != nil {
		return entitydb.Adv{}, err
	}

	var advDb entitydb.Adv
	if len(referrers) == 0 {
		if as.optionService.IsAdvAutoCreate() {
			if as.optionService.IsRefererCheck() {
				pattern := `^([0-9A-Za-z_:;.,-])*`
				re := regexp.MustCompile(pattern)
				if re.MatchString(referer1) && re.MatchString(referer2) {
					advDb, err = as.allModels.AdvModel.AddAdv(referer1, referer2)
					if err != nil {
						return entitydb.Adv{}, err
					}
				}
			} else if referer1 != "" && referer2 != "" {
				advDb, err = as.allModels.AdvModel.AddAdv(referer1, referer2)
				if err != nil {
					return entitydb.Adv{}, err
				}
			}
		}

		if as.optionService.IsAdvNa() {
			advDb, err = as.allModels.AdvModel.AddAdv("NA", "NA")
			if err != nil {
				return entitydb.Adv{}, err
			}
		}
	}
	return advDb, nil
}

func (as *AdvServices) IsExistsAdv(advUuid uuid.UUID) (bool, error) {
	return as.allModels.AdvModel.IsExistsAdv(advUuid)
}

func (as *AdvServices) AddAdvStat(advStat entitydb.AdvStat) error {
	return as.allModels.AdvModel.AddAdvStat(advStat)
}

func (as *AdvServices) AddAdvDay(day entitydb.AdvDay) error {
	return as.allModels.AdvModel.AddAdvDay(day)
}

func (as *AdvServices) GetDynamicList(filter filters.Filter, getMaxMin bool) (entitydb.AdvDynamicResult, error) {
	return as.allModels.AdvModel.GetDynamicList(filter, getMaxMin)
}

func (as *AdvServices) GetEventList(filter filters.Filter) ([]entitydb.Event, error) {
	return as.allModels.AdvModel.GetEventList(filter)
}

func (as *AdvServices) Find(filter filters.Filter) ([]entitydb.Adv, error) {
	return as.allModels.AdvModel.Find(filter)
}

func (as *AdvServices) FindAll(u uint32, u2 uint32) ([]entitydb.Adv, error) {
	return nil, nil
}

func (as *AdvServices) ConvertToJSONListAdv(adv []entitydb.Adv) any {
	return nil
}
