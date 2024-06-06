create table if not exists session
(
    `id`              int(18) not null auto_increment,
    `guest_id`        int(18),
    `new_guest`       char(1) not null default 'N',
    `user_id`         int(18),
    `user_auth`       char(1),
    `c_events`        int(18) not null default '0',
    `hits`            int(18) not null default '0',
    `favorites`       char(1) not null default 'N',
    `url_from`        text,
    `url_to`          text,
    `url_to_404`      char(1) not null default 'N',
    `url_last`        text,
    `url_last_404`    char(1) not null default 'N',
    `user_agent`      text,
    `date_stat`       date,
    `date_first`      datetime,
    `date_last`       datetime,
    `ip_first`        varchar(15),
    `ip_first_number` bigint(20),
    `ip_last`         varchar(15),
    `ip_last_number`  bigint(20),
    `first_hit_id`    int(18),
    `last_hit_id`     int(18),
    `phpsessid`       varchar(255),
    `adv_id`          int(18),
    `adv_back`        char(1),
    `referer1`        varchar(255),
    `referer2`        varchar(255),
    `referer3`        varchar(255),
    `stop_list_id`    int(18),
    `country_id`      char(2),
    `city_id`         int(18),
    `first_site_id`   char(2),
    `last_site_id`    char(2),
    primary key (`id`),
    index IX_IP_FIRST_NUMBER_DATE_STAT (`ip_first_number`, `date_stat`),
    index IX_SESSION_4 (`user_id`, `date_stat`),
    index IX_DATE_STAT (`date_stat`),
    INDEX IX_SESSION_5 (`date_last`),
    INDEX IX_SESSION_6 (`guest_id`)
);

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