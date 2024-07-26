package services

import (
	"bitrix-statistic/internal/cache"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type OptionService struct {
	optionModel *models.Option
}

func NewOption(ctx context.Context, chClient driver.Conn) *OptionService {
	return &OptionService{
		optionModel: models.NewOption(ctx, chClient),
	}
}

func (o OptionService) Add(options entitydb.Option) error {
	return o.optionModel.Add(options)
}

func (o OptionService) Set(option entitydb.Option) error {
	switch option.Value.(type) {

	case int:
	case float64:
	case string:
	default:
		return errors.New("option value type is invalid")
	}

	if err := o.optionModel.Set(option); err != nil {
		return err
	}
	cache.Cache().Set(utils.StringConcat(option.Name, ":", option.SiteId), option.Value)
	return nil
}

func (o OptionService) IsSaveVisits() bool {
	val, ok := cache.Cache().Get("save_visits")
	if !ok {

		return true
	}
	return val.(bool)
}

func (o OptionService) IsSaveReferrers() bool {
	val, ok := cache.Cache().Get("save_referrers")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) IsSaveHits() bool {
	val, ok := cache.Cache().Get("save_hits")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) IsSaveAdditional() bool {
	val, ok := cache.Cache().Get("save_additional")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) IsSavePathData() bool {
	val, ok := cache.Cache().Get("save_path_data")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) IsAdvNa() bool {
	val, ok := cache.Cache().Get("adv_na")
	if !ok {
		return false
	}
	return val.(bool)
}

func (o OptionService) AvdNaReferer1() string {
	val, ok := cache.Cache().Get("avd_na_referer1")
	if !ok {
		return "NA"
	}
	return val.(string)
}

func (o OptionService) AvdNaReferer2() string {
	val, ok := cache.Cache().Get("avd_na_referer2")
	if !ok {
		return "NA"
	}
	return val.(string)
}

func (o OptionService) Referer1Syn() string {
	val, ok := cache.Cache().Get("referer1_syn")
	if !ok {
		return "r1"
	}
	return val.(string)
}

func (o OptionService) Referer2Syn() string {
	val, ok := cache.Cache().Get("referer2_syn")
	if !ok {
		return "r2"
	}
	return val.(string)
}

func (o OptionService) Referer3Syn() string {
	val, ok := cache.Cache().Get("referer3_syn")
	if !ok {
		return "r3"
	}
	return val.(string)
}

func (o OptionService) IsRefererCheck() bool {
	val, ok := cache.Cache().Get("referer_check")
	if !ok {
		return false
	}
	return val.(bool)
}

func (o OptionService) OpenStatActive() bool {
	val, ok := cache.Cache().Get("open_stat_active")
	if !ok {
		return false
	}
	return val.(bool)
}

func (o OptionService) OpenStatR1Template() string {
	val, ok := cache.Cache().Get("open_stat_r1_template")
	if !ok {
		return "#service-name#_#campaign-id#"
	}
	return val.(string)
}

func (o OptionService) OpenStatR2Template() string {
	val, ok := cache.Cache().Get("open_stat_r2_template")
	if !ok {
		return "#ad-id#_#source-id#"
	}
	return val.(string)
}

func (o OptionService) IsAdvAutoCreate() bool {
	val, ok := cache.Cache().Get("adv_auto_create")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) AdvGuestDays() uint32 {
	val, ok := cache.Cache().Get("adv_guestDays")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) AdvDays() uint32 {
	val, ok := cache.Cache().Get("adv_days")
	if !ok {
		return 365
	}
	return val.(uint32)
}

func (o OptionService) SearcherHitDays() uint32 {
	val, ok := cache.Cache().Get("searcher_hit_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) SearcherDays() uint32 {
	val, ok := cache.Cache().Get("searcher_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func (o OptionService) EventsDays() uint32 {
	val, ok := cache.Cache().Get("events_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) EventDynamicDays() uint32 {
	val, ok := cache.Cache().Get("event_dynamic_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func (o OptionService) VisitDays() uint32 {
	val, ok := cache.Cache().Get("visit_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func (o OptionService) CityDays() uint32 {
	val, ok := cache.Cache().Get("city_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func (o OptionService) CountryDays() uint32 {
	val, ok := cache.Cache().Get("country_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func (o OptionService) PathDays() uint32 {
	val, ok := cache.Cache().Get("path_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func (o OptionService) GuestDays() uint32 {
	val, ok := cache.Cache().Get("guest_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) SessionDays() uint32 {
	val, ok := cache.Cache().Get("session_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) HitDays() uint32 {
	val, ok := cache.Cache().Get("hit_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) PhrasesDays() uint32 {
	val, ok := cache.Cache().Get("phrases_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func (o OptionService) RefererListDays() uint32 {
	val, ok := cache.Cache().Get("referer_list_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func (o OptionService) RefererDays() uint32 {
	val, ok := cache.Cache().Get("referer_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func (o OptionService) RefererTop() uint32 {
	val, ok := cache.Cache().Get("referer_top")
	if !ok {
		return 500
	}
	return val.(uint32)
}

func (o OptionService) MaxPathSteps() uint32 {
	val, ok := cache.Cache().Get("max_path_steps")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func (o OptionService) OnlineInterval() uint32 {
	val, ok := cache.Cache().Get("online_interval")
	if !ok {
		return 180
	}
	return val.(uint32)
}

func (o OptionService) RecordsLimit() uint32 {
	val, ok := cache.Cache().Get("records_limit")
	if !ok {
		return 500
	}
	return val.(uint32)
}

func (o OptionService) EventGidBase64Encode() bool {
	val, ok := cache.Cache().Get("event_gid_base64_encode")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) AdvEventsDefault() string {
	val, ok := cache.Cache().Get("adv_events_default")
	if !ok {
		return "list"
	}
	return val.(string)
}

func (o OptionService) UseAutoOptimize() bool {
	val, ok := cache.Cache().Get("use_auto_optimize")
	if !ok {
		return false
	}
	return val.(bool)
}

func (o OptionService) BaseCurrency() string {
	val, ok := cache.Cache().Get("base_currency")
	if !ok {
		return "xxx"
	}
	return val.(string)
}

func (o OptionService) IsDefenceOn() bool {
	val, ok := cache.Cache().Get("defence_on")
	if !ok {
		return true
	}
	return val.(bool)
}

func (o OptionService) DefenceStackTime() uint32 {
	val, ok := cache.Cache().Get("defence_stack_time")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func (o OptionService) DefenceMaxStackHits() uint32 {
	val, ok := cache.Cache().Get("defence_max_stack_hits")
	if !ok {
		return 15
	}
	return val.(uint32)
}

func (o OptionService) DefenceDelay() uint32 {
	val, ok := cache.Cache().Get("defence_delay")
	if !ok {
		return 300
	}
	return val.(uint32)
}

func (o OptionService) IsDefenceLog() bool {
	val, ok := cache.Cache().Get("defence_log")
	if !ok {
		return false
	}
	return val.(bool)
}

func (o OptionService) ImportantPageParams() string {
	val, ok := cache.Cache().Get("important_page_params")
	if !ok {
		return "ID, IBLOCK_ID, SECTION_ID, PARENT_ELEMENT_ID, FID, TID, MID, UID, VOTE_ID, print, goto"
	}
	return val.(string)
}

func (o OptionService) SkipStatisticWhat() string {
	val, ok := cache.Cache().Get("skip_statistic_what")
	if !ok {
		return "none"
	}
	return val.(string)
}

func (o OptionService) SkipStatisticGroups() string {
	val, ok := cache.Cache().Get("skip_statistic_groups")
	if !ok {
		return ""
	}
	return val.(string)
}

func (o OptionService) SkipStatisticIpRanges() string {
	val, ok := cache.Cache().Get("skip_statistic_ip_ranges")
	if !ok {
		return ""
	}
	return val.(string)
}

func (o OptionService) DirectoryIndex() string {
	val, ok := cache.Cache().Get("directory_index")
	if !ok {
		return ""
	}
	return val.(string)
}

// IsSearcherEvents Учитывать события рекламных кампаний для поисковиков
func (o OptionService) IsSearcherEvents() bool {
	val, ok := cache.Cache().Get("searcher_events")
	if !ok {
		return true
	}
	return val.(bool)
}
