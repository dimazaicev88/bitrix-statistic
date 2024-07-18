package optioncache

import (
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	once       sync.Once
	otterCache otter.Cache[string, interface{}]
	builder    *otter.Builder[string, interface{}]
)

func init() {
	once.Do(func() {
		var err error
		builder = otter.MustBuilder[string, interface{}](100)
		otterCache, err = builder[string, interface{}].CollectStats().
			Build()

		if err != nil {
			logrus.Fatal(err)
		}
	})
}

func Set(key string, value interface{}) {

	switch value.(type) {

	case int:
	case float64:
	case string:
	default:
		logrus.Panic("unknown type")
	}

	otterCache.Set(key, value)
}

func IsSaveVisits() bool {
	val, ok := otterCache.Get("save_visits")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSaveReferrers() bool {
	val, ok := otterCache.Get("save_referrers")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSaveHits() bool {
	val, ok := otterCache.Get("save_hits")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSaveAdditional() bool {
	val, ok := otterCache.Get("save_additional")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsSavePathData() bool {
	val, ok := otterCache.Get("save_path_data")
	if !ok {
		return true
	}
	return val.(bool)
}

func IsAdvNa() bool {
	val, ok := otterCache.Get("adv_na")
	if !ok {
		return false
	}
	return val.(bool)
}

func AvdNaReferer1() string {
	val, ok := otterCache.Get("avd_na_referer1")
	if !ok {
		return "NA"
	}
	return val.(string)
}

func AvdNaReferer2() string {
	val, ok := otterCache.Get("avd_na_referer2")
	if !ok {
		return "NA"
	}
	return val.(string)
}

func Referer1Syn() string {
	val, ok := otterCache.Get("referer1_syn")
	if !ok {
		return "r1"
	}
	return val.(string)
}

func Referer2Syn() string {
	val, ok := otterCache.Get("referer2_syn")
	if !ok {
		return "r2"
	}
	return val.(string)
}

func Referer3Syn() string {
	val, ok := otterCache.Get("referer3_syn")
	if !ok {
		return "r3"
	}
	return val.(string)
}

func IsRefererCheck() bool {
	val, ok := otterCache.Get("referer_check")
	if !ok {
		return false
	}
	return val.(bool)
}

func OpenStatActive() bool {
	val, ok := otterCache.Get("open_stat_active")
	if !ok {
		return false
	}
	return val.(bool)
}

func OpenStatR1Template() string {
	val, ok := otterCache.Get("open_stat_r1_template")
	if !ok {
		return "#service-name#_#campaign-id#"
	}
	return val.(string)
}

func OpenStatR2Template() string {
	val, ok := otterCache.Get("open_stat_r2_template")
	if !ok {
		return "#ad-id#_#source-id#"
	}
	return val.(string)
}

func IsAdvAutoCreate() bool {
	val, ok := otterCache.Get("adv_auto_create")
	if !ok {
		return true
	}
	return val.(bool)
}

func AdvGuestDays() uint32 {
	val, ok := otterCache.Get("adv_guestDays")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func AdvDays() uint32 {
	val, ok := otterCache.Get("adv_days")
	if !ok {
		return 365
	}
	return val.(uint32)
}

func SearcherHitDays() uint32 {
	val, ok := otterCache.Get("searcher_hit_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func SearcherDays() uint32 {
	val, ok := otterCache.Get("searcher_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func EventsDays() uint32 {
	val, ok := otterCache.Get("events_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func EventDynamicDays() uint32 {
	val, ok := otterCache.Get("event_dynamic_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func VisitDays() uint32 {
	val, ok := otterCache.Get("visit_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func CityDays() uint32 {
	val, ok := otterCache.Get("city_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func CountryDays() uint32 {
	val, ok := otterCache.Get("country_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func PathDays() uint32 {
	val, ok := otterCache.Get("path_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func GuestDays() uint32 {
	val, ok := otterCache.Get("guest_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func SessionDays() uint32 {
	val, ok := otterCache.Get("session_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func HitDays() uint32 {
	val, ok := otterCache.Get("hit_days")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func PhrasesDays() uint32 {
	val, ok := otterCache.Get("phrases_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func RefererListDays() uint32 {
	val, ok := otterCache.Get("referer_list_days")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func RefererDays() uint32 {
	val, ok := otterCache.Get("referer_days")
	if !ok {
		return 360
	}
	return val.(uint32)
}

func RefererTop() uint32 {
	val, ok := otterCache.Get("referer_top")
	if !ok {
		return 500
	}
	return val.(uint32)
}

func MaxPathSteps() uint32 {
	val, ok := otterCache.Get("max_path_steps")
	if !ok {
		return 3
	}
	return val.(uint32)
}

func OnlineInterval() uint32 {
	val, ok := otterCache.Get("online_interval")
	if !ok {
		return 180
	}
	return val.(uint32)
}

func RecordsLimit() uint32 {
	val, ok := otterCache.Get("records_limit")
	if !ok {
		return 500
	}
	return val.(uint32)
}

func EventGidBase64Encode() bool {
	val, ok := otterCache.Get("event_gid_base64_encode")
	if !ok {
		return true
	}
	return val.(bool)
}

func AdvEventsDefault() string {
	val, ok := otterCache.Get("adv_events_default")
	if !ok {
		return "list"
	}
	return val.(string)
}

func UseAutoOptimize() bool {
	val, ok := otterCache.Get("use_auto_optimize")
	if !ok {
		return false
	}
	return val.(bool)
}

func BaseCurrency() string {
	val, ok := otterCache.Get("base_currency")
	if !ok {
		return "xxx"
	}
	return val.(string)
}

func IsDefenceOn() bool {
	val, ok := otterCache.Get("defence_on")
	if !ok {
		return true
	}
	return val.(bool)
}

func DefenceStackTime() uint32 {
	val, ok := otterCache.Get("defence_stack_time")
	if !ok {
		return 10
	}
	return val.(uint32)
}

func DefenceMaxStackHits() uint32 {
	val, ok := otterCache.Get("defence_max_stack_hits")
	if !ok {
		return 15
	}
	return val.(uint32)
}

func DefenceDelay() uint32 {
	val, ok := otterCache.Get("defence_delay")
	if !ok {
		return 300
	}
	return val.(uint32)
}

func IsDefenceLog() bool {
	val, ok := otterCache.Get("defence_log")
	if !ok {
		return false
	}
	return val.(bool)
}

func ImportantPageParams() string {
	val, ok := otterCache.Get("important_page_params")
	if !ok {
		return "ID, IBLOCK_ID, SECTION_ID, PARENT_ELEMENT_ID, FID, TID, MID, UID, VOTE_ID, print, goto"
	}
	return val.(string)
}

func SkipStatisticWhat() string {
	val, ok := otterCache.Get("skip_statistic_what")
	if !ok {
		return "none"
	}
	return val.(string)
}

func SkipStatisticGroups() string {
	val, ok := otterCache.Get("skip_statistic_groups")
	if !ok {
		return ""
	}
	return val.(string)
}

func SkipStatisticIpRanges() string {
	val, ok := otterCache.Get("skip_statistic_ip_ranges")
	if !ok {
		return ""
	}
	return val.(string)
}

func DirectoryIndex() string {
	val, ok := otterCache.Get("directory_index")
	if !ok {
		return ""
	}
	return val.(string)
}

func IsSearcherEvents() bool {
	val, ok := otterCache.Get("searcher_events")
	if !ok {
		return true
	}
	return val.(bool)
}

func Close() {
	otterCache.Close()
}

type AppOptions struct {
}

//SAVE_REFERERS = > Y,
//SAVE_HITS = > Y,
//SAVE_ADDITIONAL = > Y,
//SAVE_SESSION_DATA = > Y,
//SAVE_PATH_DATA = > Y,
//ADV_NA = > N,
//AVD_NA_REFERER1 = > NA,
//AVD_NA_REFERER2 = > NA,
//REFERER1_SYN = > referrer1, r1, // Синонимы
//REFERER2_SYN = > referrer2, r2, // Синонимы
//REFERER3_SYN = > referrer3, r3, // Синонимы
//REFERER_CHECK = > Y,
//OPENSTAT_ACTIVE = > N,
//OPENSTAT_R1_TEMPLATE = > #service-name#_#campaign-id#,
//OPENSTAT_R2_TEMPLATE = > #ad-id#_#source-id#,
//ADV_AUTO_CREATE = > Y,
//ADV_GUEST_DAYS = > 3,
//ADV_DAYS = > 360,
//SEARCHER_HIT_DAYS = > 3,
//SEARCHER_DAYS = > 360,
//EVENTS_DAYS = > 3,
//EVENT_DYNAMIC_DAYS = > 360,
//VISIT_DAYS = > 10,
//CITY_DAYS = > 360,
//COUNTRY_DAYS = > 360,
//PATH_DAYS = > 10,
//GUEST_DAYS = > 3,
//SESSION_DAYS = > 3,
//HIT_DAYS = > 3,
//PHRASES_DAYS = > 10,
//REFERER_LIST_DAYS = > 10,
//REFERER_DAYS = > 360,
//REFERER_TOP = > 500,
//MAX_PATH_STEPS = > 3,
//ONLINE_INTERVAL = > 180,
//RECORDS_LIMIT = > 500,
//EVENT_GID_SITE_ID = > COption::GetOptionString(main, cookie_name, BITRIX_SM),
//EVENT_GID_BASE64_ENCODE = > Y,
//ADV_EVENTS_DEFAULT = > list,
//USE_AUTO_OPTIMIZE = > N,
//EVENTS_LOAD_HANDLERS_PATH = > /bitrix/modules/statistic/loading/,
//USER_EVENTS_LOAD_HANDLERS_PATH => BX_PERSONAL_ROOT./php_interface/include/statistic/,
//BASE_CURRENCY = > xxx,
//GRAPH_WEIGHT = > 576,
//GRAPH_HEIGHT = > 400,
//DIAGRAM_DIAMETER = > 180,
//DEFENCE_ON = > Y,
//DEFENCE_STACK_TIME = > 10,
//DEFENCE_MAX_STACK_HITS = > 15,
//DEFENCE_DELAY = > 300,
//DEFENCE_LOG = > N,
//IMPORTANT_PAGE_PARAMS = > ID, IBLOCK_ID, SECTION_ID, PARENT_ELEMENT_ID, FID, TID, MID, UID, VOTE_ID, print, goto,
//STAT_LIST_TOP_SIZE = > 10,
//ADV_DETAIL_TOP_SIZE = > 10,
//SKIP_STATISTIC_WHAT = > none,
//SKIP_STATISTIC_GROUPS = >,
//SKIP_STATISTIC_IP_RANGES = >,
//DIRECTORY_INDEX = >,
//SEARCHER_EVENTS => Y,
