package options

import (
	"bitrix-statistic/internal/cache"
	"github.com/sirupsen/logrus"
)

func Set(key string, value interface{}) {

	switch value.(type) {

	case int:
	case float64:
	case string:
	default:
		logrus.Panic("unknown type")
	}

	cache.Cache().Set(key, value)
}

func IsSaveVisits() bool {
	val, ok := cache.Cache().Get("save_visits")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSaveReferrers() bool {
	val, ok := cache.Cache().Get("save_referrers")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSaveHits() bool {
	val, ok := cache.Cache().Get("save_hits")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSaveAdditional() bool {
	val, ok := cache.Cache().Get("save_additional")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSavePathData() bool {
	val, ok := cache.Cache().Get("save_path_data")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsAdvNa() bool {
	val, ok := cache.Cache().Get("adv_na")
	if !ok {
		return false
	}
	return val.(bool)
}

func AvdNaReferer1() string {
	val, ok := cache.Cache().Get("avd_na_referer1")
	if !ok {
		return "NA"
	}
	return val.(string)
}

func AvdNaReferer2() string {
	val, ok := cache.Cache().Get("avd_na_referer2")
	if !ok {
		return "NA"
	}
	return val.(string)
}

func Referer1Syn() string {
	val, ok := cache.Cache().Get("referer1_syn")
	if !ok {
		return "r1"
	}
	return val.(string)
}

func Referer2Syn() string {
	val, ok := cache.Cache().Get("referer2_syn")
	if !ok {
		return "r2"
	}
	return val.(string)
}

func Referer3Syn() string {
	val, ok := cache.Cache().Get("referer3_syn")
	if !ok {
		return "r3"
	}
	return val.(string)
}

func IsRefererCheck() bool {
	val, ok := cache.Cache().Get("referer_check")
	if !ok {
		return false
	}
	return val.(bool)
}

func OpenStatActive() bool {
	val, ok := cache.Cache().Get("open_stat_active")
	if !ok {
		return false
	}
	return val.(bool)
}

func OpenStatR1Template() string {
	val, ok := cache.Cache().Get("open_stat_r1_template")
	if !ok {
		return "#service-name#_#campaign-id#"
	}
	return val.(string)
}

func OpenStatR2Template() string {
	val, ok := cache.Cache().Get("open_stat_r2_template")
	if !ok {
		return "#ad-id#_#source-id#"
	}
	return val.(string)
}

func IsAdvAutoCreate() bool {
	val, ok := cache.Cache().Get("adv_auto_create")
	if !ok {
		return true
	}
	return val.(bool)
}

func AdvGuestDays() uint32 {
	val, ok := cache.Cache().Get("adv_guestDays")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func AdvDays() uint32 {
	val, ok := cache.Cache().Get("adv_days")
	if !ok {
		return 365
	}
	return val.(uint32)
}

func SearcherHitDays() uint32 {
	val, ok := cache.Cache().Get("searcher_hit_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func SearcherDays() uint32 {
	val, ok := cache.Cache().Get("searcher_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func EventsDays() uint32 {
	val, ok := cache.Cache().Get("events_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func EventDynamicDays() uint32 {
	val, ok := cache.Cache().Get("event_dynamic_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func VisitDays() uint32 {
	val, ok := cache.Cache().Get("visit_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func CityDays() uint32 {
	val, ok := cache.Cache().Get("city_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func CountryDays() uint32 {
	val, ok := cache.Cache().Get("country_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func PathDays() uint32 {
	val, ok := cache.Cache().Get("path_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func GuestDays() uint32 {
	val, ok := cache.Cache().Get("guest_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func SessionDays() uint32 {
	val, ok := cache.Cache().Get("session_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func HitDays() uint32 {
	val, ok := cache.Cache().Get("hit_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func PhrasesDays() uint32 {
	val, ok := cache.Cache().Get("phrases_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func RefererListDays() uint32 {
	val, ok := cache.Cache().Get("referer_list_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func RefererDays() uint32 {
	val, ok := cache.Cache().Get("referer_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func RefererTop() uint32 {
	val, ok := cache.Cache().Get("referer_top")
	if !ok {
		return 500
	}
	return val.(uint32)
}

func MaxPathSteps() uint32 {
	val, ok := cache.Cache().Get("max_path_steps")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func OnlineInterval() uint32 {
	val, ok := cache.Cache().Get("online_interval")
	if !ok {
		return 180
	}
	return val.(uint32)
}

func RecordsLimit() uint32 {
	val, ok := cache.Cache().Get("records_limit")
	if !ok {
		return 500
	}
	return val.(uint32)
}

func EventGidBase64Encode() bool {
	val, ok := cache.Cache().Get("event_gid_base64_encode")
	if !ok {
		return true
	}
	return val.(bool)
}

func AdvEventsDefault() string {
	val, ok := cache.Cache().Get("adv_events_default")
	if !ok {
		return "list"
	}
	return val.(string)
}

func UseAutoOptimize() bool {
	val, ok := cache.Cache().Get("use_auto_optimize")
	if !ok {
		return false
	}
	return val.(bool)
}

func BaseCurrency() string {
	val, ok := cache.Cache().Get("base_currency")
	if !ok {
		return "xxx"
	}
	return val.(string)
}

func IsDefenceOn() bool {
	val, ok := cache.Cache().Get("defence_on")
	if !ok {
		return true
	}
	return val.(bool)
}

func DefenceStackTime() uint32 {
	val, ok := cache.Cache().Get("defence_stack_time")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func DefenceMaxStackHits() uint32 {
	val, ok := cache.Cache().Get("defence_max_stack_hits")
	if !ok {
		return 15
	}
	return val.(uint32)
}

func DefenceDelay() uint32 {
	val, ok := cache.Cache().Get("defence_delay")
	if !ok {
		return 300
	}
	return val.(uint32)
}

func IsDefenceLog() bool {
	val, ok := cache.Cache().Get("defence_log")
	if !ok {
		return false
	}
	return val.(bool)
}

func ImportantPageParams() string {
	val, ok := cache.Cache().Get("important_page_params")
	if !ok {
		return "ID, IBLOCK_ID, SECTION_ID, PARENT_ELEMENT_ID, FID, TID, MID, UID, VOTE_ID, print, goto"
	}
	return val.(string)
}

func SkipStatisticWhat() string {
	val, ok := cache.Cache().Get("skip_statistic_what")
	if !ok {
		return "none"
	}
	return val.(string)
}

func SkipStatisticGroups() string {
	val, ok := cache.Cache().Get("skip_statistic_groups")
	if !ok {
		return ""
	}
	return val.(string)
}

func SkipStatisticIpRanges() string {
	val, ok := cache.Cache().Get("skip_statistic_ip_ranges")
	if !ok {
		return ""
	}
	return val.(string)
}

func DirectoryIndex() string {
	val, ok := cache.Cache().Get("directory_index")
	if !ok {
		return ""
	}
	return val.(string)
}

// IsSearcherEvents Учитывать события рекламных кампаний для поисковиков
func IsSearcherEvents() bool {
	val, ok := cache.Cache().Get("searcher_events")
	if !ok {
		return true
	}
	return val.(bool)
}
