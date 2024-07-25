USE statistic;
---------------------- ADV -------------------------

create table if not exists adv
(
    uuid        UUID,
    referer1    String,
    referer2    String,
    date_create DateTime32('Europe/Moscow'),
    cost        decimal(18, 4) default 0.0000,
    events_view String,
    description String,
    priority    UInt32         default 1
) ENGINE = MergeTree
      ORDER BY (referer1);

create table if not exists adv_stat
(
    adv_uuid       UUID,
    revenue        decimal(18, 4) default 0.0000,
    guests         UInt32         default 0,
    new_guests     UInt32         default 0,
    favorites      UInt32         default 0,
    hosts          UInt32         default 0,
    sessions       UInt32         default 0,
    hits           UInt32         default 0,
    guests_back    UInt32         default 0,
    favorites_back UInt32         default 0,
    hosts_back     UInt32         default 0,
    sessions_back  UInt32         default 0,
    hits_back      UInt32         default 0
) engine = SummingMergeTree((revenue,
                             guests,
                             new_guests,
                             favorites,
                             hosts,
                             sessions,
                             hits,
                             guests_back,
                             favorites_back,
                             hosts_back,
                             sessions_back,
                             hits_back))
      ORDER BY (adv_uuid);


create table if not exists adv_day
(
    adv_uuid        UUID,
    date_stat       Date32,
    guests          UInt32 default 0,
    guests_day      UInt32 default 0,
    new_guests      UInt32 default 0,
    favorites       UInt32 default 0,
    hosts           UInt32 default 0,
    hosts_day       UInt32 default 0,
    sessions        UInt32 default 0,
    hits            UInt32 default 0,
    guests_back     UInt32 default 0,
    guests_day_back UInt32 default 0,
    favorites_back  UInt32 default 0,
    hosts_back      UInt32 default 0,
    hosts_day_back  UInt32 default 0,
    sessions_back   UInt32 default 0,
    hits_back       UInt32 default 0
) ENGINE = SummingMergeTree(
        (guests,
         guests_day,
         new_guests,
         favorites,
         hosts,
         hosts_day,
         sessions,
         hits,
         guests_back,
         guests_day_back,
         favorites_back,
         hosts_back,
         hosts_day_back,
         sessions_back,
         hits_back)
           )
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (adv_uuid, date_stat);

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
      ORDER BY (money);

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
      ORDER BY (adv_uuid, event_uuid, date_stat);

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
      ORDER BY (adv_uuid, guest_uuid);


create table if not exists adv_page
(
    uuid     UUID,
    adv_uuid UUID,
    page     String,
    type     String default 'TO'
) ENGINE = MergeTree
      ORDER BY (adv_uuid, type);

create table if not exists adv_searcher
(
    uuid          UUID,
    adv_uuid      UUID,
    searcher_uuid UUID
) ENGINE = MergeTree
      ORDER BY (adv_uuid, searcher_uuid);

----------------------- Browser --------------------------
create table if not exists browser
(
    uuid       UUID,
    user_agent String
) ENGINE = MergeTree
      ORDER BY (user_agent);

------------------------ City -----------------------------

create table if not exists city
(
    uuid       UUID,
    country_id String,
    region     String,
    name       String,
    xml_id     String,
    sessions   UInt32 default 0,
    new_guests UInt32 default 0,
    hits       UInt32 default 0,
    events     UInt32 default 0
) engine = MergeTree
      ORDER BY (country_id, region, name);

create table if not exists city_day
(
    city_uuid  UUID,
    date_stat  Date32,
    sessions   UInt32 default 0,
    new_guests UInt32 default 0,
    hits       UInt32 default 0,
    events     UInt32 default 0
) engine = SummingMergeTree((sessions, new_guests, hits, events))
      ORDER BY (city_uuid, date_stat);

create table if not exists city_ip
(
    start_ip   UInt32,
    end_ip     UInt32,
    country_id String,
    city_uuid  UUID
) engine = MergeTree
      ORDER BY (end_ip);

------------------ Country ---------------------

create table if not exists country
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

create table if not exists country_day
(
    country_id FixedString(2),
    date_stat  Date32,
    sessions   UInt32 default 0,
    new_guests UInt32 default 0,
    hits       UInt32 default 0,
    events     UInt32 default 0
) engine = SummingMergeTree((sessions, new_guests, hits, events))
      ORDER BY (country_id, date_stat);

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
    date_enter        DateTime32('Europe/Moscow'),
    date_cleanup      DateTime32('Europe/Moscow'),
    sort              UInt32         default 100,
    counter           UInt32         default 0,
    adv_visible       BOOLEAN        default true,
    name              String,
    description       String,
    keep_days         UInt32,
    dynamic_keep_days UInt32,
    diagram_default   BOOLEAN        default true
) engine = MergeTree
      ORDER BY (event1, event2, keep_days);

