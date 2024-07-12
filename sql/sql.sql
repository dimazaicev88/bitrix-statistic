---------------------- ADV -------------------------

create table if not exists adv
(
    `uuid`           UUID,
    `referer1`       String,
    `referer2`       String,
    `cost`           decimal(18, 4) default 0.0000,
    `revenue`        decimal(18, 4) default 0.0000,
    `events_view`    String,
    `guests`         UInt32         default 0,
    `new_guests`     UInt32         default 0,
    `favorites`      UInt32         default 0,
    `hosts`          UInt32         default 0,
    `sessions`       UInt32         default 0,
    `hits`           UInt32         default 0,
    `date_first`     DateTime32('Europe/Moscow'),
    `date_last`      DateTime32('Europe/Moscow'),
    `guests_back`    UInt32         default 0,
    `favorites_back` UInt32         default 0,
    `hosts_back`     UInt32         default 0,
    `sessions_back`  UInt32         default 0,
    `hits_back`      UInt32         default 0,
    `description`    String,
    `priority`       Int32          default 1
) ENGINE = MergeTree
--      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (referer1)
      PRIMARY KEY (referer1);

create table if not exists adv_day
(
    `uuid`            UUID,
    `adv_uuid`        UUID,
    `date_stat`       DateTime32('Europe/Moscow'),
    `guests`          UInt32 default 0,
    `guests_day`      UInt32 default 0,
    `new_guests`      UInt32 default 0,
    `favorites`       UInt32 default 0,
    `hosts`           UInt32 default 0,
    `hosts_day`       UInt32 default 0,
    `sessions`        UInt32 default 0,
    `hits`            UInt32 default 0,
    `guests_back`     UInt32 default 0,
    `guests_day_back` UInt32 default 0,
    `favorites_back`  UInt32 default 0,
    `hosts_back`      UInt32 default 0,
    `hosts_day_back`  UInt32 default 0,
    `sessions_back`   UInt32 default 0,
    `hits_back`       UInt32 default 0
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (`adv_uuid`, `date_stat`);

create table if not exists adv_event
(
    uuid         UUID,
    adv_uuid     UUID,
    event_uuid   UUID,
    counter      UInt32         default 0,
    counter_back UInt32         default 0,
    money        decimal(18, 4) default 0.0000,
    money_back   decimal(18, 4) default 0.0000
) ENGINE = MergeTree;
--       PARTITION BY toYYYYMM(date_stat)
--       ORDER BY (`adv_uuid`, `date_stat`);

create table if not exists adv_event_day
(
    uuid         String,
    adv_uuid     String,
    event_uuid   String,
    date_stat    DateTime32('Europe/Moscow'),
    counter      UInt32,
    counter_back UInt32,
    money        decimal(18, 4) default 0.0000,
    money_back   decimal(18, 4) default 0.0000
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (`adv_uuid`, `event_uuid`, `date_stat`);

create table if not exists adv_guest
(
    uuid           UUID,
    adv_uuid       UUID,
    back           BOOLEAN default false,
    guest_uuid     UUID,
    date_guest_hit DateTime32('Europe/Moscow'),
    date_host_hit  DateTime32('Europe/Moscow'),
    session_uuid   UUID,
    ip             IPv4
) ENGINE = MergeTree
--      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (`adv_uuid`, `guest_uuid`);


create table if not exists adv_page
(
    `uuid`     UUID,
    `adv_uuid` UUID,
    `page`     String,
    `type`     String default 'TO'
) ENGINE = MergeTree
      ORDER BY (`adv_uuid`, `type`);

create table if not exists adv_searcher
(
    uuid          UUID,
    adv_uuid      UUID,
    searcher_uuid UUID
) ENGINE = MergeTree
      ORDER BY (`adv_uuid`, `searcher_uuid`);

----------------------- Browser --------------------------
create table browser
(
    uuid       UUID,
    user_agent String
) ENGINE = MergeTree
      ORDER BY (`user_agent`);

------------------------ City -----------------------------

create table if not exists city
(
    `uuid`       UUID,
    `country_id` String,
    `region`     String,
    `name`       String,
    `xml_id`     String,
    `sessions`   UInt32 default 0,
    `new_guests` UInt32 default 0,
    `hits`       UInt32 default 0,
    `events`     UInt32 default 0
) engine = MergeTree
      ORDER BY (country_id, region, name);

create table if not exists city_day
(
    uuid       UUID,
    city_uuid  UUID,
    date_stat  DateTime32('Europe/Moscow'),
    sessions   UInt32 default 0,
    new_guests UInt32 default 0,
    hits       UInt32 default 0,
    events     UInt32 default 0
) engine = MergeTree
      ORDER BY (city_uuid, date_stat)
      SETTINGS index_granularity = 8192;

create table city_ip
(
    start_ip   UInt32,
    end_ip     UInt32,
    country_id String,
    city_uuid  UUID
) engine = MergeTree
      ORDER BY (end_ip);

------------------ Country ---------------------

create table country
(
    uuid       UUID,
    short_name String,
    name       String,
    sessions   UInt32 default 0,
    new_guests UInt32 default 0,
    hits       UInt32 default 0,
    events     UInt32 default 0
) engine = MergeTree
      ORDER BY (name);

create table country_day
(
    Id         UUID,
    country_id FixedString(2),
    date_stat  DateTime32('Europe/Moscow'),
    sessions   UInt32 default 0,
    new_guests UInt32 default 0,
    hits       UInt32 default 0,
    events     UInt32 default 0
) engine = MergeTree
      ORDER BY (country_id, date_stat);

---------------------day---------------------

create table day
(
    uuid                UUID,
    date_stat           DateTime32('Europe/Moscow'),
    hits                UInt32         default 0,
    hosts               UInt32         default 0,
    sessions            UInt32         default 0,
    events              UInt32         default 0,
    guests              UInt32         default 0,
    new_guests          UInt32         default 0,
    favorites           UInt32         default 0,
    total_hosts         UInt32         default 0,
    am_average_time     decimal(18, 2) default 0.00,
    am_1                UInt32         default 0,
    am_1_3              UInt32         default 0,
    am_3_6              UInt32         default 0,
    am_6_9              UInt32         default 0,
    am_9_12             UInt32         default 0,
    am_12_15            UInt32         default 0,
    am_15_18            UInt32         default 0,
    am_18_21            UInt32         default 0,
    am_21_24            UInt32         default 0,
    am_24               UInt32         default 0,
    ah_average_hits     decimal(18, 2) default 0.00,
    ah_1                UInt32         default 0,
    ah_2_5              UInt32         default 0,
    ah_6_9              UInt32         default 0,
    ah_10_13            UInt32         default 0,
    ah_14_17            UInt32         default 0,
    ah_18_21            UInt32         default 0,
    ah_22_25            UInt32         default 0,
    ah_26_29            UInt32         default 0,
    ah_30_33            UInt32         default 0,
    ah_34               UInt32         default 0,
    hour_host_0         UInt32         default 0,
    hour_host_1         UInt32         default 0,
    hour_host_2         UInt32         default 0,
    hour_host_3         UInt32         default 0,
    hour_host_4         UInt32         default 0,
    hour_host_5         UInt32         default 0,
    hour_host_6         UInt32         default 0,
    hour_host_7         UInt32         default 0,
    hour_host_8         UInt32         default 0,
    hour_host_9         UInt32         default 0,
    hour_host_10        UInt32         default 0,
    hour_host_11        UInt32         default 0,
    hour_host_12        UInt32         default 0,
    hour_host_13        UInt32         default 0,
    hour_host_14        UInt32         default 0,
    hour_host_15        UInt32         default 0,
    hour_host_16        UInt32         default 0,
    hour_host_17        UInt32         default 0,
    hour_host_18        UInt32         default 0,
    hour_host_19        UInt32         default 0,
    hour_host_20        UInt32         default 0,
    hour_host_21        UInt32         default 0,
    hour_host_22        UInt32         default 0,
    hour_host_23        UInt32         default 0,
    hour_guest_0        UInt32         default 0,
    hour_guest_1        UInt32         default 0,
    hour_guest_2        UInt32         default 0,
    hour_guest_3        UInt32         default 0,
    hour_guest_4        UInt32         default 0,
    hour_guest_5        UInt32         default 0,
    hour_guest_6        UInt32         default 0,
    hour_guest_7        UInt32         default 0,
    hour_guest_8        UInt32         default 0,
    hour_guest_9        UInt32         default 0,
    hour_guest_10       UInt32         default 0,
    hour_guest_11       UInt32         default 0,
    hour_guest_12       UInt32         default 0,
    hour_guest_13       UInt32         default 0,
    hour_guest_14       UInt32         default 0,
    hour_guest_15       UInt32         default 0,
    hour_guest_16       UInt32         default 0,
    hour_guest_17       UInt32         default 0,
    hour_guest_18       UInt32         default 0,
    hour_guest_19       UInt32         default 0,
    hour_guest_20       UInt32         default 0,
    hour_guest_21       UInt32         default 0,
    hour_guest_22       UInt32         default 0,
    hour_guest_23       UInt32         default 0,
    hour_new_guest_0    UInt32         default 0,
    hour_new_guest_1    UInt32         default 0,
    hour_new_guest_2    UInt32         default 0,
    hour_new_guest_3    UInt32         default 0,
    hour_new_guest_4    UInt32         default 0,
    hour_new_guest_5    UInt32         default 0,
    hour_new_guest_6    UInt32         default 0,
    hour_new_guest_7    UInt32         default 0,
    hour_new_guest_8    UInt32         default 0,
    hour_new_guest_9    UInt32         default 0,
    hour_new_guest_10   UInt32         default 0,
    hour_new_guest_11   UInt32         default 0,
    hour_new_guest_12   UInt32         default 0,
    hour_new_guest_13   UInt32         default 0,
    hour_new_guest_14   UInt32         default 0,
    hour_new_guest_15   UInt32         default 0,
    hour_new_guest_16   UInt32         default 0,
    hour_new_guest_17   UInt32         default 0,
    hour_new_guest_18   UInt32         default 0,
    hour_new_guest_19   UInt32         default 0,
    hour_new_guest_20   UInt32         default 0,
    hour_new_guest_21   UInt32         default 0,
    hour_new_guest_22   UInt32         default 0,
    hour_new_guest_23   UInt32         default 0,
    hour_session_0      UInt32         default 0,
    hour_session_1      UInt32         default 0,
    hour_session_2      UInt32         default 0,
    hour_session_3      UInt32         default 0,
    hour_session_4      UInt32         default 0,
    hour_session_5      UInt32         default 0,
    hour_session_6      UInt32         default 0,
    hour_session_7      UInt32         default 0,
    hour_session_8      UInt32         default 0,
    hour_session_9      UInt32         default 0,
    hour_session_10     UInt32         default 0,
    hour_session_11     UInt32         default 0,
    hour_session_12     UInt32         default 0,
    hour_session_13     UInt32         default 0,
    hour_session_14     UInt32         default 0,
    hour_session_15     UInt32         default 0,
    hour_session_16     UInt32         default 0,
    hour_session_17     UInt32         default 0,
    hour_session_18     UInt32         default 0,
    hour_session_19     UInt32         default 0,
    hour_session_20     UInt32         default 0,
    hour_session_21     UInt32         default 0,
    hour_session_22     UInt32         default 0,
    hour_session_23     UInt32         default 0,
    hour_hit_0          UInt32         default 0,
    hour_hit_1          UInt32         default 0,
    hour_hit_2          UInt32         default 0,
    hour_hit_3          UInt32         default 0,
    hour_hit_4          UInt32         default 0,
    hour_hit_5          UInt32         default 0,
    hour_hit_6          UInt32         default 0,
    hour_hit_7          UInt32         default 0,
    hour_hit_8          UInt32         default 0,
    hour_hit_9          UInt32         default 0,
    hour_hit_10         UInt32         default 0,
    hour_hit_11         UInt32         default 0,
    hour_hit_12         UInt32         default 0,
    hour_hit_13         UInt32         default 0,
    hour_hit_14         UInt32         default 0,
    hour_hit_15         UInt32         default 0,
    hour_hit_16         UInt32         default 0,
    hour_hit_17         UInt32         default 0,
    hour_hit_18         UInt32         default 0,
    hour_hit_19         UInt32         default 0,
    hour_hit_20         UInt32         default 0,
    hour_hit_21         UInt32         default 0,
    hour_hit_22         UInt32         default 0,
    hour_hit_23         UInt32         default 0,
    hour_event_0        UInt32         default 0,
    hour_event_1        UInt32         default 0,
    hour_event_2        UInt32         default 0,
    hour_event_3        UInt32         default 0,
    hour_event_4        UInt32         default 0,
    hour_event_5        UInt32         default 0,
    hour_event_6        UInt32         default 0,
    hour_event_7        UInt32         default 0,
    hour_event_8        UInt32         default 0,
    hour_event_9        UInt32         default 0,
    hour_event_10       UInt32         default 0,
    hour_event_11       UInt32         default 0,
    hour_event_12       UInt32         default 0,
    hour_event_13       UInt32         default 0,
    hour_event_14       UInt32         default 0,
    hour_event_15       UInt32         default 0,
    hour_event_16       UInt32         default 0,
    hour_event_17       UInt32         default 0,
    hour_event_18       UInt32         default 0,
    hour_event_19       UInt32         default 0,
    hour_event_20       UInt32         default 0,
    hour_event_21       UInt32         default 0,
    hour_event_22       UInt32         default 0,
    hour_event_23       UInt32         default 0,
    hour_favorite_0     UInt32         default 0,
    hour_favorite_1     UInt32         default 0,
    hour_favorite_2     UInt32         default 0,
    hour_favorite_3     UInt32         default 0,
    hour_favorite_4     UInt32         default 0,
    hour_favorite_5     UInt32         default 0,
    hour_favorite_6     UInt32         default 0,
    hour_favorite_7     UInt32         default 0,
    hour_favorite_8     UInt32         default 0,
    hour_favorite_9     UInt32         default 0,
    hour_favorite_10    UInt32         default 0,
    hour_favorite_11    UInt32         default 0,
    hour_favorite_12    UInt32         default 0,
    hour_favorite_13    UInt32         default 0,
    hour_favorite_14    UInt32         default 0,
    hour_favorite_15    UInt32         default 0,
    hour_favorite_16    UInt32         default 0,
    hour_favorite_17    UInt32         default 0,
    hour_favorite_18    UInt32         default 0,
    hour_favorite_19    UInt32         default 0,
    hour_favorite_20    UInt32         default 0,
    hour_favorite_21    UInt32         default 0,
    hour_favorite_22    UInt32         default 0,
    hour_favorite_23    UInt32         default 0,
    weekday_host_0      UInt32         default 0,
    weekday_host_1      UInt32         default 0,
    weekday_host_2      UInt32         default 0,
    weekday_host_3      UInt32         default 0,
    weekday_host_4      UInt32         default 0,
    weekday_host_5      UInt32         default 0,
    weekday_host_6      UInt32         default 0,
    weekday_guest_0     UInt32         default 0,
    weekday_guest_1     UInt32         default 0,
    weekday_guest_2     UInt32         default 0,
    weekday_guest_3     UInt32         default 0,
    weekday_guest_4     UInt32         default 0,
    weekday_guest_5     UInt32         default 0,
    weekday_guest_6     UInt32         default 0,
    weekday_new_guest_0 UInt32         default 0,
    weekday_new_guest_1 UInt32         default 0,
    weekday_new_guest_2 UInt32         default 0,
    weekday_new_guest_3 UInt32         default 0,
    weekday_new_guest_4 UInt32         default 0,
    weekday_new_guest_5 UInt32         default 0,
    weekday_new_guest_6 UInt32         default 0,
    weekday_session_0   UInt32         default 0,
    weekday_session_1   UInt32         default 0,
    weekday_session_2   UInt32         default 0,
    weekday_session_3   UInt32         default 0,
    weekday_session_4   UInt32         default 0,
    weekday_session_5   UInt32         default 0,
    weekday_session_6   UInt32         default 0,
    weekday_hit_0       UInt32         default 0,
    weekday_hit_1       UInt32         default 0,
    weekday_hit_2       UInt32         default 0,
    weekday_hit_3       UInt32         default 0,
    weekday_hit_4       UInt32         default 0,
    weekday_hit_5       UInt32         default 0,
    weekday_hit_6       UInt32         default 0,
    weekday_event_0     UInt32         default 0,
    weekday_event_1     UInt32         default 0,
    weekday_event_2     UInt32         default 0,
    weekday_event_3     UInt32         default 0,
    weekday_event_4     UInt32         default 0,
    weekday_event_5     UInt32         default 0,
    weekday_event_6     UInt32         default 0,
    weekday_favorite_0  UInt32         default 0,
    weekday_favorite_1  UInt32         default 0,
    weekday_favorite_2  UInt32         default 0,
    weekday_favorite_3  UInt32         default 0,
    weekday_favorite_4  UInt32         default 0,
    weekday_favorite_5  UInt32         default 0,
    weekday_favorite_6  UInt32         default 0,
    month_host_1        UInt32         default 0,
    month_host_2        UInt32         default 0,
    month_host_3        UInt32         default 0,
    month_host_4        UInt32         default 0,
    month_host_5        UInt32         default 0,
    month_host_6        UInt32         default 0,
    month_host_7        UInt32         default 0,
    month_host_8        UInt32         default 0,
    month_host_9        UInt32         default 0,
    month_host_10       UInt32         default 0,
    month_host_11       UInt32         default 0,
    month_host_12       UInt32         default 0,
    month_guest_1       UInt32         default 0,
    month_guest_2       UInt32         default 0,
    month_guest_3       UInt32         default 0,
    month_guest_4       UInt32         default 0,
    month_guest_5       UInt32         default 0,
    month_guest_6       UInt32         default 0,
    month_guest_7       UInt32         default 0,
    month_guest_8       UInt32         default 0,
    month_guest_9       UInt32         default 0,
    month_guest_10      UInt32         default 0,
    month_guest_11      UInt32         default 0,
    month_guest_12      UInt32         default 0,
    month_new_guest_1   UInt32         default 0,
    month_new_guest_2   UInt32         default 0,
    month_new_guest_3   UInt32         default 0,
    month_new_guest_4   UInt32         default 0,
    month_new_guest_5   UInt32         default 0,
    month_new_guest_6   UInt32         default 0,
    month_new_guest_7   UInt32         default 0,
    month_new_guest_8   UInt32         default 0,
    month_new_guest_9   UInt32         default 0,
    month_new_guest_10  UInt32         default 0,
    month_new_guest_11  UInt32         default 0,
    month_new_guest_12  UInt32         default 0,
    month_session_1     UInt32         default 0,
    month_session_2     UInt32         default 0,
    month_session_3     UInt32         default 0,
    month_session_4     UInt32         default 0,
    month_session_5     UInt32         default 0,
    month_session_6     UInt32         default 0,
    month_session_7     UInt32         default 0,
    month_session_8     UInt32         default 0,
    month_session_9     UInt32         default 0,
    month_session_10    UInt32         default 0,
    month_session_11    UInt32         default 0,
    month_session_12    UInt32         default 0,
    month_hit_1         UInt32         default 0,
    month_hit_2         UInt32         default 0,
    month_hit_3         UInt32         default 0,
    month_hit_4         UInt32         default 0,
    month_hit_5         UInt32         default 0,
    month_hit_6         UInt32         default 0,
    month_hit_7         UInt32         default 0,
    month_hit_8         UInt32         default 0,
    month_hit_9         UInt32         default 0,
    month_hit_10        UInt32         default 0,
    month_hit_11        UInt32         default 0,
    month_hit_12        UInt32         default 0,
    month_event_1       UInt32         default 0,
    month_event_2       UInt32         default 0,
    month_event_3       UInt32         default 0,
    month_event_4       UInt32         default 0,
    month_event_5       UInt32         default 0,
    month_event_6       UInt32         default 0,
    month_event_7       UInt32         default 0,
    month_event_8       UInt32         default 0,
    month_event_9       UInt32         default 0,
    month_event_10      UInt32         default 0,
    month_event_11      UInt32         default 0,
    month_event_12      UInt32         default 0,
    month_favorite_1    UInt32         default 0,
    month_favorite_2    UInt32         default 0,
    month_favorite_3    UInt32         default 0,
    month_favorite_4    UInt32         default 0,
    month_favorite_5    UInt32         default 0,
    month_favorite_6    UInt32         default 0,
    month_favorite_7    UInt32         default 0,
    month_favorite_8    UInt32         default 0,
    month_favorite_9    UInt32         default 0,
    month_favorite_10   UInt32         default 0,
    month_favorite_11   UInt32         default 0,
    month_favorite_12   UInt32         default 0
);


---------------------------- Event --------------------------------

create table event
(
    uuid              UUID,
    event1            String,
    event2            String,
    money             decimal(18, 4) default 0.0000,
    date_enter        DateTime32('Europe/Moscow'),
    date_cleanup      DateTime32('Europe/Moscow'),
    sort              UInt32         default 100,
    counter           UInt32         default 0,
    ADV_VISIBLE       BOOLEAN        default true,
    NAME              String,
    description       String,
    keep_days         UInt32,
    dynamic_keep_days UInt32,
    diagram_default   BOOLEAN        default true
) engine = MergeTree
      ORDER BY (event1, event2, keep_days);

create table event_day
(
    uuid       UUID,
    date_stat  DateTime32('Europe/Moscow'),
    date_last  DateTime32('Europe/Moscow'),
    event_uuid UInt32         default 0,
    money      decimal(18, 4) default 0.0000,
    counter    UInt32         default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (event_uuid, date_stat);

create table event_list
(
    uuid            UUID,
    event_uuid      int            default 0,
    event3          String,
    money           decimal(18, 4) default 0.0000,
    date_enter      DateTime32('Europe/Moscow'),
    referer_url     String,
    url             String,
    redirect_url    String,
    session_uuid    int,
    guest_uuid      int,
    guest_adv_uuid  int,
    adv_back        BOOLEAN        default false,
    hit_uuid        UUID,
    country_id      FixedString(2),
    keep_days       int,
    chargeback      char           default 'N',
    site_id         FixedString(2),
    referer_site_id FixedString(2)
) engine = MergeTree
      PARTITION BY toYYYYMM(date_enter)
      ORDER BY (date_enter);

----------------------- Guest ---------------------------

create table guest
(
    uuid              UUID,
    timestamp         DateTime32('Europe/Moscow'),
    favorites         UInt8 default 0,
    events            Int32 default 0,
    sessions          Int32 default 0,
    hits              Int32 default 0,
    repair            UInt8 default 0,
    first_session_id  UUID,
    first_date        DateTime32('Europe/Moscow'),
    first_url_from    String,
    first_url_to      String,
    first_url_to_404  UInt8 default 0,
    first_site_id     String,
    first_adv_id      UUID,
    first_referer1    String,
    first_referer2    String,
    first_referer3    String,
    last_session_id   UUID,
    last_date         DateTime32('Europe/Moscow'),
    last_user_id      Int32,
    last_user_auth    UInt8,
    last_url_last     String,
    last_url_last_404 UInt8 default 0,
    last_user_agent   String,
    last_ip           IPv4,
    last_cookie       String,
    last_language     String,
    last_adv_id       UUID,
    last_adv_back     UInt8 default 0,
    last_referer1     String,
    last_referer2     String,
    last_referer3     String,
    last_city_id      UUID,
    token             String
)
    engine = MergeTree
        PARTITION BY toYYYYMM(timestamp)
        ORDER BY timestamp;


----------------------- Hit ---------------------------
create table hit
(
    uuid         UUID,
    session_id   UInt32  default 0,
    date_hit     DateTime32('Europe/Moscow'),
    guest_uuid   UUID,
    new_guest    BOOLEAN default false,
    user_id      int,
    user_auth    BOOLEAN default false,
    url          String,
    url_404      BOOLEAN default false,
    url_from     String,
    ip           IPv4,
    method       String,
    cookies      String,
    user_agent   String,
    stop_list_id int,
    country_id   FixedString(2),
    city_uuid    UUID,
    site_id      FixedString(2)
)
    engine = MergeTree
        PARTITION BY toYYYYMM(date_hit)
        ORDER BY date_hit;

------------------ Page ----------------------

create table page
(
    uuid          UUID,
    date_stat     DateTime32('Europe/Moscow'),
    dir           BOOLEAN default false,
    url           String,
    url_404       BOOLEAN default false,
    url_hash      UInt32,
    site_id       FixedString(2),
    counter       UInt32  default 0,
    enter_counter UInt32  default 0,
    exit_counter  UInt32  default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY date_stat;

---------------------- Path ------------------------

create table path
(
    uuid               UUID,
    path_id            UInt32  default 0,
    parent_path_id     UInt32,
    date_stat          DateTime32('Europe/Moscow'),
    counter            UInt32  default 0,
    counter_abnormal   UInt32  default 0,
    counter_full_path  UInt32  default 0,
    pages              String,
    first_page         String,
    first_page_404     BOOLEAN default false,
    first_page_site_id FixedString(2),
    prev_page          String,
    prev_page_hash     UInt32,
    last_page          String,
    last_page_404      BOOLEAN default false,
    last_page_site_id  FixedString(2),
    last_page_hash     UInt32,
    steps              UInt32  default 1
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY date_stat;

create table path_adv
(
    uuid                   UUID,
    adv_uuid               UUID   default 0,
    path_uuid              UUID   default 0,
    date_stat              DateTime32('Europe/Moscow'),
    counter                UInt32 default 0,
    counter_back           UInt32 default 0,
    counter_full_path      UInt32 default 0,
    counter_full_path_back UInt32 default 0,
    steps                  UInt32 default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY date_stat;


create table path_cache
(
    uuid                    UUID,
    session_uuid            UUID    default 0,
    date_hit                DateTime32('Europe/Moscow'),
    path_uuid               UUID,
    path_pages              String,
    path_first_page         String,
    path_first_page_404     BOOLEAN default false,
    path_first_page_site_id FixedString(2),
    path_last_page          String,
    path_last_page_404      BOOLEAN default false,
    path_last_page_site_id  FixedString(2),
    path_steps              UInt32  default 1,
    is_last_page            BOOLEAN default true
)engine = MergeTree
     PARTITION BY toYYYYMM(date_hit)
     ORDER BY date_hit;


----------------------- Phrase ----------------------------

create table phrase_list
(
    ID          int auto_increment
        primary key,
    date_hit    DateTime32('Europe/Moscow'),
    searcher_id int,
    referer_id  int,
    phrase      String,
    url_from    String,
    url_to      String,
    url_to_404  char default 'N',
    session_id  int,
    site_id     FixedString(2)
);

create
index IX_DATE_HIT
    on b_stat_phrase_list (DATE_HIT);

create
index IX_URL_TO_SEARCHER_ID
    on b_stat_phrase_list (URL_TO(100), SEARCHER_ID);


--------------------- Referer -----------------------------

create table referer
(
    ID         int auto_increment
        primary key,
    date_first DateTime32('Europe/Moscow'),
    date_last  DateTime32('Europe/Moscow'),
    site_name  String,
    sessions   UInt32 default 0,
    hits       UInt32 default 0
);

create
index IX_B_STAT_REFERER_2
    on b_stat_referer (DATE_LAST, ID);

create
index IX_SITE_NAME
    on b_stat_referer (SITE_NAME);

create table referer_list
(
    ID         int auto_increment
        primary key,
    referer_id int,
    date_hit   DateTime32('Europe/Moscow'),
    protocol   String,
    site_name  String,
    url_from   String,
    url_to     String,
    url_to_404 char default 'N',
    session_id int,
    adv_id     int,
    site_id    FixedString(2)
);

create
index IX_DATE_HIT
    on b_stat_referer_list (DATE_HIT);

create
index IX_SITE_NAME
    on b_stat_referer_list (SITE_NAME(100), URL_TO(100));

--------------------- Searcher -------------------------
create table if not exists searcher
(
    `id`                UInt32 auto_increment,
    `date_cleanup`      DateTime32('Europe/Moscow'),
    `total_hits`        UInt32  default '0',
    `save_statistic`    BOOLEAN default true,
    `active`            BOOLEAN default true,
    `name`              String,
    `user_agent`        String,
    `diagram_default`   BOOLEAN default false,
    `hit_keep_days`     UInt32,
    `dynamic_keep_days` UInt32,
    `phrases`           UInt32  default '0',
    `phrases_hits`      UInt32  default '0',
    `check_activity`    BOOLEAN default true,
    primary             key (`id`)
);
CREATE
INDEX IX_SEARCHER_1 ON searcher (`hit_keep_days`);

create table if not exists searcher_day
(
    `id`          UInt32 auto_increment,
    `date_stat`   date,
    `date_last`   DateTime32('Europe/Moscow'),
    `searcher_id` UInt32 default '0',
    `total_hits`  UInt32 default '0',
    primary       key (`id`),
    index IX_SEARCHER_ID_DATE_STAT (`searcher_id`, `date_stat`)
);

create table if not exists searcher_hit
(
    `uuid`          UUID,
    `date_hit`      DateTime32('Europe/Moscow'),
    `searcher_id`   UInt32  default '0',
    `url`           String,
    `url_404`       BOOLEAN default false,
    `ip`            String,
    `user_agent`    String,
    `hit_keep_days` UInt32,
    `site_id`       FixedString(2)
);
CREATE
INDEX IX_SEARCHER_HIT_1 ON searcher_hit (`searcher_id`, `date_hit`);
CREATE
INDEX IX_SEARCHER_HIT_2 ON searcher_hit (`hit_keep_days`, `date_hit`);

create table searcher_params
(
    ID          int auto_increment
        primary key,
    searcher_id int default 0,
    domain      String,
    variable    String,
    char_set    String
);

create
index IX_SEARCHER_DOMAIN
    on b_stat_searcher_params (SEARCHER_ID, DOMAIN);


--------------------- session ---------------------------

create table if not exists session
(
    uuid          UUID,
    guest_id      UUID,
    new_guest     UInt8,
    user_id       Int32,
    user_auth     UInt8,
    events        Int32 default 0,
    hits          Int32 default 0,
    favorites     UInt8,
    url_from      String,
    url_to        String,
    url_to_404    UInt8,
    url_last      String,
    url_last_404  UInt8,
    user_agent    String,
    date_stat     DateTime32('Europe/Moscow'),
    date_first    DateTime32('Europe/Moscow'),
    date_last     DateTime32('Europe/Moscow'),
    ip_first      IPv4,
    ip_last       IPv4,
    first_hit_id  UUID,
    last_hit_id   UUID,
    phpsessid     String,
    adv_id        UUID,
    adv_back      UInt8,
    referer1      String,
    referer2      String,
    referer3      String,
    STOP_LIST_ID  UUID,
--     COUNTRY_ID      FixedString(2) ,
    FIRST_SITE_ID String,
    LAST_SITE_ID  String
--     CITY_ID         UInt32
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (uuid, date_stat)
      PRIMARY KEY (uuid, date_stat);



