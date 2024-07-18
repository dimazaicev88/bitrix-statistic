package config

type AppOptions struct {
	SaveVisits         bool
	SaveReferrers      bool
	SaveHits           bool
	SaveAdditional     bool
	SaveSessionData    bool
	SavePathData       bool
	AdvNa              bool
	AvdNaReferer1      string
	AvdNaReferer2      string
	Referer1Syn        string // Синонимы
	Referer2Syn        string // Синонимы
	Referer3Syn        string // Синонимы
	RefererCheck       bool
	OpenStatActive     bool
	OpenStatR1Template string
	OpenStatR2Template string
	AdvAutoCreate      bool
	AdvGuestDays       uint16
	AdvDays            uint16
	SearcherHitDays    uint16
	SearcherDays       uint16
	EventsDays         uint16
	EventDynamicDays   uint16
	VisitDays          uint16
	CityDays           uint16
	CountryDays        uint16
	PathDays           uint16
	GuestDays          uint16
	SessionDays        uint16
	HitDays            uint16
	PhrasesDays        uint16
	RefererListDays    uint16
	RefererDays        uint16
	RefererTop         uint16
	MaxPathSteps       uint16
	OnlineInterval     uint16
	RecordsLimit       uint16
	//event_gid_site_id = > COption::GetOptionString(main, cookie_name, BITRIX_SM),
	EventGidBase64Encode       bool
	AdvEventsDefault           string
	UseAutoOptimize            bool
	EventsLoadHandlersPath     string
	UserEventsLoadHandlersPath string
	BaseCurrency               string
	GraphWeight                uint16
	GraphHeight                uint16
	DiagramDiameter            uint16
	DefenceOn                  bool
	DefenceStackTime           uint16
	DefenceMaxStackHits        uint16
	DefenceDelay               uint16
	DefenceLog                 bool
	ImportantPageParams        string
	StatListTopSize            uint16
	AdvDetailTopSize           uint16
	SkipStatisticWhat          string
	SkipStatisticGroups        string
	SkipStatisticIpRanges      string
	DirectoryIndex             string
	SearcherEvents             bool
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
//SEARCHER_EVENTS => Y, //Учитывать события рекламных кампаний для поисковиков
