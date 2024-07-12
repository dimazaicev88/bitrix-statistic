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
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (`adv_uuid`, `date_stat`);

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
    ID                  int(18) auto_increment
        primary key,
    DATE_STAT           date null,
    HITS                int(18)        default 0    not null,
    C_HOSTS             int(18)        default 0    not null,
    SESSIONS            int(18)        default 0    not null,
    C_EVENTS            int(18)        default 0    not null,
    GUESTS              int(18)        default 0    not null,
    NEW_GUESTS          int(18)        default 0    not null,
    FAVORITES           int(18)        default 0    not null,
    TOTAL_HOSTS         int(18)        default 0    not null,
    AM_AVERAGE_TIME     decimal(18, 2) default 0.00 not null,
    AM_1                int(18)        default 0    not null,
    AM_1_3              int(18)        default 0    not null,
    AM_3_6              int(18)        default 0    not null,
    AM_6_9              int(18)        default 0    not null,
    AM_9_12             int(18)        default 0    not null,
    AM_12_15            int(18)        default 0    not null,
    AM_15_18            int(18)        default 0    not null,
    AM_18_21            int(18)        default 0    not null,
    AM_21_24            int(18)        default 0    not null,
    AM_24               int(18)        default 0    not null,
    AH_AVERAGE_HITS     decimal(18, 2) default 0.00 not null,
    AH_1                int(18)        default 0    not null,
    AH_2_5              int(18)        default 0    not null,
    AH_6_9              int(18)        default 0    not null,
    AH_10_13            int(18)        default 0    not null,
    AH_14_17            int(18)        default 0    not null,
    AH_18_21            int(18)        default 0    not null,
    AH_22_25            int(18)        default 0    not null,
    AH_26_29            int(18)        default 0    not null,
    AH_30_33            int(18)        default 0    not null,
    AH_34               int(18)        default 0    not null,
    HOUR_HOST_0         int(18)        default 0    not null,
    HOUR_HOST_1         int(18)        default 0    not null,
    HOUR_HOST_2         int(18)        default 0    not null,
    HOUR_HOST_3         int(18)        default 0    not null,
    HOUR_HOST_4         int(18)        default 0    not null,
    HOUR_HOST_5         int(18)        default 0    not null,
    HOUR_HOST_6         int(18)        default 0    not null,
    HOUR_HOST_7         int(18)        default 0    not null,
    HOUR_HOST_8         int(18)        default 0    not null,
    HOUR_HOST_9         int(18)        default 0    not null,
    HOUR_HOST_10        int(18)        default 0    not null,
    HOUR_HOST_11        int(18)        default 0    not null,
    HOUR_HOST_12        int(18)        default 0    not null,
    HOUR_HOST_13        int(18)        default 0    not null,
    HOUR_HOST_14        int(18)        default 0    not null,
    HOUR_HOST_15        int(18)        default 0    not null,
    HOUR_HOST_16        int(18)        default 0    not null,
    HOUR_HOST_17        int(18)        default 0    not null,
    HOUR_HOST_18        int(18)        default 0    not null,
    HOUR_HOST_19        int(18)        default 0    not null,
    HOUR_HOST_20        int(18)        default 0    not null,
    HOUR_HOST_21        int(18)        default 0    not null,
    HOUR_HOST_22        int(18)        default 0    not null,
    HOUR_HOST_23        int(18)        default 0    not null,
    HOUR_GUEST_0        int(18)        default 0    not null,
    HOUR_GUEST_1        int(18)        default 0    not null,
    HOUR_GUEST_2        int(18)        default 0    not null,
    HOUR_GUEST_3        int(18)        default 0    not null,
    HOUR_GUEST_4        int(18)        default 0    not null,
    HOUR_GUEST_5        int(18)        default 0    not null,
    HOUR_GUEST_6        int(18)        default 0    not null,
    HOUR_GUEST_7        int(18)        default 0    not null,
    HOUR_GUEST_8        int(18)        default 0    not null,
    HOUR_GUEST_9        int(18)        default 0    not null,
    HOUR_GUEST_10       int(18)        default 0    not null,
    HOUR_GUEST_11       int(18)        default 0    not null,
    HOUR_GUEST_12       int(18)        default 0    not null,
    HOUR_GUEST_13       int(18)        default 0    not null,
    HOUR_GUEST_14       int(18)        default 0    not null,
    HOUR_GUEST_15       int(18)        default 0    not null,
    HOUR_GUEST_16       int(18)        default 0    not null,
    HOUR_GUEST_17       int(18)        default 0    not null,
    HOUR_GUEST_18       int(18)        default 0    not null,
    HOUR_GUEST_19       int(18)        default 0    not null,
    HOUR_GUEST_20       int(18)        default 0    not null,
    HOUR_GUEST_21       int(18)        default 0    not null,
    HOUR_GUEST_22       int(18)        default 0    not null,
    HOUR_GUEST_23       int(18)        default 0    not null,
    HOUR_NEW_GUEST_0    int(18)        default 0    not null,
    HOUR_NEW_GUEST_1    int(18)        default 0    not null,
    HOUR_NEW_GUEST_2    int(18)        default 0    not null,
    HOUR_NEW_GUEST_3    int(18)        default 0    not null,
    HOUR_NEW_GUEST_4    int(18)        default 0    not null,
    HOUR_NEW_GUEST_5    int(18)        default 0    not null,
    HOUR_NEW_GUEST_6    int(18)        default 0    not null,
    HOUR_NEW_GUEST_7    int(18)        default 0    not null,
    HOUR_NEW_GUEST_8    int(18)        default 0    not null,
    HOUR_NEW_GUEST_9    int(18)        default 0    not null,
    HOUR_NEW_GUEST_10   int(18)        default 0    not null,
    HOUR_NEW_GUEST_11   int(18)        default 0    not null,
    HOUR_NEW_GUEST_12   int(18)        default 0    not null,
    HOUR_NEW_GUEST_13   int(18)        default 0    not null,
    HOUR_NEW_GUEST_14   int(18)        default 0    not null,
    HOUR_NEW_GUEST_15   int(18)        default 0    not null,
    HOUR_NEW_GUEST_16   int(18)        default 0    not null,
    HOUR_NEW_GUEST_17   int(18)        default 0    not null,
    HOUR_NEW_GUEST_18   int(18)        default 0    not null,
    HOUR_NEW_GUEST_19   int(18)        default 0    not null,
    HOUR_NEW_GUEST_20   int(18)        default 0    not null,
    HOUR_NEW_GUEST_21   int(18)        default 0    not null,
    HOUR_NEW_GUEST_22   int(18)        default 0    not null,
    HOUR_NEW_GUEST_23   int(18)        default 0    not null,
    HOUR_SESSION_0      int(18)        default 0    not null,
    HOUR_SESSION_1      int(18)        default 0    not null,
    HOUR_SESSION_2      int(18)        default 0    not null,
    HOUR_SESSION_3      int(18)        default 0    not null,
    HOUR_SESSION_4      int(18)        default 0    not null,
    HOUR_SESSION_5      int(18)        default 0    not null,
    HOUR_SESSION_6      int(18)        default 0    not null,
    HOUR_SESSION_7      int(18)        default 0    not null,
    HOUR_SESSION_8      int(18)        default 0    not null,
    HOUR_SESSION_9      int(18)        default 0    not null,
    HOUR_SESSION_10     int(18)        default 0    not null,
    HOUR_SESSION_11     int(18)        default 0    not null,
    HOUR_SESSION_12     int(18)        default 0    not null,
    HOUR_SESSION_13     int(18)        default 0    not null,
    HOUR_SESSION_14     int(18)        default 0    not null,
    HOUR_SESSION_15     int(18)        default 0    not null,
    HOUR_SESSION_16     int(18)        default 0    not null,
    HOUR_SESSION_17     int(18)        default 0    not null,
    HOUR_SESSION_18     int(18)        default 0    not null,
    HOUR_SESSION_19     int(18)        default 0    not null,
    HOUR_SESSION_20     int(18)        default 0    not null,
    HOUR_SESSION_21     int(18)        default 0    not null,
    HOUR_SESSION_22     int(18)        default 0    not null,
    HOUR_SESSION_23     int(18)        default 0    not null,
    HOUR_HIT_0          int(18)        default 0    not null,
    HOUR_HIT_1          int(18)        default 0    not null,
    HOUR_HIT_2          int(18)        default 0    not null,
    HOUR_HIT_3          int(18)        default 0    not null,
    HOUR_HIT_4          int(18)        default 0    not null,
    HOUR_HIT_5          int(18)        default 0    not null,
    HOUR_HIT_6          int(18)        default 0    not null,
    HOUR_HIT_7          int(18)        default 0    not null,
    HOUR_HIT_8          int(18)        default 0    not null,
    HOUR_HIT_9          int(18)        default 0    not null,
    HOUR_HIT_10         int(18)        default 0    not null,
    HOUR_HIT_11         int(18)        default 0    not null,
    HOUR_HIT_12         int(18)        default 0    not null,
    HOUR_HIT_13         int(18)        default 0    not null,
    HOUR_HIT_14         int(18)        default 0    not null,
    HOUR_HIT_15         int(18)        default 0    not null,
    HOUR_HIT_16         int(18)        default 0    not null,
    HOUR_HIT_17         int(18)        default 0    not null,
    HOUR_HIT_18         int(18)        default 0    not null,
    HOUR_HIT_19         int(18)        default 0    not null,
    HOUR_HIT_20         int(18)        default 0    not null,
    HOUR_HIT_21         int(18)        default 0    not null,
    HOUR_HIT_22         int(18)        default 0    not null,
    HOUR_HIT_23         int(18)        default 0    not null,
    HOUR_EVENT_0        int(18)        default 0    not null,
    HOUR_EVENT_1        int(18)        default 0    not null,
    HOUR_EVENT_2        int(18)        default 0    not null,
    HOUR_EVENT_3        int(18)        default 0    not null,
    HOUR_EVENT_4        int(18)        default 0    not null,
    HOUR_EVENT_5        int(18)        default 0    not null,
    HOUR_EVENT_6        int(18)        default 0    not null,
    HOUR_EVENT_7        int(18)        default 0    not null,
    HOUR_EVENT_8        int(18)        default 0    not null,
    HOUR_EVENT_9        int(18)        default 0    not null,
    HOUR_EVENT_10       int(18)        default 0    not null,
    HOUR_EVENT_11       int(18)        default 0    not null,
    HOUR_EVENT_12       int(18)        default 0    not null,
    HOUR_EVENT_13       int(18)        default 0    not null,
    HOUR_EVENT_14       int(18)        default 0    not null,
    HOUR_EVENT_15       int(18)        default 0    not null,
    HOUR_EVENT_16       int(18)        default 0    not null,
    HOUR_EVENT_17       int(18)        default 0    not null,
    HOUR_EVENT_18       int(18)        default 0    not null,
    HOUR_EVENT_19       int(18)        default 0    not null,
    HOUR_EVENT_20       int(18)        default 0    not null,
    HOUR_EVENT_21       int(18)        default 0    not null,
    HOUR_EVENT_22       int(18)        default 0    not null,
    HOUR_EVENT_23       int(18)        default 0    not null,
    HOUR_FAVORITE_0     int(18)        default 0    not null,
    HOUR_FAVORITE_1     int(18)        default 0    not null,
    HOUR_FAVORITE_2     int(18)        default 0    not null,
    HOUR_FAVORITE_3     int(18)        default 0    not null,
    HOUR_FAVORITE_4     int(18)        default 0    not null,
    HOUR_FAVORITE_5     int(18)        default 0    not null,
    HOUR_FAVORITE_6     int(18)        default 0    not null,
    HOUR_FAVORITE_7     int(18)        default 0    not null,
    HOUR_FAVORITE_8     int(18)        default 0    not null,
    HOUR_FAVORITE_9     int(18)        default 0    not null,
    HOUR_FAVORITE_10    int(18)        default 0    not null,
    HOUR_FAVORITE_11    int(18)        default 0    not null,
    HOUR_FAVORITE_12    int(18)        default 0    not null,
    HOUR_FAVORITE_13    int(18)        default 0    not null,
    HOUR_FAVORITE_14    int(18)        default 0    not null,
    HOUR_FAVORITE_15    int(18)        default 0    not null,
    HOUR_FAVORITE_16    int(18)        default 0    not null,
    HOUR_FAVORITE_17    int(18)        default 0    not null,
    HOUR_FAVORITE_18    int(18)        default 0    not null,
    HOUR_FAVORITE_19    int(18)        default 0    not null,
    HOUR_FAVORITE_20    int(18)        default 0    not null,
    HOUR_FAVORITE_21    int(18)        default 0    not null,
    HOUR_FAVORITE_22    int(18)        default 0    not null,
    HOUR_FAVORITE_23    int(18)        default 0    not null,
    WEEKDAY_HOST_0      int(18)        default 0    not null,
    WEEKDAY_HOST_1      int(18)        default 0    not null,
    WEEKDAY_HOST_2      int(18)        default 0    not null,
    WEEKDAY_HOST_3      int(18)        default 0    not null,
    WEEKDAY_HOST_4      int(18)        default 0    not null,
    WEEKDAY_HOST_5      int(18)        default 0    not null,
    WEEKDAY_HOST_6      int(18)        default 0    not null,
    WEEKDAY_GUEST_0     int(18)        default 0    not null,
    WEEKDAY_GUEST_1     int(18)        default 0    not null,
    WEEKDAY_GUEST_2     int(18)        default 0    not null,
    WEEKDAY_GUEST_3     int(18)        default 0    not null,
    WEEKDAY_GUEST_4     int(18)        default 0    not null,
    WEEKDAY_GUEST_5     int(18)        default 0    not null,
    WEEKDAY_GUEST_6     int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_0 int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_1 int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_2 int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_3 int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_4 int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_5 int(18)        default 0    not null,
    WEEKDAY_NEW_GUEST_6 int(18)        default 0    not null,
    WEEKDAY_SESSION_0   int(18)        default 0    not null,
    WEEKDAY_SESSION_1   int(18)        default 0    not null,
    WEEKDAY_SESSION_2   int(18)        default 0    not null,
    WEEKDAY_SESSION_3   int(18)        default 0    not null,
    WEEKDAY_SESSION_4   int(18)        default 0    not null,
    WEEKDAY_SESSION_5   int(18)        default 0    not null,
    WEEKDAY_SESSION_6   int(18)        default 0    not null,
    WEEKDAY_HIT_0       int(18)        default 0    not null,
    WEEKDAY_HIT_1       int(18)        default 0    not null,
    WEEKDAY_HIT_2       int(18)        default 0    not null,
    WEEKDAY_HIT_3       int(18)        default 0    not null,
    WEEKDAY_HIT_4       int(18)        default 0    not null,
    WEEKDAY_HIT_5       int(18)        default 0    not null,
    WEEKDAY_HIT_6       int(18)        default 0    not null,
    WEEKDAY_EVENT_0     int(18)        default 0    not null,
    WEEKDAY_EVENT_1     int(18)        default 0    not null,
    WEEKDAY_EVENT_2     int(18)        default 0    not null,
    WEEKDAY_EVENT_3     int(18)        default 0    not null,
    WEEKDAY_EVENT_4     int(18)        default 0    not null,
    WEEKDAY_EVENT_5     int(18)        default 0    not null,
    WEEKDAY_EVENT_6     int(18)        default 0    not null,
    WEEKDAY_FAVORITE_0  int(18)        default 0    not null,
    WEEKDAY_FAVORITE_1  int(18)        default 0    not null,
    WEEKDAY_FAVORITE_2  int(18)        default 0    not null,
    WEEKDAY_FAVORITE_3  int(18)        default 0    not null,
    WEEKDAY_FAVORITE_4  int(18)        default 0    not null,
    WEEKDAY_FAVORITE_5  int(18)        default 0    not null,
    WEEKDAY_FAVORITE_6  int(18)        default 0    not null,
    MONTH_HOST_1        int(18)        default 0    not null,
    MONTH_HOST_2        int(18)        default 0    not null,
    MONTH_HOST_3        int(18)        default 0    not null,
    MONTH_HOST_4        int(18)        default 0    not null,
    MONTH_HOST_5        int(18)        default 0    not null,
    MONTH_HOST_6        int(18)        default 0    not null,
    MONTH_HOST_7        int(18)        default 0    not null,
    MONTH_HOST_8        int(18)        default 0    not null,
    MONTH_HOST_9        int(18)        default 0    not null,
    MONTH_HOST_10       int(18)        default 0    not null,
    MONTH_HOST_11       int(18)        default 0    not null,
    MONTH_HOST_12       int(18)        default 0    not null,
    MONTH_GUEST_1       int(18)        default 0    not null,
    MONTH_GUEST_2       int(18)        default 0    not null,
    MONTH_GUEST_3       int(18)        default 0    not null,
    MONTH_GUEST_4       int(18)        default 0    not null,
    MONTH_GUEST_5       int(18)        default 0    not null,
    MONTH_GUEST_6       int(18)        default 0    not null,
    MONTH_GUEST_7       int(18)        default 0    not null,
    MONTH_GUEST_8       int(18)        default 0    not null,
    MONTH_GUEST_9       int(18)        default 0    not null,
    MONTH_GUEST_10      int(18)        default 0    not null,
    MONTH_GUEST_11      int(18)        default 0    not null,
    MONTH_GUEST_12      int(18)        default 0    not null,
    MONTH_NEW_GUEST_1   int(18)        default 0    not null,
    MONTH_NEW_GUEST_2   int(18)        default 0    not null,
    MONTH_NEW_GUEST_3   int(18)        default 0    not null,
    MONTH_NEW_GUEST_4   int(18)        default 0    not null,
    MONTH_NEW_GUEST_5   int(18)        default 0    not null,
    MONTH_NEW_GUEST_6   int(18)        default 0    not null,
    MONTH_NEW_GUEST_7   int(18)        default 0    not null,
    MONTH_NEW_GUEST_8   int(18)        default 0    not null,
    MONTH_NEW_GUEST_9   int(18)        default 0    not null,
    MONTH_NEW_GUEST_10  int(18)        default 0    not null,
    MONTH_NEW_GUEST_11  int(18)        default 0    not null,
    MONTH_NEW_GUEST_12  int(18)        default 0    not null,
    MONTH_SESSION_1     int(18)        default 0    not null,
    MONTH_SESSION_2     int(18)        default 0    not null,
    MONTH_SESSION_3     int(18)        default 0    not null,
    MONTH_SESSION_4     int(18)        default 0    not null,
    MONTH_SESSION_5     int(18)        default 0    not null,
    MONTH_SESSION_6     int(18)        default 0    not null,
    MONTH_SESSION_7     int(18)        default 0    not null,
    MONTH_SESSION_8     int(18)        default 0    not null,
    MONTH_SESSION_9     int(18)        default 0    not null,
    MONTH_SESSION_10    int(18)        default 0    not null,
    MONTH_SESSION_11    int(18)        default 0    not null,
    MONTH_SESSION_12    int(18)        default 0    not null,
    MONTH_HIT_1         int(18)        default 0    not null,
    MONTH_HIT_2         int(18)        default 0    not null,
    MONTH_HIT_3         int(18)        default 0    not null,
    MONTH_HIT_4         int(18)        default 0    not null,
    MONTH_HIT_5         int(18)        default 0    not null,
    MONTH_HIT_6         int(18)        default 0    not null,
    MONTH_HIT_7         int(18)        default 0    not null,
    MONTH_HIT_8         int(18)        default 0    not null,
    MONTH_HIT_9         int(18)        default 0    not null,
    MONTH_HIT_10        int(18)        default 0    not null,
    MONTH_HIT_11        int(18)        default 0    not null,
    MONTH_HIT_12        int(18)        default 0    not null,
    MONTH_EVENT_1       int(18)        default 0    not null,
    MONTH_EVENT_2       int(18)        default 0    not null,
    MONTH_EVENT_3       int(18)        default 0    not null,
    MONTH_EVENT_4       int(18)        default 0    not null,
    MONTH_EVENT_5       int(18)        default 0    not null,
    MONTH_EVENT_6       int(18)        default 0    not null,
    MONTH_EVENT_7       int(18)        default 0    not null,
    MONTH_EVENT_8       int(18)        default 0    not null,
    MONTH_EVENT_9       int(18)        default 0    not null,
    MONTH_EVENT_10      int(18)        default 0    not null,
    MONTH_EVENT_11      int(18)        default 0    not null,
    MONTH_EVENT_12      int(18)        default 0    not null,
    MONTH_FAVORITE_1    int(18)        default 0    not null,
    MONTH_FAVORITE_2    int(18)        default 0    not null,
    MONTH_FAVORITE_3    int(18)        default 0    not null,
    MONTH_FAVORITE_4    int(18)        default 0    not null,
    MONTH_FAVORITE_5    int(18)        default 0    not null,
    MONTH_FAVORITE_6    int(18)        default 0    not null,
    MONTH_FAVORITE_7    int(18)        default 0    not null,
    MONTH_FAVORITE_8    int(18)        default 0    not null,
    MONTH_FAVORITE_9    int(18)        default 0    not null,
    MONTH_FAVORITE_10   int(18)        default 0    not null,
    MONTH_FAVORITE_11   int(18)        default 0    not null,
    MONTH_FAVORITE_12   int(18)        default 0    not null,
    constraint          IX_DATE_STAT unique (DATE_STAT)
);


