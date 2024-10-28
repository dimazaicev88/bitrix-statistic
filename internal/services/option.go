package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type OptionService struct {
	allModels   *models.Models
	ctx         context.Context
	optionCache otter.Cache[string, string]
}

func NewOption(ctx context.Context, allModels *models.Models) *OptionService {
	otterCache, err := otter.MustBuilder[string, string](100).
		CollectStats().
		WithTTL(time.Hour * 730).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}
	return &OptionService{
		ctx:         ctx,
		allModels:   allModels,
		optionCache: otterCache,
	}
}

func (o *OptionService) Add(options entitydb.Option) error {
	return o.allModels.Option.Add(options)
}

func (o *OptionService) Set(option entitydb.Option) error {
	if err := o.allModels.Option.Set(option); err != nil {
		return err
	}
	o.optionCache.Set(option.Name, option.Value)
	return nil
}

func (o *OptionService) IsSaveVisits() bool {
	val, _ := o.get("saveVisits", "true")
	return val == "true"
}

func (o *OptionService) IsSaveReferrers() bool {
	val, _ := o.get("saveReferrers", "true")
	return val == "true"
}

func (o *OptionService) IsSaveHits() bool {
	val, _ := o.get("saveHits", "true")
	return val == "true"
}

func (o *OptionService) IsSaveAdditional() bool {
	val, _ := o.get("saveAdditional", "true")
	return val == "true"

}

func (o *OptionService) IsSavePathData() bool {
	val, _ := o.get("savePathData", "true")
	return val == "true"
}

func (o *OptionService) IsAdvNa() bool {
	val, _ := o.get("advNa", "true")
	return val == "true"
}

func (o *OptionService) AvdNaReferer1() string {
	val, _ := o.get("avdNaReferer1", "NA")
	return val
}

func (o *OptionService) AvdNaReferer2() string {
	val, _ := o.get("avdNaReferer2", "NA")
	return val
}

func (o *OptionService) Referer1Syn() string {
	val, _ := o.get("referer1Syn", "r1")
	return val
}

func (o *OptionService) Referer2Syn() string {
	val, _ := o.get("referer2Syn", "r2")
	return val
}

func (o *OptionService) Referer3Syn() string {
	val, _ := o.get("referer3Syn", "r3")
	return val
}

func (o *OptionService) IsRefererCheck() bool {
	val, _ := o.get("refererCheck", "false")
	return val == "true"

}

func (o *OptionService) IsAdvAutoCreate() bool {
	val, _ := o.get("advAutoCreate", "true")
	return val == "true"
}

