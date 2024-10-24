use default;
DROP DATABASE if exists statistic;
create database statistic;
USE statistic;
---------------------- ADV -------------------------

create table if not exists adv
(
    uuid        UUID,
    referer1    String,
    referer2    String,
    dateCreate DateTime32('Europe/Moscow'),
    cost        decimal(18, 4) default 0.0000,
    eventsView String,
    description String
) ENGINE = MergeTree
      ORDER BY (referer1);

create table if not exists adv_stat
(
    advUuid       UUID,
    guests        UInt32 default 0,
    newGuests     UInt32 default 0,
    favorites     UInt32 default 0,
    hosts         UInt32 default 0,
    sessions      UInt32 default 0,
    hits          UInt32 default 0,
    guestsBack    UInt32 default 0,
    favoritesBack UInt32 default 0,
    hostsBack     UInt32 default 0,
    sessionsBack  UInt32 default 0,
    hitsBack      UInt32 default 0
) engine = SummingMergeTree((
                             guests,
                             newGuests,
                             favorites,
                             hosts,
                             sessions,
                             hits,
                             guestsBack,
                             favoritesBack,
                             hostsBack,
                             sessionsBack,
                             hitsBack))
      ORDER BY (advUuid);


create table if not exists adv_day
(
    advUuid       UUID,
    dateStat      Date32,
    guests        UInt32 default 0,
    guestsDay     UInt32 default 0,
    newGuests     UInt32 default 0,
    favorites     UInt32 default 0,
    hosts         UInt32 default 0,
    hostsDay      UInt32 default 0,
    sessions      UInt32 default 0,
    hits          UInt32 default 0,
    guestsBack    UInt32 default 0,
    guestsDayBack UInt32 default 0,
    favoritesBack UInt32 default 0,
    hostsBack     UInt32 default 0,
    hostsDayBack  UInt32 default 0,
    sessionsBack  UInt32 default 0,
    hitsBack      UInt32 default 0
) ENGINE = SummingMergeTree(
        (guests,
         guestsDay,
         newGuests,
         favorites,
         hosts,
         hostsDay,
         sessions,
         hits,
         guestsBack,
         guestsDayBack,
         favoritesBack,
         hostsBack,
         hostsDayBack,
         sessionsBack,
         hitsBack)
           )
      PARTITION BY toMonth(dateStat)
      ORDER BY (advUuid, dateStat);

create table if not exists adv_event
(
    uuid        UUID,
    advUuid     UUID,
    eventUuid   UUID,
    counter     UInt32         default 0,
    counterBack UInt32         default 0,
    money       decimal(18, 4) default 0.0000,
    moneyBack   decimal(18, 4) default 0.0000
) ENGINE = MergeTree
      ORDER BY (money);

create table if not exists adv_event_day
(
    uuid        String,
    advUuid     String,
    eventUuid   String,
    dateStat    DateTime32('Europe/Moscow'),
    counter     UInt32,
    counterBack UInt32,
    money       decimal(18, 4) default 0.0000,
    moneyBack   decimal(18, 4) default 0.0000
) ENGINE = MergeTree
      ORDER BY (advUuid, eventUuid, dateStat);

create table if not exists adv_guest
(
    uuid         UUID,
    advUuid      UUID,
    back         BOOLEAN default false,
    guestUuid    UUID,
    dateGuestHit DateTime32('Europe/Moscow'),
    dateHostHit  DateTime32('Europe/Moscow'),
    sessionUuid  UUID,
    ip           IPv4
) ENGINE = MergeTree
      ORDER BY (advUuid, guestUuid);


create table if not exists adv_page
(
    uuid    UUID,
    advUuid UUID,
    page    String,
    type    String default 'TO'
) ENGINE = MergeTree
      ORDER BY (advUuid, type);

create table if not exists adv_searcher
(
    uuid         UUID,
    advUuid      UUID,
    searcherUuid UUID
) ENGINE = MergeTree
      ORDER BY (advUuid, searcherUuid);

----------------------- Browser --------------------------
create table if not exists browser
(
    uuid      UUID,
    userAgent String
) ENGINE = MergeTree
      ORDER BY (userAgent);

------------------------ City -----------------------------

create table if not exists city
(
    uuid      UUID,
    countryId String,
    region    String,
    name      String,
    xmlId     String,
    sessions  UInt32 default 0,
    newGuests UInt32 default 0,
    hits      UInt32 default 0,
    events    UInt32 default 0
) engine = MergeTree
      ORDER BY (countryId, region, name);

create table if not exists city_day
(
    cityUuid  UUID,
    dateStat  Date32,
    sessions  UInt32 default 0,
    newGuests UInt32 default 0,
    hits      UInt32 default 0,
    events    UInt32 default 0
) engine = SummingMergeTree((sessions, newGuests, hits, events))
      ORDER BY (cityUuid, dateStat);

create table if not exists city_ip
(
    startIp   UInt32,
    endIp     UInt32,
    countryId String,
    cityUuid  UUID
) engine = MergeTree
      ORDER BY (endIp);

