use default;
DROP DATABASE if exists statistic;
create database statistic;
USE statistic;

create table if not exists guests
(
    guestHash  String(32),
    dateInsert DateTime32('Europe/Moscow')
)
    engine = MergeTree()
        PARTITION BY toMonth(dateInsert)
        ORDER BY dateInsert;

create table hits
(
    uuid         UUID,
    dateHit      DateTime('Europe/Moscow'),
    cookies      String,
    event1       String,
    event2       String,
    guestHash    FixedString(32),
    isNewGuest   Bool DEFAULT false,
    ip           IPv4,
    method       String,
    phpSessionId String,
    referrer     String,
    siteId       FixedString(2),
    url          String,
    url404       Bool DEFAULT false,
    urlFrom      String,
    userAgent    String,
    userId       UInt32
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
    httpXForwardedFor String,
    isError404        bool,
    siteId            String,
    event1            String,
    event2            String,
    isUserAuth        bool
) ENGINE = MergeTree
      PARTITION BY toMonth(date)
      ORDER BY (date);