---------------------------- Event --------------------------------

create table b_stat_event
(
    ID                int(18) auto_increment
        primary key,
    EVENT1            varchar(166) null,
    EVENT2            varchar(166) null,
    MONEY             decimal(18, 4) default 0.0000 not null,
    DATE_ENTER        datetime null,
    DATE_CLEANUP      datetime null,
    C_SORT            int(18)        default 100    null,
    COUNTER           int(18)        default 0      not null,
    ADV_VISIBLE       char           default 'Y' not null,
    NAME              varchar(50) null,
    DESCRIPTION       text null,
    KEEP_DAYS         int(18)                       null,
    DYNAMIC_KEEP_DAYS int(18)                       null,
    DIAGRAM_DEFAULT   char           default 'Y' not null
);

create
index IX_B_STAT_EVENT_2
    on b_stat_event (KEEP_DAYS);

create
index IX_EVENT1_EVENT2
    on b_stat_event (EVENT1, EVENT2);


create table event_day
(
    ID        int(18) auto_increment
        primary key,
    DATE_STAT date null,
    DATE_LAST datetime null,
    EVENT_ID  int(18)        default 0      not null,
    MONEY     decimal(18, 4) default 0.0000 not null,
    COUNTER   int(18)        default 0      not null
);

create
index IX_EVENT_ID_DATE_STAT
    on b_stat_event_day (EVENT_ID, DATE_STAT);