------------------ Country ---------------------

create table if not exists country
(
    uuid      UUID,
    shortName String,
    name      String,
    sessions  UInt32 default 0,
    newGuests UInt32 default 0,
    hits      UInt32 default 0,
    events    UInt32 default 0
) engine = MergeTree
      ORDER BY (name);

create table if not exists country_day
(
    countryId FixedString(2),
    dateStat  Date32,
    sessions  UInt32 default 0,
    newGuests UInt32 default 0,
    hits      UInt32 default 0,
    events    UInt32 default 0
) engine = SummingMergeTree((sessions, newGuests, hits, events))
      ORDER BY (countryId, dateStat);

---------------------day---------------------

create table if not exists day
(
    uuid                UUID,
    date_stat           Date32,
    site_id             FixedString(2),
    hits                UInt32 default 0,
    hosts               UInt32 default 0,
    sessions            UInt32 default 0,
    events              UInt32 default 0,
    guests              UInt32 default 0,
    new_guests          UInt32 default 0,
    favorites           UInt32 default 0,
    total_hosts         UInt32 default 0,
    am_1                UInt32 default 0,
    am_1_3              UInt32 default 0,
    am_3_6              UInt32 default 0,
    am_6_9              UInt32 default 0,
    am_9_12             UInt32 default 0,
    am_12_15            UInt32 default 0,
    am_15_18            UInt32 default 0,
    am_18_21            UInt32 default 0,
    am_21_24            UInt32 default 0,
    am_24               UInt32 default 0,
    ah_1                UInt32 default 0,
    ah_2_5              UInt32 default 0,
    ah_6_9              UInt32 default 0,
    ah_10_13            UInt32 default 0,
    ah_14_17            UInt32 default 0,
    ah_18_21            UInt32 default 0,
    ah_22_25            UInt32 default 0,
    ah_26_29            UInt32 default 0,
    ah_30_33            UInt32 default 0,
    ah_34               UInt32 default 0,
    hour_host_0         UInt32 default 0,
    hour_host_1         UInt32 default 0,
    hour_host_2         UInt32 default 0,
    hour_host_3         UInt32 default 0,
    hour_host_4         UInt32 default 0,
    hour_host_5         UInt32 default 0,
    hour_host_6         UInt32 default 0,
    hour_host_7         UInt32 default 0,
    hour_host_8         UInt32 default 0,
    hour_host_9         UInt32 default 0,
    hour_host_10        UInt32 default 0,
    hour_host_11        UInt32 default 0,
    hour_host_12        UInt32 default 0,
    hour_host_13        UInt32 default 0,
    hour_host_14        UInt32 default 0,
    hour_host_15        UInt32 default 0,
    hour_host_16        UInt32 default 0,
    hour_host_17        UInt32 default 0,
    hour_host_18        UInt32 default 0,
    hour_host_19        UInt32 default 0,
    hour_host_20        UInt32 default 0,
    hour_host_21        UInt32 default 0,
    hour_host_22        UInt32 default 0,
    hour_host_23        UInt32 default 0,
    hour_guest_0        UInt32 default 0,
    hour_guest_1        UInt32 default 0,
    hour_guest_2        UInt32 default 0,
    hour_guest_3        UInt32 default 0,
    hour_guest_4        UInt32 default 0,
    hour_guest_5        UInt32 default 0,
    hour_guest_6        UInt32 default 0,
    hour_guest_7        UInt32 default 0,
    hour_guest_8        UInt32 default 0,
    hour_guest_9        UInt32 default 0,
    hour_guest_10       UInt32 default 0,
    hour_guest_11       UInt32 default 0,
    hour_guest_12       UInt32 default 0,
    hour_guest_13       UInt32 default 0,
    hour_guest_14       UInt32 default 0,
    hour_guest_15       UInt32 default 0,
    hour_guest_16       UInt32 default 0,
    hour_guest_17       UInt32 default 0,
    hour_guest_18       UInt32 default 0,
    hour_guest_19       UInt32 default 0,
    hour_guest_20       UInt32 default 0,
    hour_guest_21       UInt32 default 0,
    hour_guest_22       UInt32 default 0,
    hour_guest_23       UInt32 default 0,
    hour_new_guest_0    UInt32 default 0,
    hour_new_guest_1    UInt32 default 0,
    hour_new_guest_2    UInt32 default 0,
    hour_new_guest_3    UInt32 default 0,
    hour_new_guest_4    UInt32 default 0,
    hour_new_guest_5    UInt32 default 0,
    hour_new_guest_6    UInt32 default 0,
    hour_new_guest_7    UInt32 default 0,
    hour_new_guest_8    UInt32 default 0,
    hour_new_guest_9    UInt32 default 0,
    hour_new_guest_10   UInt32 default 0,
    hour_new_guest_11   UInt32 default 0,
    hour_new_guest_12   UInt32 default 0,
    hour_new_guest_13   UInt32 default 0,
    hour_new_guest_14   UInt32 default 0,
    hour_new_guest_15   UInt32 default 0,
    hour_new_guest_16   UInt32 default 0,
    hour_new_guest_17   UInt32 default 0,
    hour_new_guest_18   UInt32 default 0,
    hour_new_guest_19   UInt32 default 0,
    hour_new_guest_20   UInt32 default 0,
    hour_new_guest_21   UInt32 default 0,
    hour_new_guest_22   UInt32 default 0,
    hour_new_guest_23   UInt32 default 0,
    hour_session_0      UInt32 default 0,
    hour_session_1      UInt32 default 0,
    hour_session_2      UInt32 default 0,
    hour_session_3      UInt32 default 0,
    hour_session_4      UInt32 default 0,
    hour_session_5      UInt32 default 0,
    hour_session_6      UInt32 default 0,
    hour_session_7      UInt32 default 0,
    hour_session_8      UInt32 default 0,
    hour_session_9      UInt32 default 0,
    hour_session_10     UInt32 default 0,
    hour_session_11     UInt32 default 0,
    hour_session_12     UInt32 default 0,
    hour_session_13     UInt32 default 0,
    hour_session_14     UInt32 default 0,
    hour_session_15     UInt32 default 0,
    hour_session_16     UInt32 default 0,
    hour_session_17     UInt32 default 0,
    hour_session_18     UInt32 default 0,
    hour_session_19     UInt32 default 0,
    hour_session_20     UInt32 default 0,
    hour_session_21     UInt32 default 0,
    hour_session_22     UInt32 default 0,
    hour_session_23     UInt32 default 0,
    hour_hit_0          UInt32 default 0,
    hour_hit_1          UInt32 default 0,
    hour_hit_2          UInt32 default 0,
    hour_hit_3          UInt32 default 0,
    hour_hit_4          UInt32 default 0,
    hour_hit_5          UInt32 default 0,
    hour_hit_6          UInt32 default 0,
    hour_hit_7          UInt32 default 0,
    hour_hit_8          UInt32 default 0,
    hour_hit_9          UInt32 default 0,
    hour_hit_10         UInt32 default 0,
    hour_hit_11         UInt32 default 0,
    hour_hit_12         UInt32 default 0,
    hour_hit_13         UInt32 default 0,
    hour_hit_14         UInt32 default 0,
    hour_hit_15         UInt32 default 0,
    hour_hit_16         UInt32 default 0,
    hour_hit_17         UInt32 default 0,
    hour_hit_18         UInt32 default 0,
    hour_hit_19         UInt32 default 0,
    hour_hit_20         UInt32 default 0,
    hour_hit_21         UInt32 default 0,
    hour_hit_22         UInt32 default 0,
    hour_hit_23         UInt32 default 0,
    hour_event_0        UInt32 default 0,
    hour_event_1        UInt32 default 0,
    hour_event_2        UInt32 default 0,
    hour_event_3        UInt32 default 0,
    hour_event_4        UInt32 default 0,
    hour_event_5        UInt32 default 0,
    hour_event_6        UInt32 default 0,
    hour_event_7        UInt32 default 0,
    hour_event_8        UInt32 default 0,
    hour_event_9        UInt32 default 0,
    hour_event_10       UInt32 default 0,
    hour_event_11       UInt32 default 0,
    hour_event_12       UInt32 default 0,
    hour_event_13       UInt32 default 0,
    hour_event_14       UInt32 default 0,
    hour_event_15       UInt32 default 0,
    hour_event_16       UInt32 default 0,
    hour_event_17       UInt32 default 0,
    hour_event_18       UInt32 default 0,
    hour_event_19       UInt32 default 0,
    hour_event_20       UInt32 default 0,
    hour_event_21       UInt32 default 0,
    hour_event_22       UInt32 default 0,
    hour_event_23       UInt32 default 0,
    hour_favorite_0     UInt32 default 0,
    hour_favorite_1     UInt32 default 0,
    hour_favorite_2     UInt32 default 0,
    hour_favorite_3     UInt32 default 0,
    hour_favorite_4     UInt32 default 0,
    hour_favorite_5     UInt32 default 0,
    hour_favorite_6     UInt32 default 0,
    hour_favorite_7     UInt32 default 0,
    hour_favorite_8     UInt32 default 0,
    hour_favorite_9     UInt32 default 0,
    hour_favorite_10    UInt32 default 0,
    hour_favorite_11    UInt32 default 0,
    hour_favorite_12    UInt32 default 0,
    hour_favorite_13    UInt32 default 0,
    hour_favorite_14    UInt32 default 0,
    hour_favorite_15    UInt32 default 0,
    hour_favorite_16    UInt32 default 0,
    hour_favorite_17    UInt32 default 0,
    hour_favorite_18    UInt32 default 0,
    hour_favorite_19    UInt32 default 0,
    hour_favorite_20    UInt32 default 0,
    hour_favorite_21    UInt32 default 0,
    hour_favorite_22    UInt32 default 0,
    hour_favorite_23    UInt32 default 0,
    weekday_host_0      UInt32 default 0,
    weekday_host_1      UInt32 default 0,
    weekday_host_2      UInt32 default 0,
    weekday_host_3      UInt32 default 0,
    weekday_host_4      UInt32 default 0,
    weekday_host_5      UInt32 default 0,
    weekday_host_6      UInt32 default 0,
    weekday_guest_0     UInt32 default 0,
    weekday_guest_1     UInt32 default 0,
    weekday_guest_2     UInt32 default 0,
    weekday_guest_3     UInt32 default 0,
    weekday_guest_4     UInt32 default 0,
    weekday_guest_5     UInt32 default 0,
    weekday_guest_6     UInt32 default 0,
    weekday_new_guest_0 UInt32 default 0,
    weekday_new_guest_1 UInt32 default 0,
    weekday_new_guest_2 UInt32 default 0,
    weekday_new_guest_3 UInt32 default 0,
    weekday_new_guest_4 UInt32 default 0,
    weekday_new_guest_5 UInt32 default 0,
    weekday_new_guest_6 UInt32 default 0,
    weekday_session_0   UInt32 default 0,
    weekday_session_1   UInt32 default 0,
    weekday_session_2   UInt32 default 0,
    weekday_session_3   UInt32 default 0,
    weekday_session_4   UInt32 default 0,
    weekday_session_5   UInt32 default 0,
    weekday_session_6   UInt32 default 0,
    weekday_hit_0       UInt32 default 0,
    weekday_hit_1       UInt32 default 0,
    weekday_hit_2       UInt32 default 0,
    weekday_hit_3       UInt32 default 0,
    weekday_hit_4       UInt32 default 0,
    weekday_hit_5       UInt32 default 0,
    weekday_hit_6       UInt32 default 0,
    weekday_event_0     UInt32 default 0,
    weekday_event_1     UInt32 default 0,
    weekday_event_2     UInt32 default 0,
    weekday_event_3     UInt32 default 0,
    weekday_event_4     UInt32 default 0,
    weekday_event_5     UInt32 default 0,
    weekday_event_6     UInt32 default 0,
    weekday_favorite_0  UInt32 default 0,
    weekday_favorite_1  UInt32 default 0,
    weekday_favorite_2  UInt32 default 0,
    weekday_favorite_3  UInt32 default 0,
    weekday_favorite_4  UInt32 default 0,
    weekday_favorite_5  UInt32 default 0,
    weekday_favorite_6  UInt32 default 0,
    month_host_1        UInt32 default 0,
    month_host_2        UInt32 default 0,
    month_host_3        UInt32 default 0,
    month_host_4        UInt32 default 0,
    month_host_5        UInt32 default 0,
    month_host_6        UInt32 default 0,
    month_host_7        UInt32 default 0,
    month_host_8        UInt32 default 0,
    month_host_9        UInt32 default 0,
    month_host_10       UInt32 default 0,
    month_host_11       UInt32 default 0,
    month_host_12       UInt32 default 0,
    month_guest_1       UInt32 default 0,
    month_guest_2       UInt32 default 0,
    month_guest_3       UInt32 default 0,
    month_guest_4       UInt32 default 0,
    month_guest_5       UInt32 default 0,
    month_guest_6       UInt32 default 0,
    month_guest_7       UInt32 default 0,
    month_guest_8       UInt32 default 0,
    month_guest_9       UInt32 default 0,
    month_guest_10      UInt32 default 0,
    month_guest_11      UInt32 default 0,
    month_guest_12      UInt32 default 0,
    month_new_guest_1   UInt32 default 0,
    month_new_guest_2   UInt32 default 0,
    month_new_guest_3   UInt32 default 0,
    month_new_guest_4   UInt32 default 0,
    month_new_guest_5   UInt32 default 0,
    month_new_guest_6   UInt32 default 0,
    month_new_guest_7   UInt32 default 0,
    month_new_guest_8   UInt32 default 0,
    month_new_guest_9   UInt32 default 0,
    month_new_guest_10  UInt32 default 0,
    month_new_guest_11  UInt32 default 0,
    month_new_guest_12  UInt32 default 0,
    month_session_1     UInt32 default 0,
    month_session_2     UInt32 default 0,
    month_session_3     UInt32 default 0,
    month_session_4     UInt32 default 0,
    month_session_5     UInt32 default 0,
    month_session_6     UInt32 default 0,
    month_session_7     UInt32 default 0,
    month_session_8     UInt32 default 0,
    month_session_9     UInt32 default 0,
    month_session_10    UInt32 default 0,
    month_session_11    UInt32 default 0,
    month_session_12    UInt32 default 0,
    month_hit_1         UInt32 default 0,
    month_hit_2         UInt32 default 0,
    month_hit_3         UInt32 default 0,
    month_hit_4         UInt32 default 0,
    month_hit_5         UInt32 default 0,
    month_hit_6         UInt32 default 0,
    month_hit_7         UInt32 default 0,
    month_hit_8         UInt32 default 0,
    month_hit_9         UInt32 default 0,
    month_hit_10        UInt32 default 0,
    month_hit_11        UInt32 default 0,
    month_hit_12        UInt32 default 0,
    month_event_1       UInt32 default 0,
    month_event_2       UInt32 default 0,
    month_event_3       UInt32 default 0,
    month_event_4       UInt32 default 0,
    month_event_5       UInt32 default 0,
    month_event_6       UInt32 default 0,
    month_event_7       UInt32 default 0,
    month_event_8       UInt32 default 0,
    month_event_9       UInt32 default 0,
    month_event_10      UInt32 default 0,
    month_event_11      UInt32 default 0,
    month_event_12      UInt32 default 0,
    month_favorite_1    UInt32 default 0,
    month_favorite_2    UInt32 default 0,
    month_favorite_3    UInt32 default 0,
    month_favorite_4    UInt32 default 0,
    month_favorite_5    UInt32 default 0,
    month_favorite_6    UInt32 default 0,
    month_favorite_7    UInt32 default 0,
    month_favorite_8    UInt32 default 0,
    month_favorite_9    UInt32 default 0,
    month_favorite_10   UInt32 default 0,
    month_favorite_11   UInt32 default 0,
    month_favorite_12   UInt32 default 0
) engine = SummingMergeTree((hits, hosts, sessions, events, guests, new_guests, favorites, total_hosts,
                             am_1, am_1_3, am_3_6, am_6_9, am_9_12, am_12_15, am_15_18, am_18_21, am_21_24, am_24,
                             ah_1, ah_2_5, ah_6_9, ah_10_13, ah_14_17, ah_18_21, ah_22_25, ah_26_29,
                             ah_30_33, ah_34,
                             hour_host_0, hour_host_1, hour_host_2, hour_host_3, hour_host_4,
                             hour_host_5, hour_host_6, hour_host_7, hour_host_8, hour_host_9,
                             hour_host_10, hour_host_11, hour_host_12, hour_host_13, hour_host_14, hour_host_15,
                             hour_host_16, hour_host_17, hour_host_18, hour_host_19, hour_host_20, hour_host_21,
                             hour_host_22, hour_host_23, hour_guest_0, hour_guest_1, hour_guest_2, hour_guest_3,
                             hour_guest_4, hour_guest_5, hour_guest_6, hour_guest_7, hour_guest_8, hour_guest_9,
                             hour_guest_10, hour_guest_11, hour_guest_12, hour_guest_13,
                             hour_guest_14, hour_guest_15, hour_guest_16, hour_guest_17, hour_guest_18,
                             hour_guest_19, hour_guest_20, hour_guest_21, hour_guest_22, hour_new_guest_0,
                             hour_new_guest_1, hour_new_guest_2, hour_new_guest_3, hour_new_guest_4,
                             hour_new_guest_5, hour_new_guest_6, hour_new_guest_7, hour_new_guest_8,
                             hour_new_guest_9, hour_new_guest_10, hour_new_guest_11, hour_new_guest_12,
                             hour_new_guest_13, hour_new_guest_14, hour_new_guest_15, hour_new_guest_16,
                             hour_new_guest_17, hour_new_guest_18, hour_new_guest_19, hour_new_guest_20,
                             hour_new_guest_21, hour_new_guest_22, hour_new_guest_23, hour_session_0, hour_session_1,
                             hour_session_2, hour_session_3, hour_session_4, hour_session_5,
                             hour_session_6, hour_session_7, hour_session_8, hour_session_9,
                             hour_session_10, hour_session_11, hour_session_12, hour_session_13,
                             hour_session_14, hour_session_15, hour_session_16, hour_session_17,
                             hour_session_18, hour_session_19, hour_session_20, hour_session_21,
                             hour_session_22, hour_session_23, hour_hit_0, hour_hit_1, hour_hit_2, hour_hit_3,
                             hour_hit_4, hour_hit_5, hour_hit_6, hour_hit_7, hour_hit_8,
                             hour_hit_9, hour_hit_10, hour_hit_11, hour_hit_12, hour_hit_13,
                             hour_hit_14, hour_hit_15, hour_hit_16, hour_hit_17, hour_hit_18, hour_hit_19, hour_hit_20,
                             hour_hit_21, hour_hit_22, hour_hit_23, hour_event_0, hour_event_1, hour_event_2,
                             hour_event_3, hour_event_4, hour_event_5, hour_event_6, hour_event_7, hour_event_8,
                             hour_event_9, hour_event_10, hour_event_11, hour_event_12,
                             hour_event_13, hour_event_14, hour_event_15, hour_event_16,
                             hour_event_17, hour_event_18, hour_event_19, hour_event_20,
                             hour_event_21, hour_event_22, hour_event_23, hour_favorite_0,
                             hour_favorite_1, hour_favorite_2, hour_favorite_3, hour_favorite_4,
                             hour_favorite_5, hour_favorite_6, hour_favorite_7, hour_favorite_8,
                             hour_favorite_9, hour_favorite_10, hour_favorite_11, hour_favorite_12,
                             hour_favorite_13, hour_favorite_14, hour_favorite_15, hour_favorite_16,
                             hour_favorite_17, hour_favorite_18, hour_favorite_19, hour_favorite_20,
                             hour_favorite_21, hour_favorite_22, hour_favorite_23, weekday_host_0,
                             weekday_host_1, weekday_host_2, weekday_host_3, weekday_host_4,
                             weekday_host_5, weekday_host_6, weekday_guest_0, weekday_guest_1, weekday_guest_2,
                             weekday_guest_3, weekday_guest_4, weekday_guest_5, weekday_guest_6, weekday_new_guest_0,
                             weekday_new_guest_1, weekday_new_guest_2, weekday_new_guest_3, weekday_new_guest_4,
                             weekday_new_guest_5, weekday_new_guest_6, weekday_session_0, weekday_session_1,
                             weekday_session_2, weekday_session_3, weekday_session_4, weekday_session_5,
                             weekday_session_6, weekday_hit_0, weekday_hit_1, weekday_hit_2, weekday_hit_3,
                             weekday_hit_4, weekday_hit_5, weekday_hit_6, weekday_event_0, weekday_event_1,
                             weekday_event_2, weekday_event_3, weekday_event_4, weekday_event_5, weekday_event_6,
                             weekday_favorite_0, weekday_favorite_1, weekday_favorite_2, weekday_favorite_3,
                             weekday_favorite_4, weekday_favorite_5, weekday_favorite_6, month_host_1, month_host_2,
                             month_host_3, month_host_4, month_host_5, month_host_6, month_host_7,
                             month_host_8, month_host_9, month_host_10, month_host_11,
                             month_host_12, month_guest_1, month_guest_2,
                             month_guest_3, month_guest_4, month_guest_5, month_guest_6,
                             month_guest_7, month_guest_8, month_guest_9,
                             month_guest_10, month_guest_11, month_guest_12,
                             month_new_guest_1, month_new_guest_2, month_new_guest_3, month_new_guest_4,
                             month_new_guest_5, month_new_guest_6, month_new_guest_7, month_new_guest_8,
                             month_new_guest_9, month_new_guest_10, month_new_guest_11, month_new_guest_12,
                             month_session_1, month_session_2, month_session_3,
                             month_session_4, month_session_5, month_session_6, month_session_7,
                             month_session_8, month_session_9, month_session_10, month_session_11,
                             month_session_12, month_hit_1, month_hit_2,
                             month_hit_3, month_hit_4, month_hit_5, month_hit_6,
                             month_hit_7, month_hit_8, month_hit_9, month_hit_10,
                             month_hit_11, month_hit_12, month_event_1, month_event_2,
                             month_event_3, month_event_4, month_event_5, month_event_6,
                             month_event_7, month_event_8, month_event_9, month_event_10, month_event_11,
                             month_event_12, month_favorite_1, month_favorite_2, month_favorite_3,
                             month_favorite_4, month_favorite_5, month_favorite_6, month_favorite_7,
                             month_favorite_8, month_favorite_9, month_favorite_10, month_favorite_11,
                             month_favorite_12))
      ORDER BY (date_stat, site_id);


