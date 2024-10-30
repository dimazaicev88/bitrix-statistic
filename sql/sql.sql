use default;
DROP DATABASE if exists statistic;
create database statistic;
USE statistic;

create table hit_full
(
    uuid             UUID,
    dateHit          DateTime('Europe/Moscow'),
    cookies          String,
    countryId        FixedString(2),
    city             String,
    event1           String,
    event2           String,
    guestHash        UUID,
    ip               IPv4,
    method           String,
    phpSessionId     String,
    referrerSiteName String,
    referrerUrlFrom  String,
    siteId           FixedString(2),
    url              String,
    url404           Bool DEFAULT false,
    urlFrom          String,
    userAgent        String,
    userAuth         Bool DEFAULT false,
    userId           UInt32

)
    engine = MergeTree
        PARTITION BY toMonth(dateHit)
        ORDER BY dateHit;

create table if not exists raw_request
(
    date              DateTime32('Europe/Moscow'),
    phpSessionId      String,
    url               String,
    referer           String,
    ip                IPv4,
    userAgent         String,
    userid            UInt32,
    userLogin         String,
    httpXForwardedFor String,
    isError404        bool,
    siteId            String,
    event1            String,
    event2            String,
    isUserAuth        bool
) ENGINE = MergeTree
      PARTITION BY toMonth(date)
      ORDER BY (date);