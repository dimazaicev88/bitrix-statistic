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
    `id`            int(18) not null auto_increment,
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


create table if not exists adv_page
(
    `uuid`     UUID,
    `adv_uuid` UUID,
    `page`     String,
    `type`     String default 'TO'
) ENGINE = MergeTree
      ORDER BY (`adv_uuid`, `type`);


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