---------------------------- Event --------------------------------

create table if not exists event
(
    uuid              UUID,
    event1            String,
    event2            String,
    money             decimal(18, 4) default 0.0000,
    dateEnter       DateTime32('Europe/Moscow'),
    dateCleanup     DateTime32('Europe/Moscow'),
    sort              UInt32         default 100,
    counter           UInt32         default 0,
    advVisible      BOOLEAN default true,
    name              String,
    description       String,
    keepDays        UInt32,
    dynamicKeepDays UInt32,
    diagramDefault  BOOLEAN default true
) engine = MergeTree
      ORDER BY (event1, event2, keepDays);

create table if not exists event_day
(
    uuid       UUID,
    dateStat  DateTime32('Europe/Moscow'),
    eventUuid UInt32 default 0,
    money      decimal(18, 4) default 0.0000,
    counter    UInt32         default 0
) engine = MergeTree
      PARTITION BY toMonth(dateStat)
      ORDER BY (eventUuid, dateStat);

create table if not exists event_list
(
    uuid            UUID,
    eventUuid     UUID,
    event3          String,
    money           decimal(18, 4) default 0.0000,
    dateEnter     DateTime32('Europe/Moscow'),
    refererUrl    String,
    url             String,
    redirectUrl   String,
    sessionUuid   UUID,
    guestUuid     UUID,
    guestAdvUuid  UUID,
    advBack       BOOLEAN default false,
    hitUuid       UUID,
    countryId     FixedString(2),
    keepDays      UInt32,
    chargeback      bool           default false,
    siteId        FixedString(2),
    refererSiteId FixedString(2)
) engine = MergeTree
      PARTITION BY toMonth(dateEnter)
      ORDER BY (dateEnter);

