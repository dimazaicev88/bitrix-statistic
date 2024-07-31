package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"errors"
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	"time"
)

type OptionService struct {
	allModels   *models.Models
	ctx         context.Context
	optionCache otter.Cache[string, interface{}]
}

func NewOption(ctx context.Context, allModels *models.Models) *OptionService {
	otterCache, err := otter.MustBuilder[string, interface{}](100).
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

func (o OptionService) Add(options entitydb.Option) error {
	return o.allModels.Option.Add(options)
}

func (o OptionService) Set(option entitydb.Option) error {
	switch option.Value.(type) {

	case int:
	case float64:
	case string:
	default:
		return errors.New("option value type is invalid")
	}

	if err := o.allModels.Option.Set(option); err != nil {
		return err
	}
	o.optionCache.Set(utils.StringConcat(option.Name, ":", option.SiteId), option.Value)
	return nil
}

func (o OptionService) IsSaveVisits(siteId string) bool {
	val, _ := o.get("save_visits", siteId, true)
	return val.(bool)
}

func (o OptionService) IsSaveReferrers(siteId string) bool {
	val, _ := o.get("save_referrers", siteId, true)
	return val.(bool)
}

func (o OptionService) IsSaveHits(siteId string) bool {
	val, _ := o.get("save_hits", siteId, true)
	return val.(bool)
}

func (o OptionService) IsSaveAdditional(siteId string) bool {
	val, _ := o.get("save_additional", siteId, true)
	return val.(bool)

}

func (o OptionService) IsSavePathData(siteId string) bool {
	val, _ := o.get("save_path_data", siteId, true)
	return val.(bool)
}

func (o OptionService) IsAdvNa(siteId string) bool {
	val, _ := o.get("adv_na", siteId, true)
	return val.(bool)
}

func (o OptionService) AvdNaReferer1(siteId string) string {
	val, _ := o.get("avd_na_referer1", siteId, "NA")
	return val.(string)
}

func (o OptionService) AvdNaReferer2(siteId string) string {
	val, _ := o.get("avd_na_referer2", siteId, "NA")
	return val.(string)
}

func (o OptionService) Referer1Syn(siteId string) string {
	val, _ := o.get("referer1_syn", siteId, "r1")
	return val.(string)
}

func (o OptionService) Referer2Syn(siteId string) string {
	val, _ := o.get("referer2_syn", siteId, "r2")
	return val.(string)
}

func (o OptionService) Referer3Syn(siteId string) string {
	val, _ := o.get("referer3_syn", siteId, "r3")
	return val.(string)
}

func (o OptionService) IsRefererCheck(siteId string) bool {
	val, _ := o.get("referer_check", siteId, false)
	return val.(bool)

}

func (o OptionService) OpenStatActive(siteId string) bool {
	val, _ := o.get("open_stat_active", siteId, false)
	return val.(bool)
}

func (o OptionService) OpenStatR1Template(siteId string) string {
	val, _ := o.get("open_stat_r1_template", siteId, "#service-name#_#campaign-id#")
	return val.(string)
}

func (o OptionService) OpenStatR2Template(siteId string) string {

	val, _ := o.get("open_stat_r2_template", siteId, "#ad-id#_#source-id#")
	return val.(string)
}

func (o OptionService) IsAdvAutoCreate(siteId string) bool {
	val, _ := o.get("adv_auto_create", siteId, true)
	return val.(bool)
}

func (o OptionService) AdvGuestDays(siteId string) uint32 {
	val, _ := o.get("adv_guestDays", siteId, 3)
	return val.(uint32)
}

func (o OptionService) AdvDays(siteId string) uint32 {
	val, _ := o.get("adv_days", siteId, 365)
	return val.(uint32)
}

func (o OptionService) SearcherHitDays(siteId string) uint32 {
	val, _ := o.get("searcher_hit_days", siteId, 3)
	return val.(uint32)
}

func (o OptionService) SearcherDays(siteId string) uint32 {
	val, _ := o.get("searcher_days", siteId, 360)
	return val.(uint32)
}

func (o OptionService) EventsDays(siteId string) uint32 {
	val, _ := o.get("events_days", siteId, 3)
	return val.(uint32)
}

func (o OptionService) EventDynamicDays(siteId string) uint32 {
	val, _ := o.get("event_dynamic_days", siteId, 360)
	return val.(uint32)
}

func (o OptionService) VisitDays(siteId string) uint32 {
	val, _ := o.get("visit_days", siteId, 10)
	return val.(uint32)
}

func (o OptionService) CityDays(siteId string) uint32 {
	val, _ := o.get("city_days", siteId, 360)
	return val.(uint32)
}

func (o OptionService) CountryDays(siteId string) uint32 {
	val, _ := o.get("country_days", siteId, 360)
	return val.(uint32)
}

func (o OptionService) PathDays(siteId string) uint32 {
	val, _ := o.get("path_days", siteId, 10)
	return val.(uint32)
}

func (o OptionService) GuestDays(siteId string) uint32 {
	val, _ := o.get("guest_days", siteId, 3)
	return val.(uint32)
}

func (o OptionService) SessionDays(siteId string) uint32 {
	val, _ := o.get("session_days", siteId, 3)
	return val.(uint32)
}

func (o OptionService) HitDays(siteId string) uint32 {
	val, _ := o.get("hit_days", siteId, 3)
	return val.(uint32)
}

func (o OptionService) PhrasesDays(siteId string) uint32 {
	val, _ := o.get("phrases_days", siteId, 10)
	return val.(uint32)
}

func (o OptionService) RefererListDays(siteId string) uint32 {
	val, _ := o.get("referer_list_days", siteId, 10)
	return val.(uint32)
}

func (o OptionService) RefererDays(siteId string) uint32 {
	val, _ := o.get("referer_days", siteId, 360)
	return val.(uint32)
}

func (o OptionService) RefererTop(siteId string) uint32 {
	val, _ := o.get("referer_top", siteId, 500)
	return val.(uint32)
}

func (o OptionService) MaxPathSteps(siteId string) uint32 {
	val, _ := o.get("max_path_steps", siteId, 3)
	return val.(uint32)
}

func (o OptionService) OnlineInterval(siteId string) uint32 {
	val, _ := o.get("online_interval", siteId, 180)
	return val.(uint32)
}

func (o OptionService) RecordsLimit(siteId string) uint32 {
	val, _ := o.get("records_limit", siteId, 500)
	return val.(uint32)
}

func (o OptionService) EventGidBase64Encode(siteId string) bool {
	val, _ := o.get("event_gid_base64_encode", siteId, true)
	return val.(bool)
}

func (o OptionService) AdvEventsDefault(siteId string) string {
	val, _ := o.get("adv_events_default", siteId, "list")
	return val.(string)
}

func (o OptionService) UseAutoOptimize(siteId string) bool {
	val, _ := o.get("use_auto_optimize", siteId, false)
	return val.(bool)
}

func (o OptionService) BaseCurrency(siteId string) string {
	val, _ := o.get("base_currency", siteId, "xxx")
	return val.(string)
}

func (o OptionService) IsDefenceOn(siteId string) bool {
	val, _ := o.get("defence_on", siteId, true)
	return val.(bool)
}

func (o OptionService) DefenceStackTime(siteId string) uint32 {
	val, _ := o.get("defence_stack_time", siteId, 10)
	return val.(uint32)
}

func (o OptionService) DefenceMaxStackHits(siteId string) uint32 {
	val, _ := o.get("defence_max_stack_hits", siteId, 15)
	return val.(uint32)
}

func (o OptionService) DefenceDelay(siteId string) uint32 {
	val, _ := o.get("defence_delay", siteId, 300)
	return val.(uint32)
}

func (o OptionService) IsDefenceLog(siteId string) bool {
	val, _ := o.get("defence_log", siteId, false)
	return val.(bool)
}

func (o OptionService) ImportantPageParams(siteId string) string {
	val, _ := o.get("important_page_params", siteId, "ID, IBLOCK_ID, SECTION_ID, PARENT_ELEMENT_ID, FID, TID, MID, UID, VOTE_ID, print, goto")
	return val.(string)
}

func (o OptionService) SkipStatisticWhat(siteId string) string {
	val, _ := o.get("skip_statistic_what", siteId, "none")
	return val.(string)
}

func (o OptionService) SkipStatisticGroups(siteId string) string {
	val, _ := o.get("skip_statistic_groups", siteId, "")
	return val.(string)
}

func (o OptionService) SkipStatisticIpRanges(siteId string) string {
	val, _ := o.get("skip_statistic_ip_ranges", siteId, "")
	return val.(string)
}

func (o OptionService) DirectoryIndex(siteId string) string {
	val, _ := o.get("directory_index", siteId, "")
	return val.(string)
}

// IsSearcherEvents Учитывать события рекламных кампаний для поисковиков
func (o OptionService) IsSearcherEvents(siteId string) bool {
	val, _ := o.get("searcher_events", siteId, true)
	return val.(bool)
}

func (o OptionService) get(name, site string, defValue interface{}) (interface{}, error) {
	val, ok := o.optionCache.Get(utils.StringConcat(name, ":", site))
	if !ok {
		dbVal, err := o.allModels.Option.Find(name, site)
		if err != nil {
			return nil, err
		}
		if dbVal == (entitydb.Option{}) {
			return defValue, nil
		}
		o.optionCache.Set(utils.StringConcat(name, ":", site), dbVal.Value)
	}
	return val, nil
}