func (o *OptionService) AdvGuestDays() uint64 {
	val, _ := o.get("advGuestDays", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) AdvDays() uint64 {
	val, _ := o.get("advDays", "365")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) SearcherHitDays() uint64 {
	val, _ := o.get("searcherHitDays", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) SearcherDays() uint64 {
	val, _ := o.get("searcherDays", "360")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) EventsDays() uint64 {
	val, _ := o.get("eventsDays", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) EventDynamicDays() uint64 {
	val, _ := o.get("eventDynamicDays", "360")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) VisitDays() uint64 {
	val, _ := o.get("visitDays", "10")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) CityDays() uint64 {
	val, _ := o.get("cityDays", "360")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) CountryDays() uint64 {
	val, _ := o.get("countryDays", "360")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) PathDays() uint64 {
	val, _ := o.get("pathDays", "10")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) GuestDays() uint64 {
	val, _ := o.get("guestDays", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) SessionDays() uint64 {
	val, _ := o.get("sessionDays", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) HitDays() uint64 {
	val, _ := o.get("hitDays", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) PhrasesDays() uint64 {
	val, _ := o.get("phrasesDays", "10")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) RefererListDays() uint64 {
	val, _ := o.get("refererListDays", "10")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) RefererDays() uint64 {
	val, _ := o.get("refererDays", "360")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) RefererTop() uint64 {
	val, _ := o.get("refererTop", "500")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) MaxPathSteps() uint32 {
	val, _ := o.get("maxPathSteps", "3")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return uint32(parseUint)
}

func (o *OptionService) OnlineInterval() uint64 {
	val, _ := o.get("onlineInterval", "180")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) RecordsLimit() uint64 {
	val, _ := o.get("recordsLimit", "500")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) EventGidBase64Encode() bool {
	val, _ := o.get("eventGidBase64Encode", "true")
	return val == "true"
}

func (o *OptionService) AdvEventsDefault() string {
	val, _ := o.get("advEventsDefault", "list")
	return val
}

func (o *OptionService) UseAutoOptimize() bool {
	val, _ := o.get("useAutoOptimize", "false")
	return val == "true"
}

func (o *OptionService) BaseCurrency() string {
	val, _ := o.get("baseCurrency", "xxx")
	return val
}

func (o *OptionService) DefenceOn() bool {
	val, _ := o.get("defenceOn", "true")
	return val == "true"
}

func (o *OptionService) DefenceStackTime() uint64 {
	val, _ := o.get("defenceStackTime", "10")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) DefenceMaxStackHits() uint64 {
	val, _ := o.get("defenceMaxStackHits", "15")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) DefenceDelay() uint64 {
	val, _ := o.get("defenceDelay", "300")
	parseUint, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0
	}
	return parseUint
}

func (o *OptionService) DefenceLog() bool {
	val, _ := o.get("defenceLog", "false")
	return val == "true"
}

func (o *OptionService) ImportantPageParams() string {
	val, _ := o.get("importantPageParams", "ID, IBLOCK_ID, SECTION_ID, PARENT_ELEMENT_ID, FID, TID, MID, UID, VOTE_ID, print, goto")
	return val
}

func (o *OptionService) SkipStatisticWhat() string {
	val, _ := o.get("skipStatisticWhat", "none")
	return val
}

func (o *OptionService) SkipStatisticGroups() string {
	val, _ := o.get("skipStatisticGroups", "")
	return val
}

func (o *OptionService) SkipStatisticIpRanges() string {
	val, _ := o.get("skipStatisticIpRanges", "")
	return val
}

func (o *OptionService) DirectoryIndex() string {
	val, _ := o.get("directoryIndex", "")
	return val
}

// SearcherEvents Учитывать события рекламных кампаний для поисковиков
func (o *OptionService) SearcherEvents() bool {
	val, _ := o.get("searcherEvents", "true")
	return val == "true"
}

func (o *OptionService) get(name, defValue string) (string, error) {
	val, ok := o.optionCache.Get(name)
	if !ok {
		dbVal, err := o.allModels.Option.Find(name)
		if err != nil {
			return "", err
		}
		if dbVal == (entitydb.Option{}) {
			return defValue, nil
		}
		val = dbVal.Value
		o.optionCache.Set(name, dbVal.Value)
	}
	return val, nil
}

func (o *OptionService) GetOptions() dto.Options {
	return dto.Options{
		AdvCompany: &dto.AdvCompany{
			AdvNa:          o.IsAdvNa(),
			AdvAutoCreate:  o.IsAdvAutoCreate(),
			RefererCheck:   o.IsRefererCheck(),
			SearcherEvents: o.SearcherEvents(),
			Referer1Syn:    o.Referer1Syn(),
			Referer2Syn:    o.Referer2Syn(),
			Referer3Syn:    o.Referer3Syn(),
		},
	}
}

func (o *OptionService) SetAdvCompany(company *dto.AdvCompany) error {
	if err := o.SetAdvNa(strconv.FormatBool(company.AdvNa)); err != nil {
		return err
	}

	if err := o.SetAdvAutoCreate(strconv.FormatBool(company.AdvAutoCreate)); err != nil {
		return err
	}

	if err := o.SetRefererCheck(strconv.FormatBool(company.RefererCheck)); err != nil {
		return err
	}

	if err := o.SetSearcherEvents(strconv.FormatBool(company.SearcherEvents)); err != nil {
		return err
	}

	if err := o.SetReferer1Syn(company.Referer1Syn); err != nil {
		return err
	}

	if err := o.SetReferer2Syn(company.Referer2Syn); err != nil {
		return err
	}

	if err := o.SetReferer3Syn(company.Referer3Syn); err != nil {
		return err
	}

	return nil
}

func (o *OptionService) SetAdvNa(value string) error {
	if err := o.Set(entitydb.Option{Name: "advNa", Value: value}); err != nil {
		return err
	}
	return nil
}

func (o *OptionService) SetAdvAutoCreate(value string) error {
	if err := o.Set(entitydb.Option{Name: "advAutoCreate", Value: value}); err != nil {
		return err
	}
	return nil
}

func (o *OptionService) SetRefererCheck(value string) error {
	if err := o.Set(entitydb.Option{Name: "refererCheck", Value: value}); err != nil {
		return err
	}
	return nil
}

func (o *OptionService) SetSearcherEvents(value string) error {
	if err := o.Set(entitydb.Option{Name: "searcherEvents", Value: value}); err != nil {
		return err
	}
	return nil
}

func (o *OptionService) SetReferer1Syn(value string) error {
	if err := o.Set(entitydb.Option{Name: "referer1Syn", Value: value}); err != nil {
		return err
	}
	return nil
}

func (o *OptionService) SetReferer2Syn(value string) error {
	if err := o.Set(entitydb.Option{Name: "referer2Syn", Value: value}); err != nil {
		return err
	}
	return nil
}

func (o *OptionService) SetReferer3Syn(value string) error {
	if err := o.Set(entitydb.Option{Name: "referer3Syn", Value: value}); err != nil {
		return err
	}
	return nil
}