----------------------- Guest ---------------------------

-- CREATE TABLE guest_buffer AS guest ENGINE = Buffer(statistic, guest, 1, 30, 40, 0, 10000, 0, 0);


create table if not exists guest
(
    uuid     UUID,
    dateAdd DateTime32('Europe/Moscow'),
    repair   boolean default false
)
    engine = MergeTree()
        PARTITION BY toMonth(dateAdd)
        ORDER BY dateAdd;


----------------------- Hit ---------------------------
create table if not exists hit
(
    uuid           UUID,
    sessionUuid  UUID,
    advUuid      UUID,
    dateHit      DateTime32('Europe/Moscow'),
    phpSessionId String,
    guestUuid    UUID,
    language       String,
    isNewGuest   BOOLEAN default false,
    userId       UInt32,
    userAuth     BOOLEAN default false,
    url            String,
    url404       BOOLEAN default false,
    urlFrom      String,
    ip             IPv4,
    method         String,
    cookies        String,
    userAgent    String,
    stopListUuid UUID,
    countryId    FixedString(2),
    cityUuid     UUID,
    siteId       FixedString(2),
    favorites      boolean default false
)
    engine = MergeTree
        PARTITION BY toMonth(dateHit)
        ORDER BY dateHit;

create table hit_full
(
    advBack          Bool DEFAULT false,
    advReferer1      String,
    advReferer2      String,
    advReferer3      String,
    advUuid          UUID,
    dateHit          DateTime('Europe/Moscow'),
    cookies            String,
    countryId        FixedString(2),
    city               String,
    event1             String,
    event2             String,
    eventTypeUuid    String,
    guestUuid        UUID,
    isNewGuest       Bool DEFAULT false,
    ip                 IPv4,
    method             String,
    phpSessionId     String,
    referrerSiteName String,
    referrerUrlFrom  String,
    searcherName     String,
    searcherPhrase   String,
    searcherUuid     String,
    sessionUuid      UUID,
    siteId           FixedString(2),
    url                String,
    url404           Bool DEFAULT false,
    urlFrom          String,
    urlIsDir         Bool DEFAULT false,
    userAgent        String,
    userAuth         Bool DEFAULT false,
    userId           UInt32,
    uuid               UUID
)
    engine = MergeTree
        PARTITION BY toMonth(dateHit)
        ORDER BY dateHit;