create table event_list
(
    ID              int auto_increment
        primary key,
    EVENT_ID        int            default 0 not null,
    EVENT3          varchar(255) null,
    MONEY           decimal(18, 4) default 0.0000 not null,
    DATE_ENTER      datetime not null,
    REFERER_URL     text null,
    URL             text null,
    REDIRECT_URL    text null,
    SESSION_ID      int null,
    GUEST_ID        int null,
    ADV_ID          int null,
    ADV_BACK        char           default 'N' not null,
    HIT_ID          int null,
    COUNTRY_ID      char(2) null,
    KEEP_DAYS       int null,
    CHARGEBACK      char           default 'N' not null,
    SITE_ID         char(2) null,
    REFERER_SITE_ID char(2) null
);

create
index IX_B_STAT_EVENT_LIST_2
    on b_stat_event_list (EVENT_ID, DATE_ENTER);

create
index IX_B_STAT_EVENT_LIST_3
    on b_stat_event_list (KEEP_DAYS, DATE_ENTER);

create
index IX_GUEST_ID
    on b_stat_event_list (GUEST_ID);

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
    engine = MergeTree PARTITION BY toYYYYMM(timestamp)
        ORDER BY timestamp
        SETTINGS index_granularity = 8192;


----------------------- Hit ---------------------------
create table hit
(
    ID           int auto_increment
        primary key,
    SESSION_ID   int  default 0 not null,
    DATE_HIT     datetime null,
    GUEST_ID     int null,
    NEW_GUEST    char default 'N' not null,
    USER_ID      int null,
    USER_AUTH    char null,
    URL          text null,
    URL_404      char default 'N' not null,
    URL_FROM     text null,
    IP           varchar(15) null,
    METHOD       varchar(10) null,
    COOKIES      text null,
    USER_AGENT   text null,
    STOP_LIST_ID int null,
    COUNTRY_ID   char(2) null,
    CITY_ID      int null,
    SITE_ID      char(2) null
);

