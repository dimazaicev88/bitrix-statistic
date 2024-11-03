use default;
DROP DATABASE if exists statistic;
create database statistic;
USE statistic;

create table if not exists guest
(
    guestHash String(32),
    dateAdd   DateTime32('Europe/Moscow')
)
    engine = MergeTree()
        PARTITION BY toMonth(dateAdd)
        ORDER BY dateAdd;


----------------------- Hit ---------------------------
-- create table if not exists hit
-- (
--     uuid         UUID,
--     sessionUuid  UUID,
--     advUuid      UUID,
--     dateHit      DateTime32('Europe/Moscow'),
--     phpSessionId String,
--     guestUuid    UUID,
--     language     String,
--     isNewGuest   BOOLEAN default false,
--     userId       UInt32,
--     userAuth     BOOLEAN default false,
--     url          String,
--     url404       BOOLEAN default false,
--     urlFrom      String,
--     ip           IPv4,
--     method       String,
--     cookies      String,
--     userAgent    String,
--     stopListUuid UUID,
--     countryId    FixedString(2),
--     cityUuid     UUID,
--     siteId       FixedString(2),
--     favorites    boolean default false
-- )
--     engine = MergeTree
--         PARTITION BY toMonth(dateHit)
--         ORDER BY dateHit;

create table hit
(
    uuid         UUID,
    dateHit      DateTime('Europe/Moscow'),
    cookies      String,
    countryId    FixedString(2),
    city         String,
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
    userAuth     Bool DEFAULT false,
    userId       UInt32
)
    engine = MergeTree
        PARTITION BY toMonth(dateHit)
        ORDER BY dateHit;

create table if not exists raw_request
(
    date    DateTime32('Europe/Moscow'),
    phpSessionId      String,
    url     String,
    referer String,
    ip      IPv4,
    userAgent         String,
    userid  UInt32,
    httpXForwardedFor String,
    isError404        bool,
    siteId            String,
    event1  String,
    event2  String,
    isUserAuth        bool
) ENGINE = MergeTree
      PARTITION BY toMonth(date)
      ORDER BY (date);