------------------ Page ----------------------

create table if not exists page
(
    uuid          UUID,
    dateStat     DateTime32('Europe/Moscow'),
    dir           BOOLEAN default false,
    url           String,
    url404       BOOLEAN default false,
    urlHash      UInt32,
    siteId       FixedString(2),
    counter       UInt32  default 0,
    enterCounter UInt32  default 0,
    exitCounter  UInt32  default 0,
    sign          Int8,
    version       UInt32
) engine = VersionedCollapsingMergeTree(sign, version)
      PARTITION BY toMonth(dateStat)
      ORDER BY dateStat;

create table if not exists page_adv
(
    dateStat         DateTime32('Europe/Moscow'),
    pageUuid         UUID,
    advUuid          UUID,
    enterCounter     UInt32,
    exitCounter      UInt32,
    counterBack      UInt32,
    enterCounterBack UInt32,
    exitCounterBack  UInt32

) engine = SummingMergeTree((dateStat, pageUuid, advUuid))
      PARTITION BY toMonth(dateStat)
      ORDER BY dateStat;

---------------------- Path ------------------------

create table if not exists path
(
    uuid               UUID,
    pathId          Int32   default 0,
    parentPathId    UInt32,
    dateStat        DATE,
    counter            UInt32  default 0,
    counterAbnormal UInt32  default 0,
    counterFullPath UInt32  default 0,
    pages              String,
    firstPage       String,
    firstPage404    BOOLEAN default false,
    firstPageSiteId FixedString(2),
    prevPage        String,
    prevPageHash    UInt32,
    lastPage        String,
    lastPage404     bool    default false,
    lastPageSiteId  FixedString(2),
    lastPageHash    UInt32,
    steps              UInt32  default 1,
    sign               Int8,
    version            UInt32
) engine = VersionedCollapsingMergeTree(sign, version)
      PARTITION BY toMonth(dateStat)
      ORDER BY dateStat;