create
index IX_DATE_HIT
    on b_stat_hit (DATE_HIT);

------------------ Page ----------------------

create table page
(
    ID            int auto_increment
        primary key,
    DATE_STAT     date not null,
    DIR           char default 'N' not null,
    URL           text not null,
    URL_404       char default 'N' not null,
    URL_HASH      int null,
    SITE_ID       char(2) null,
    COUNTER       int  default 0 not null,
    ENTER_COUNTER int  default 0 not null,
    EXIT_COUNTER  int  default 0 not null
);

create
index IX_DATE_STAT
    on b_stat_page (DATE_STAT);

create
index IX_URL_HASH
    on b_stat_page (URL_HASH);


---------------------- Path ------------------------

create table path
(
    ID                 int auto_increment
        primary key,
    PATH_ID            int  default 0 not null,
    PARENT_PATH_ID     int null,
    DATE_STAT          date null,
    COUNTER            int  default 0 not null,
    COUNTER_ABNORMAL   int  default 0 not null,
    COUNTER_FULL_PATH  int  default 0 not null,
    PAGES              text null,
    FIRST_PAGE         varchar(255) null,
    FIRST_PAGE_404     char default 'N' not null,
    FIRST_PAGE_SITE_ID char(2) null,
    PREV_PAGE          varchar(255) null,
    PREV_PAGE_HASH     int null,
    LAST_PAGE          varchar(255) null,
    LAST_PAGE_404      char default 'N' not null,
    LAST_PAGE_SITE_ID  char(2) null,
    LAST_PAGE_HASH     int null,
    STEPS              int  default 1 not null
);

