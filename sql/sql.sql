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

-- create table if not exists session
-- (
--     uuid         UUID,
--     guestUuid    UUID,
--     phpSessionId String,
--     dateAdd      DateTime32('Europe/Moscow')
-- ) ENGINE = MergeTree
--       PARTITION BY toMonth(dateAdd)
--       ORDER BY (dateAdd);

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

create table if not exists searcher_total_hits
(
    dateStat     Date,
    searcherUuid UUID,
    totalHits    UInt64
) engine = SummingMergeTree(totalHits)
      ORDER BY (dateStat, searcherUuid);



-- SELECT t_adv.uuid,
--        t_adv.referer1,
--        t_adv.referer2,
--        t_adv.priority,
--        t_adv.events_view,
--        t_adv.description,
-- --        A.DATE_FIRST                                                                            C_TIME_FIRST,
-- --        A.DATE_LAST                                                                             C_TIME_LAST,
--        'RUB'                                                                                           CURRENCY,
-- --        DATE_FORMAT(A.DATE_FIRST, '%d.%m.%Y')                                                   DATE_FIRST,
-- --        DATE_FORMAT(A.DATE_LAST, '%d.%m.%Y')                                                    DATE_LAST,
--        toUnixTimestamp(ifNull(A.DATE_LAST, 0)) - toUnixTimestamp(ifNull(A.DATE_FIRST, 0))              ADV_TIME,
--
--        -- TODAY
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.guests_day, 0), 0))  GUESTS_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.new_guests, 0), 0))  NEW_GUESTS_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.favorites, 0), 0))   FAVORITES_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.hosts_day, 0), 0))   C_HOSTS_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.sessions, 0), 0))    SESSIONS_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.hits, 0), 0))        HITS_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.guests_day_back, 0),
--               0))                                                                                      GUESTS_BACK_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.favorites_back, 0),
--               0))                                                                                      FAVORITES_BACK_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.hosts_day_back, 0),
--               0))                                                                                      HOSTS_BACK_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.sessions_back, 0),
--               0))                                                                                      SESSIONS_BACK_TODAY,
--        sum(if(TO_DAYS(curdate()) = TO_DAYS(t_adv_day.date_stat), ifNull(t_adv_day.hits_back, 0), 0))   HITS_BACK_TODAY,
--
--        -- YESTERDAY
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.guests_day, 0),
--               0))                                                                                      GUESTS_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.new_guests, 0),
--               0))                                                                                      NEW_GUESTS_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.favorites, 0),
--               0))                                                                                      FAVORITES_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.hosts_day, 0),
--               0))                                                                                      C_HOSTS_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.sessions, 0),
--               0))                                                                                      SESSIONS_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.hits, 0), 0))    HITS_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.guests_day_back, 0),
--               0))                                                                                      GUESTS_BACK_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.favorites_back, 0),
--               0))                                                                                      FAVORITES_BACK_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.hosts_day_back, 0),
--               0))                                                                                      HOSTS_BACK_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.sessions_back, 0),
--               0))                                                                                      SESSIONS_BACK_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 1, ifNull(t_adv_day.hits_back, 0),
--               0))                                                                                      HITS_BACK_YESTERDAY,
--
--        -- THE DAY BEFORE YESTERDAY
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.guests_day, 0),
--               0))                                                                                      GUESTS_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.new_guests, 0),
--               0))                                                                                      NEW_GUESTS_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.favorites, 0),
--               0))                                                                                      FAVORITES_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.hosts_day, 0),
--               0))                                                                                      C_HOSTS_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.sessions, 0),
--               0))                                                                                      SESSIONS_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.hits, 0),
--               0))                                                                                      HITS_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.guests_day_back, 0),
--               0))                                                                                      GUESTS_BACK_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.favorites_back, 0),
--               0))                                                                                      FAVORITES_BACK_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.hosts_day_back, 0),
--               0))                                                                                      HOSTS_BACK_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.sessions_back, 0),
--               0))                                                                                      SESSIONS_BACK_BEF_YESTERDAY,
--        sum(if(TO_DAYS(curdate()) - TO_DAYS(t_adv_day.date_stat) = 2, ifNull(t_adv_day.hits_back, 0),
--               0))                                                                                      HITS_BACK_BEF_YESTERDAY,
--
--
--        -- PERIOD
--        t_adv_stat.guests                                                                               GUESTS_PERIOD,
--        t_adv_stat.hosts                                                                                C_HOSTS_PERIOD,
--        t_adv_stat.new_guests                                                                           NEW_GUESTS_PERIOD,
--        t_adv_stat.favorites                                                                            FAVORITES_PERIOD,
--        t_adv_stat.sessions                                                                             SESSIONS_PERIOD,
--        t_adv_stat.hits                                                                                 HITS_PERIOD,
--        t_adv_stat.guests_back                                                                          GUESTS_BACK_PERIOD,
--        t_adv_stat.hosts_back                                                                           HOSTS_BACK_PERIOD,
--        t_adv_stat.favorites                                                                            FAVORITES_BACK_PERIOD,
--        t_adv_stat.sessions_back                                                                        SESSIONS_BACK_PERIOD,
--        t_adv_stat.hits_back                                                                            HITS_BACK_PERIOD,
--
--        -- TOTAL
--        t_adv_stat.favorites_back,
--
--        -- AUDIENCE
--        if(t_adv_stat.sessions > 0, round(t_adv_stat.hits / t_adv_stat.sessions, 2), -1)                ATTENT,
--        if(t_adv_stat.sessions_back > 0, round(t_adv_stat.hits_back / t_adv_stat.sessions_back, 2), -1) ATTENT_BACK,
--        if(t_adv_stat.guests > 0, round((t_adv_stat.new_guests / t_adv_stat.guests) * 100, 2), -1)      NEW_VISITORS,
--        if(t_adv_stat.guests > 0, round((t_adv_stat.guests_back / t_adv_stat.guests) * 100, 2),
--           -1)                                                                                          RETURNED_VISITORS,
--        if(
--                round((((toUnixTimestamp(ifNull(A.DATE_LAST, 0)) - toUnixTimestamp(ifNull(A.DATE_FIRST, 0))) / 86400)),
--                      0) >= 1, round(A.GUESTS / ((toUnixTimestamp(ifNull(A.DATE_LAST, 0)) -
--                                                  toUnixTimestamp(ifNull(A.DATE_FIRST, 0))) / 86400), 2),
--                -1)                                                                                     VISITORS_PER_DAY,
--
--        -- FINANCES
--        round(round(t_adv.cost, 2) * 1, 2)                                                              COST,
--        round(round(t_adv_stat.revenue, 2) * 1, 2)                                                      REVENUE,
--        round(round(t_adv_stat.revenue - t_adv.cost, 2) * 1, 2)                                         BENEFIT,
--        round(round(if(t_adv_stat.sessions > 0, t_adv.cost / t_adv_stat.sessions, 0), 2) * 1, 2)        SESSION_COST,
--        round(round(if(t_adv_stat.guests > 0, t_adv.cost / t_adv_stat.guests, 0), 2) * 1, 2)            VISITOR_COST,
--        if(t_adv.cost > 0, round(((t_adv_stat.revenue - t_adv.cost) / t_adv.cost) * 100, 2), -1)        ROI
--
-- FROM adv t_adv
--          JOIN adv_stat t_adv_stat ON t_adv_stat.adv_uuid = t_adv.uuid
--          LEFT JOIN adv_day t_adv_day ON (t_adv_day.adv_uuid = t_adv_stat.adv_uuid)
-- GROUP BY t_adv.uuid, t_adv.referer1, t_adv.referer2, t_adv.cost, t_adv_stat.revenue, t_adv.priority, t_adv.events_view,
--          t_adv.description, A.date_first,
--          A.date_last, t_adv_stat.guests, t_adv_stat.new_guests, t_adv_stat.favorites, t_adv_stat.hosts,
--          t_adv_stat.sessions, t_adv_stat.hits,
--          t_adv_stat.guests_back,
--          t_adv_stat.favorites_back, t_adv_stat.hosts_back, t_adv_stat.sessions_back, t_adv_stat.hits_back
-- -- ORDER BY sessions_today desc, sessions_yesterday desc, sessions_bef_yesterday desc, sessions_period desc, sessions desc
-- LIMIT 500