create table if not exists path_adv
(
    advUuid             UUID,
    pathId              Int32,
    dateStat            DATE,
    counter                UInt32 default 0,
    counterBack         UInt32 default 0,
    counterFullPath     UInt32 default 0,
    counterFullPathBack UInt32 default 0,
    steps                  UInt32 default 0,
    sign                   Int8,
    version                UInt32
) engine = VersionedCollapsingMergeTree(sign, version)
      PARTITION BY toMonth(dateStat)
      ORDER BY dateStat;


create table if not exists path_cache
(
    uuid                    UUID,
    sessionUuid         UUID,
    dateHit             DateTime32('Europe/Moscow'),
    pathId              Int32,
    pathPages           String,
    pathFirstPage       String,
    pathFirstPage404    BOOLEAN default false,
    pathFirstPageSiteId FixedString(2),
    pathLastPage        String,
    pathLastPage404     BOOLEAN default false,
    pathPageSiteId      FixedString(2),
    pathSteps           UInt32  default 1,
    isLastPage          BOOLEAN default true,
    sign                    Int8,
    version                 UInt32
) engine = VersionedCollapsingMergeTree(sign, version)
      PARTITION BY toMonth(dateHit)
      ORDER BY dateHit;


----------------------- Phrase ----------------------------

create table if not exists phrase_list
(
    uuid          UUID,
    dateHit      DateTime32('Europe/Moscow'),
    searcherUuid UUID,
    refererUuid  UUID,
    phrase        String,
    urlFrom      String,
    urlTo        String,
    urlTo404     bool default false,
    sessionUuid  UUID,
    siteId       FixedString(2)
) engine = MergeTree
      PARTITION BY toMonth(dateHit)
      ORDER BY dateHit;