create
index IX_DATE_STAT
    on b_stat_path (DATE_STAT);

create
index IX_PATH_ID_DATE_STAT
    on b_stat_path (PATH_ID, DATE_STAT);

create
index IX_PREV_PAGE_HASH_LAST_PAGE_HASH
    on b_stat_path (PREV_PAGE_HASH, LAST_PAGE_HASH);

create table path_adv
(
    ID                     int auto_increment
        primary key,
    ADV_ID                 int default 0 not null,
    PATH_ID                int default 0 not null,
    DATE_STAT              date null,
    COUNTER                int default 0 not null,
    COUNTER_BACK           int default 0 not null,
    COUNTER_FULL_PATH      int default 0 not null,
    COUNTER_FULL_PATH_BACK int default 0 not null,
    STEPS                  int default 0 not null
);

create
index IX_DATE_STAT
    on b_stat_path_adv (DATE_STAT);

create
index IX_PATH_ID_ADV_ID_DATE_STAT
    on b_stat_path_adv (PATH_ID, ADV_ID, DATE_STAT);


create table path_cache
(
    ID                      int auto_increment
        primary key,
    SESSION_ID              int  default 0 not null,
    DATE_HIT                datetime null,
    PATH_ID                 int null,
    PATH_PAGES              text null,
    PATH_FIRST_PAGE         varchar(255) null,
    PATH_FIRST_PAGE_404     char default 'N' not null,
    PATH_FIRST_PAGE_SITE_ID char(2) null,
    PATH_LAST_PAGE          varchar(255) null,
    PATH_LAST_PAGE_404      char default 'N' not null,
    PATH_LAST_PAGE_SITE_ID  char(2) null,
    PATH_STEPS              int  default 1 not null,
    IS_LAST_PAGE            char default 'Y' not null
);