create table if not exists event_day
(
    uuid       UUID,
    date_stat  DateTime32('Europe/Moscow'),
    event_uuid UInt32         default 0,
    money      decimal(18, 4) default 0.0000,
    counter    UInt32         default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (event_uuid, date_stat);

create table if not exists event_list
(
    uuid            UUID,
    event_uuid      int            default 0,
    event3          String,
    money           decimal(18, 4) default 0.0000,
    date_enter      DateTime32('Europe/Moscow'),
    referer_url     String,
    url             String,
    redirect_url    String,
    session_uuid    int,
    guest_uuid      int,
    guest_adv_uuid  int,
    adv_back        BOOLEAN        default false,
    hit_uuid        UUID,
    country_id      FixedString(2),
    keep_days       int,
    chargeback      char           default 'N',
    site_id         FixedString(2),
    referer_site_id FixedString(2)
) engine = MergeTree
      PARTITION BY toYYYYMM(date_enter)
      ORDER BY (date_enter);

----------------------- Guest ---------------------------

CREATE TABLE statistic.guest_buffer AS statistic.guest ENGINE = Buffer(statistic, guest, 1, 30, 40, 0, 10000, 0, 0);

create table if not exists guest
(
    uuid         UUID,
    timestamp    DateTime32('Europe/Moscow'),
    favorites    boolean default false,
    events       Int32   default 0,
    sessions     Int32   default 0,
    hits         Int32   default 0,
    repair       boolean default false,
    session_id   UUID,
    date         DateTime32('Europe/Moscow'),
    url_from     String,
    url_to       String,
    url_404      boolean default false,
    site_id      String,
    adv_uuid     UUID,
    referer1     String,
    referer2     String,
    referer3     String,
    session_uuid UUID,
    user_id      Int32,
    user_auth    boolean,
    url          String,
    user_agent   String,
    ip           IPv4,
    cookie       String,
    language     String,
    adv_id       UUID,
    adv_back     boolean default false,
    city_id      UUID,
    guest_hash   String
)
    engine = MergeTree
        PARTITION BY toYYYYMM(timestamp)
        ORDER BY timestamp;


----------------------- Hit ---------------------------
create table if not exists hit
(
    uuid         UUID,
    session_id   UInt32  default 0,
    date_hit     DateTime32('Europe/Moscow'),
    guest_uuid   UUID,
    new_guest    BOOLEAN default false,
    user_id      int,
    user_auth    BOOLEAN default false,
    url          String,
    url_404      BOOLEAN default false,
    url_from     String,
    ip           IPv4,
    method       String,
    cookies      String,
    user_agent   String,
    stop_list_id int,
    country_id   FixedString(2),
    city_uuid    UUID,
    site_id      FixedString(2)
)
    engine = MergeTree
        PARTITION BY toYYYYMM(date_hit)
        ORDER BY date_hit;

------------------ Page ----------------------

create table if not exists page
(
    uuid          UUID,
    date_stat     DateTime32('Europe/Moscow'),
    dir           BOOLEAN default false,
    url           String,
    url_404       BOOLEAN default false,
    url_hash      UInt32,
    site_id       FixedString(2),
    counter       UInt32  default 0,
    enter_counter UInt32  default 0,
    exit_counter  UInt32  default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY date_stat;

---------------------- Path ------------------------

create table if not exists path
(
    uuid              UUID,
    path_id           UInt32  default 0,
    parent_path_id    UInt32,
    date_stat         DateTime32('Europe/Moscow'),
    counter           UInt32  default 0,
    counter_abnormal  UInt32  default 0,
    counter_full_path UInt32  default 0,
    pages             String,
    page              String,
    page_404          BOOLEAN default false,
    page_site_id      FixedString(2),
    prev_page         String,
    prev_page_hash    UInt32,
    page_hash         UInt32,
    steps             UInt32  default 1
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY date_stat;

create table if not exists path_adv
(
    uuid                   UUID,
    adv_uuid               UUID,
    path_uuid              UUID,
    date_stat              DateTime32('Europe/Moscow'),
    counter                UInt32 default 0,
    counter_back           UInt32 default 0,
    counter_full_path      UInt32 default 0,
    counter_full_path_back UInt32 default 0,
    steps                  UInt32 default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY date_stat;


create table if not exists path_cache
(
    uuid                    UUID,
    session_uuid            UUID,
    date_hit                DateTime32('Europe/Moscow'),
    path_uuid               UUID,
    path_pages              String,
    path_first_page         String,
    path_first_page_404     BOOLEAN default false,
    path_first_page_site_id FixedString(2),
    path_page               String,
    path_page_404           BOOLEAN default false,
    path_page_site_id       FixedString(2),
    path_steps              UInt32  default 1,
    is_last_page            BOOLEAN default true
) engine = MergeTree
      PARTITION BY toYYYYMM(date_hit)
      ORDER BY date_hit;


----------------------- Phrase ----------------------------

create table if not exists phrase_list
(
    uuid        UUID,
    date_hit    DateTime32('Europe/Moscow'),
    searcher_id int,
    referer_id  int,
    phrase      String,
    url_from    String,
    url_to      String,
    url_to_404  char default 'N',
    session_id  int,
    site_id     FixedString(2)
) engine = MergeTree
      PARTITION BY toYYYYMM(date_hit)
      ORDER BY date_hit;

--------------------- Referer -----------------------------

create table if not exists referer
(
    uuid      UUID,
    date      DateTime32('Europe/Moscow'),
    site_name String,
    sessions  UInt32 default 0,
    hits      UInt32 default 0
) engine = MergeTree
      PARTITION BY toYYYYMM(date)
      ORDER BY date;

create table if not exists referer_list
(
    uuid         UUID,
    referer_uuid int,
    date_hit     DateTime32('Europe/Moscow'),
    protocol     String,
    site_name    String,
    url_from     String,
    url_to       String,
    url_to_404   char default 'N',
    session_uuid int,
    adv_uuid     UUID,
    site_id      FixedString(2)
) engine = MergeTree
      PARTITION BY toMonth(date_hit)
      ORDER BY date_hit;

--------------------- Searcher -------------------------
create table if not exists searcher
(
    uuid              UUID,
    date_cleanup      Nullable(DateTime32('Europe/Moscow')),
    total_hits        UInt32  default '0',
    save_statistic    BOOLEAN default true,
    active            BOOLEAN default true,
    name              String,
    user_agent        String,
    diagram_default   BOOLEAN default false,
    hit_keep_days     UInt32,
    dynamic_keep_days UInt32,
    phrases           UInt32  default '0',
    phrases_hits      UInt32  default '0',
    check_activity    BOOLEAN default true
) engine = MergeTree
--       PARTITION BY toYYYYMM(date_hit)
      ORDER BY name;

create table if not exists statistic.searcher_day_hits
(
    hit_day       Date,
    searcher_uuid UUID,
    total_hits    UInt64
) engine = SummingMergeTree(total_hits)
      ORDER BY (hit_day, searcher_uuid);

create table if not exists searcher_hit
(
    uuid          UUID,
    date_hit      DateTime32('Europe/Moscow'),
    searcher_uuid UUID,
    url           String,
    url_404       BOOLEAN default false,
    ip            IPv4,
    user_agent    String,
    site_id       FixedString(2)
) engine = MergeTree
--       PARTITION BY toYYYYMM(date_hit)
      ORDER BY (date_hit, searcher_uuid);

create table if not exists searcher_params
(
    uuid          UUID,
    searcher_uuid UUID,
    domain        String,
    variable      String,
    char_set      String
) engine = MergeTree
      ORDER BY (domain);


--------------------- session ---------------------------
create table if not exists session
(
    uuid         UUID,
    guest_id     UUID,
    new_guest    boolean,
    user_id      Int32,
    user_auth    boolean,
    events       Int32 default 0,
    hits         Int32 default 0,
    favorites    boolean,
    url_from     String,
    url_to       String,
    url_to_404   boolean,
    url          String,
    user_agent   String,
    date_stat    DateTime32('Europe/Moscow'),
    date         DateTime32('Europe/Moscow'),
    ip           IPv4,
    hit_id       UUID,
    phpsessid    String,
    adv_id       UUID,
    adv_back     boolean,
    referer1     String,
    referer2     String,
    referer3     String,
    stop_list_id UUID,
    country_id   FixedString(2),
    site_id      String,
    city_id      UInt32

) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date_stat)
      ORDER BY (uuid, date_stat)
      PRIMARY KEY (uuid, date_stat);


------------------- Option -----------------------

create table if not exists option
(
    name        String,
    value       String,
    description String,
    siteId      FixedString(2)
) ENGINE = MergeTree
      PARTITION BY siteId
      ORDER BY (name);

create table if not exists raw_request
(
    date                 DateTime32('Europe/Moscow'),
    php_session_id       String,
    url                  String,
    referer              String,
    ip                   IPv4,
    user_agent           String,
    userid               UInt32,
    user_login           String,
    http_x_forwarded_for String,
    is_error404          bool,
    site_id              String,
    event1               String,
    event2               String,
    is_user_auth         bool
) ENGINE = MergeTree
      PARTITION BY toYYYYMM(date)
      ORDER BY (date);

create table if not exists statistic.searcher_total_hits
(
    date_stat     Date,
    searcher_uuid UUID,
    total_hits    UInt64
) engine = SummingMergeTree(total_hits)
      ORDER BY (date_stat, searcher_uuid);