--------------------- Referer -----------------------------

create table if not exists referer
(
    uuid      UUID,
    siteName String,
    sessions  UInt32 default 0,
    hits      UInt32 default 0
) engine = SummingMergeTree((sessions, hits))
      ORDER BY siteName;

create table if not exists referer_list
(
    uuid         UUID,
    refererUuid UUID,
    dateHit     DateTime32('Europe/Moscow'),
    protocol     String,
    siteName    String,
    urlFrom     String,
    urlTo       String,
    urlTo404    bool default false,
    sessionUuid UUID,
    advUuid     UUID,
    siteId      FixedString(2)
) engine = MergeTree
      PARTITION BY toMonth(dateHit)
      ORDER BY dateHit;

--------------------- Searcher -------------------------
create table if not exists searcher
(
    uuid              UUID,
    date_cleanup      Nullable(DateTime32('Europe/Moscow')),
    totalHits       UInt32  default 0,
    saveStatistic   BOOLEAN default true,
    active            BOOLEAN default true,
    name              String,
    userAgent       String,
    diagramDefault  BOOLEAN default false,
    hitKeepDays     UInt32,
    dynamicKeepDays UInt32,
    checkActivity   BOOLEAN default true
) engine = MergeTree
      ORDER BY name;

create table searcher_phrase_stat
(
    searcherUuid UUID,
    phrases       UInt32 default 0,
    phrasesHits  UInt32 default 0
)
    engine = SummingMergeTree((phrases, phrasesHits))
        ORDER BY searcherUuid;