create
index IX_SESSION_ID
    on b_stat_path_cache (SESSION_ID);

----------------------- Phrase ----------------------------

create table phrase_list
(
    ID          int auto_increment
        primary key,
    DATE_HIT    datetime null,
    SEARCHER_ID int null,
    REFERER_ID  int null,
    PHRASE      varchar(255) not null,
    URL_FROM    text null,
    URL_TO      text null,
    URL_TO_404  char default 'N' not null,
    SESSION_ID  int null,
    SITE_ID     char(2) null
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
    DATE_FIRST datetime null,
    DATE_LAST  datetime not null,
    SITE_NAME  varchar(255) not null,
    SESSIONS   int default 0 not null,
    HITS       int default 0 not null
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
    REFERER_ID int null,
    DATE_HIT   datetime null,
    PROTOCOL   varchar(10) not null,
    SITE_NAME  varchar(255) not null,
    URL_FROM   text not null,
    URL_TO     text null,
    URL_TO_404 char default 'N' not null,
    SESSION_ID int null,
    ADV_ID     int null,
    SITE_ID    char(2) null
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
    `id`                int(18)      not null auto_increment,
    `date_cleanup`      datetime,
    `total_hits`        int(18)      not null default '0',
    `save_statistic`    char(1) not null default 'Y',
    `active`            char(1) not null default 'Y',
    `name`              varchar(255) not null,
    `user_agent`        text,
    `diagram_default`   char(1) not null default 'N',
    `hit_keep_days`     int(18),
    `dynamic_keep_days` int(18),
    `phrases`           int(18)      not null default '0',
    `phrases_hits`      int(18)      not null default '0',
    `check_activity`    char(1) not null default 'Y',
    primary             key (`id`)
);
CREATE
INDEX IX_SEARCHER_1 ON searcher (`hit_keep_days`);

