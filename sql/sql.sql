---------------------- ADV -------------------------

create table if not exists adv
(
    `uuid`           UUID,
    `referer1`       String,
    `referer2`       String,
    `cost`           decimal(18, 4) default 0.0000,
    `revenue`        decimal(18, 4) default 0.0000,
    `events_view`    String,
    `guests`         Int32          default 0,
    `new_guests`     Int32          default 0,
    `favorites`      Int32          default 0,
    `hosts`          Int32          default 0,
    `sessions`       Int32          default 0,
    `hits`           Int32          default 0,
    `date_first`     DateTime32('Europe/Moscow'),
    `date_last`      DateTime32('Europe/Moscow'),
    `guests_back`    Int32          default 0,
    `favorites_back` Int32          default 0,
    `hosts_back`     Int32          default 0,
    `sessions_back`  Int32          default 0,
    `hits_back`      Int32          default 0,
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
    `guests`          Int32 default 0,
    `guests_day`      Int32 default 0,
    `new_guests`      Int32 default 0,
    `favorites`       Int32 default 0,
    `hosts`           Int32 default 0,
    `hosts_day`       Int32 default 0,
    `sessions`        Int32 default 0,
    `hits`            Int32 default 0,
    `guests_back`     Int32 default 0,
    `guests_day_back` Int32 default 0,
    `favorites_back`  Int32 default 0,
    `hosts_back`      Int32 default 0,
    `hosts_day_back`  Int32 default 0,
    `sessions_back`   Int32 default 0,
    `hits_back`       Int32 default 0
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (`adv_uuid`, `date_stat`);

create table adv_event
(
    ID           int(18) auto_increment
        primary key,
    ADV_ID       int(18)        default 0      null,
    EVENT_ID     int(18)        default 0      null,
    COUNTER      int(18)        default 0      not null,
    COUNTER_BACK int(18)        default 0      not null,
    MONEY        decimal(18, 4) default 0.0000 not null,
    MONEY_BACK   decimal(18, 4) default 0.0000 not null
);

create table adv_event_day
(
    ID           int(18) auto_increment
        primary key,
    ADV_ID       int(18)        default 0      null,
    EVENT_ID     int(18)        default 0      null,
    DATE_STAT    date                          null,
    COUNTER      int(18)        default 0      not null,
    COUNTER_BACK int(18)        default 0      not null,
    MONEY        decimal(18, 4) default 0.0000 not null,
    MONEY_BACK   decimal(18, 4) default 0.0000 not null
);

create index IX_ADV_ID_EVENT_ID_DATE_STAT
    on b_stat_adv_event_day (ADV_ID, EVENT_ID, DATE_STAT);

create index IX_DATE_STAT
    on b_stat_adv_event_day (DATE_STAT);

create index IX_ADV_EVENT_ID
    on b_stat_adv_event (ADV_ID, EVENT_ID);

create table adv_guest
(
    ID             int auto_increment
        primary key,
    ADV_ID         int  default 0   not null,
    BACK           char default 'N' not null,
    GUEST_ID       int  default 0   not null,
    DATE_GUEST_HIT datetime         null,
    DATE_HOST_HIT  datetime         null,
    SESSION_ID     int  default 0   not null,
    IP             varchar(15)      null,
    IP_NUMBER      bigint           null
);

create index IX_ADV_ID_GUEST
    on b_stat_adv_guest (ADV_ID, GUEST_ID);

create index IX_ADV_ID_IP_NUMBER
    on b_stat_adv_guest (ADV_ID, IP_NUMBER);


create table if not exists adv_page
(
    `uuid`     UUID,
    `adv_uuid` UUID,
    `page`     String,
    `type`     String default 'TO'
) ENGINE = MergeTree
      ORDER BY (`adv_uuid`, `type`);

create index IX_ADV_ID_TYPE
    on b_stat_adv_page (ADV_ID, C_TYPE);


create table adv_searcher
(
    ID          int(18) auto_increment
        primary key,
    ADV_ID      int(18) not null,
    SEARCHER_ID int(18) not null
);

create index idx_search_adv
    on b_stat_adv_searcher (SEARCHER_ID, ADV_ID);

----------------------- Browser --------------------------
create table browser
(
    ID         int(18) auto_increment
        primary key,
    USER_AGENT varchar(255) not null
);

------------------------ City -----------------------------

create table if not exists city
(
    `uuid`       UUID,
    `country_id` String,
    `region`     String,
    `name`       String,
    `xml_id`     String,
    `sessions`   Int32 default 0,
    `new_guests` Int32 default 0,
    `hits`       Int32 default 0,
    `events`     Int32 default 0
) engine = MergeTree
      ORDER BY (country_id, region, name)
      SETTINGS index_granularity = 8192;

create table b_stat_city_day
(
    ID         int(18) auto_increment
        primary key,
    CITY_ID    int(18)           not null,
    DATE_STAT  date              not null,
    SESSIONS   int(18) default 0 not null,
    NEW_GUESTS int(18) default 0 not null,
    HITS       int(18) default 0 not null,
    C_EVENTS   int(18) default 0 not null
);

create index IX_B_STAT_CITY_DAY_1
    on b_stat_city_day (CITY_ID, DATE_STAT);

create index IX_B_STAT_CITY_DAY_2
    on b_stat_city_day (DATE_STAT);


create table b_stat_city_ip
(
    START_IP   bigint(18) not null
        primary key,
    END_IP     bigint(18) not null,
    COUNTRY_ID char(2)    not null,
    CITY_ID    int(18)    not null
);

create index IX_B_STAT_CITY_IP_END_IP
    on b_stat_city_ip (END_IP);

create table b_stat_city_ip
(
    START_IP   bigint(18) not null
        primary key,
    END_IP     bigint(18) not null,
    COUNTRY_ID char(2)    not null,
    CITY_ID    int(18)    not null
);

create index IX_B_STAT_CITY_IP_END_IP
    on b_stat_city_ip (END_IP);

------------------ Country ---------------------

create table b_stat_country
(
    ID         char(2)           not null
        primary key,
    SHORT_NAME char(3)           null,
    NAME       varchar(50)       null,
    SESSIONS   int(18) default 0 not null,
    NEW_GUESTS int(18) default 0 not null,
    HITS       int(18) default 0 not null,
    C_EVENTS   int(18) default 0 not null
);

create table b_stat_country_day
(
    ID         int(18) auto_increment
        primary key,
    COUNTRY_ID char(2)           not null,
    DATE_STAT  date              null,
    SESSIONS   int(18) default 0 not null,
    NEW_GUESTS int(18) default 0 not null,
    HITS       int(18) default 0 not null,
    C_EVENTS   int(18) default 0 not null
);

create index IX_COUNTRY_ID_DATE_STAT
    on b_stat_country_day (COUNTRY_ID, DATE_STAT);

---------------------day---------------------

create table b_stat_day
(
    ID                  int(18) auto_increment
        primary key,
    DATE_STAT           date                        null,
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
    constraint IX_DATE_STAT
        unique (DATE_STAT)
);


---------------------------- Event --------------------------------

create table b_stat_event
(
    ID                int(18) auto_increment
        primary key,
    EVENT1            varchar(166)                  null,
    EVENT2            varchar(166)                  null,
    MONEY             decimal(18, 4) default 0.0000 not null,
    DATE_ENTER        datetime                      null,
    DATE_CLEANUP      datetime                      null,
    C_SORT            int(18)        default 100    null,
    COUNTER           int(18)        default 0      not null,
    ADV_VISIBLE       char           default 'Y'    not null,
    NAME              varchar(50)                   null,
    DESCRIPTION       text                          null,
    KEEP_DAYS         int(18)                       null,
    DYNAMIC_KEEP_DAYS int(18)                       null,
    DIAGRAM_DEFAULT   char           default 'Y'    not null
);

create index IX_B_STAT_EVENT_2
    on b_stat_event (KEEP_DAYS);

create index IX_EVENT1_EVENT2
    on b_stat_event (EVENT1, EVENT2);


create table b_stat_event_day
(
    ID        int(18) auto_increment
        primary key,
    DATE_STAT date                          null,
    DATE_LAST datetime                      null,
    EVENT_ID  int(18)        default 0      not null,
    MONEY     decimal(18, 4) default 0.0000 not null,
    COUNTER   int(18)        default 0      not null
);

create index IX_EVENT_ID_DATE_STAT
    on b_stat_event_day (EVENT_ID, DATE_STAT);


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
    `uuid`            int(18) not null auto_increment,
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

create table guest
(
    uuid       UUID,
    timestamp  DateTime('Europe/Moscow'),
    favorites  UInt8 default 0,
    events     Int32 default 0,
    sessions   Int32 default 0,
    hits       Int32 default 0,
    repair     UInt8 default 0,
    session_id UUID,
    url_from   String,
    url_to     String,
    url_to_404 UInt8 default 0,
    site_id    String,
    adv_id     UUID,
    referer1   String,
    referer2   String,
    referer3   String,
    user_id    Int32,
    user_auth  UInt8,
    url        String,
    url_404    UInt8 default 0,
    user_agent String,
    ip         IPv4,
    cookie     String,
    language   String,
    adv_back   UInt8 default 0,
    token      String
)
    engine = MergeTree PARTITION BY toYYYYMM(timestamp)
        ORDER BY timestamp
        SETTINGS index_granularity = 8192;

