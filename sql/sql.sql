create table if not exists session_data
(
    `id`              int(18)      not null auto_increment,
    `date_first`      datetime     null,
    `date_last`       datetime     null,
    `guest_md5`       varchar(255) not null,
    `sess_session_id` int(18)      null,
    `session_data`    text         null,
    primary key (`id`),
    index IX_GUEST_MD5 (`guest_md5`)
);
create table if not exists searcher
(
    `id`                int(18)      not null auto_increment,
    `date_cleanup`      datetime,
    `total_hits`        int(18)      not null default '0',
    `save_statistic`    char(1)      not null default 'Y',
    `active`            char(1)      not null default 'Y',
    `name`              varchar(255) not null,
    `user_agent`        text,
    `diagram_default`   char(1)      not null default 'N',
    `hit_keep_days`     int(18),
    `dynamic_keep_days` int(18),
    `phrases`           int(18)      not null default '0',
    `phrases_hits`      int(18)      not null default '0',
    `check_activity`    char(1)      not null default 'Y',
    primary key (`id`)
);
CREATE INDEX IX_SEARCHER_1 ON searcher (`hit_keep_days`);

create table if not exists searcher_day
(
    `id`          int(18) not null auto_increment,
    `date_stat`   date,
    `date_last`   datetime,
    `searcher_id` int(18) not null default '0',
    `total_hits`  int(18) not null default '0',
    primary key (`id`),
    index IX_SEARCHER_ID_DATE_STAT (`searcher_id`, `date_stat`)
);

create table if not exists searcher_hit
(
    `id`            int(18) not null auto_increment,
    `date_hit`      datetime,
    `searcher_id`   int(18) not null default '0',
    `url`           text    not null,
    `url_404`       char(1) not null default 'N',
    `ip`            varchar(15),
    `user_agent`    text,
    `hit_keep_days` int(18),
    `site_id`       char(2),
    primary key (`id`)
);
CREATE INDEX IX_SEARCHER_HIT_1 ON searcher_hit (`searcher_id`, `date_hit`);
CREATE INDEX IX_SEARCHER_HIT_2 ON searcher_hit (`hit_keep_days`, `date_hit`);


create table if not exists adv_day
(
    `id`              int(18) not null auto_increment,
    `adv_id`          int(18) not null default '0',
    `date_stat`       date,
    `guests`          int(18) not null default '0',
    `guests_day`      int(18) not null default '0',
    `new_guests`      int(18) not null default '0',
    `favorites`       int(18) not null default '0',
    `c_hosts`         int(18) not null default '0',
    `c_hosts_day`     int(18) not null default '0',
    `sessions`        int(18) not null default '0',
    `hits`            int(18) not null default '0',
    `guests_back`     int(18) not null default '0',
    `guests_day_back` int(18) not null default '0',
    `favorites_back`  int(18) not null default '0',
    `hosts_back`      int(18) not null default '0',
    `hosts_day_back`  int(18) not null default '0',
    `sessions_back`   int(18) not null default '0',
    `hits_back`       int(18) not null default '0',
    primary key (`id`),
    index IX_ADV_ID_DATE_STAT (`adv_id`, `date_stat`),
    index IX_DATE_STAT (`date_stat`)
);

create table if not exists adv_page
(
    `id`     int(18)      not null auto_increment,
    `adv_id` int(18)      not null default '0',
    `page`   varchar(255) not null,
    `c_type` varchar(5)   not null default 'TO',
    primary key (`id`),
    index IX_ADV_ID_TYPE (`adv_id`, `c_type`)
);


create table if not exists adv
(
    `id`             varchar(15)    not null not null,
    `referer1`       varchar(255),
    `referer2`       varchar(255),
    `cost`           decimal(18, 4) not null default '0.0000',
    `revenue`        decimal(18, 4) not null default '0.0000',
    `events_view`    varchar(255),
    `guests`         int(18)        not null default '0',
    `new_guests`     int(18)        not null default '0',
    `favorites`      int(18)        not null default '0',
    `c_hosts`        int(18)        not null default '0',
    `sessions`       int(18)        not null default '0',
    `hits`           int(18)        not null default '0',
    `date_first`     datetime,
    `date_last`      datetime,
    `guests_back`    int(18)        not null default '0',
    `favorites_back` int(18)        not null default '0',
    `hosts_back`     int(18)        not null default '0',
    `sessions_back`  int(18)        not null default '0',
    `hits_back`      int(18)        not null default '0',
    `description`    text,
    `priority`       int(18)        not null default '100',
    primary key (`id`),
    index IX_REFERER1 (`referer1`)
);

create table if not exists session
(
    `id`           varchar(15) not null,
    `guest_id`     int(18),
    `events`       int(18)     not null default '0',
    `hits`         int(18)     not null default '0',
    `timestamp`    timestamp,
    `phpsessid`    varchar(255),
    `stop_list_id` int(18),
    primary key (`id`),
    INDEX IX_GUEST_ID_SESSION (`guest_id`)
);

create table if not exists guest
(
    `id`          varchar(15) not null,
    `timestamp_x` timestamp,
    `favorites`   char(1)     not null default 0,
    `events`      int(18)     not null default 0,
    `sessions`    int(18)     not null default 0,
    `hits`        int(18)     not null default 0,
    `repair`      tinyint     not null default 0,
    `session_id`  varchar(15),
    `date`        datetime,
    `url_from`    text,
    `url_to`      text,
    `url_to_404`  tinyint     not null default 0,
    `site_id`     char(2),
    `adv_id`      int(18),
    `referer1`    varchar(255),
    `referer2`    varchar(255),
    `referer3`    varchar(255),
    `user_id`     int(18),
    `user_auth`   char(1),
    `url`         text,
    `url_404`     tinyint     not null default 0,
    `user_agent`  text,
    `ip`          varchar(15),
    `cookie`      text,
    `language`    varchar(255),
    `adv_back`    tinyint     not null default 0,
    `country_id`  char(2),
    `city_id`     int(18),
    `city_info`   text,
    `token`       varchar(15),
    primary key (`id`)
);

create table if not exists city
(
    `id`         int(18) not null auto_increment,
    `country_id` char(2) not null,
    `region`     varchar(200),
    `name`       varchar(255),
    `xml_id`     varchar(255),
    `sessions`   int(18) not null default '0',
    `new_guests` int(18) not null default '0',
    `hits`       int(18) not null default '0',
    `c_events`   int(18) not null default '0',
    primary key (`id`),
    index UX_CITY (`country_id`, `region`(50), `name`(50)),
    index IX_CITY_XML_ID (`xml_id`)
);