create table if not exists searcher_day
(
    `id`          int(18) not null auto_increment,
    `date_stat`   date,
    `date_last`   datetime,
    `searcher_id` int(18) not null default '0',
    `total_hits`  int(18) not null default '0',
    primary       key (`id`),
    index IX_SEARCHER_ID_DATE_STAT (`searcher_id`, `date_stat`)
);

create table if not exists searcher_hit
(
    `uuid`          int(18) not null auto_increment,
    `date_hit`      datetime,
    `searcher_id`   int(18) not null default '0',
    `url`           text not null,
    `url_404`       char(1) not null default 'N',
    `ip`            varchar(15),
    `user_agent`    text,
    `hit_keep_days` int(18),
    `site_id`       char(2),
    primary         key (`id`)
);
CREATE
INDEX IX_SEARCHER_HIT_1 ON searcher_hit (`searcher_id`, `date_hit`);
CREATE
INDEX IX_SEARCHER_HIT_2 ON searcher_hit (`hit_keep_days`, `date_hit`);

create table searcher_params
(
    ID          int auto_increment
        primary key,
    SEARCHER_ID int default 0 not null,
    DOMAIN      varchar(255) null,
    VARIABLE    varchar(255) null,
    CHAR_SET    varchar(255) null
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
--     COUNTRY_ID      char(2) null,
    FIRST_SITE_ID String,
    LAST_SITE_ID  String
--     CITY_ID         int(18)             null
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (uuid, date_stat)
      PRIMARY KEY (uuid, date_stat);