create table if not exists searcher_day_hits
(
    hitDay       Date,
    searcherUuid UUID,
    totalHits    UInt64
) engine = SummingMergeTree(totalHits)
      ORDER BY (hitDay, searcherUuid);

create table if not exists searcher_hit
(
    uuid          UUID,
    dateHit      DateTime32('Europe/Moscow'),
    searcherUuid UUID,
    url           String,
    url404       BOOLEAN default false,
    ip            IPv4,
    userAgent    String,
    siteId       FixedString(2)
) engine = MergeTree
      PARTITION BY toMonth(dateHit)
      ORDER BY (dateHit, searcherUuid);

create table if not exists searcher_params
(
    uuid          UUID,
    searcherUuid UUID,
    domain        String,
    variable      String,
    charSet      String
) engine = MergeTree
      ORDER BY (domain);


--------------------- session ---------------------------
-- create table if not exists session
-- (
--     uuid           UUID,
--     guest_uuid     UUID,
--     new_guest      boolean,
--     user_id        UInt32,
--     user_auth      boolean,
--     events         UInt32 default 0,
--     hits           UInt32 default 0,
--     favorites      boolean,
--     url_from       String,
--     url_to         String,
--     url_to_404     boolean,
--     url_last       String,
--     url_last_404   bool,
--     user_agent     String,
--     date_stat      DATE,
--     date_first     DateTime32('Europe/Moscow'),
--     date_last      DateTime32('Europe/Moscow'),
--     ip_first       IPv4,
--     ip_last        IPv4,
--     first_hit_uuid UUID,
--     last_hit_uuid  UUID,
--     php_session_id String
--
--
-- ) ENGINE = VersionedCollapsingMergeTree(sign, version)
--       PARTITION BY toMonth(date_stat)
--       ORDER BY (date_stat);

create table if not exists session
(
    uuid           UUID,
    guestUuid    UUID,
    phpSessionId String,
    dateAdd      DateTime32('Europe/Moscow')
) ENGINE = MergeTree
      PARTITION BY toMonth(dateAdd)
      ORDER BY (dateAdd);


------------------- Option -----------------------

create table if not exists options
(
    name  String,
    value String
) ENGINE = MergeTree
      ORDER BY (name);

create table if not exists raw_request
(
    date                 DateTime32('Europe/Moscow'),
    phpSessionId      String,
    url                  String,
    referer              String,
    ip                   IPv4,
    userAgent         String,
    userid               UInt32,
    userLogin         String,
    httpXForwardedFor String,
    isError404        bool,
    siteId            String,
    event1               String,
    event2               String,
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


