package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Guest struct {
	ctx          context.Context
	chClient     driver.Conn
	sessionModel *SessionModel
}

func NewGuest(ctx context.Context, chClient driver.Conn) *Guest {
	return &Guest{
		ctx:          ctx,
		chClient:     chClient,
		sessionModel: NewSessionModel(ctx, chClient),
	}
}

//func (gm Guest) FindLastById(id int) (int, string, int, int, string, error) {
//	row, err := gm.chClient.Query(gm.ctx, `
//				SELECT
//					G.id,
//					G.FAVORITES,
//					G.LAST_USER_ID,
//					A.ID as LAST_ADV_ID,
//					if(to_days(curdate())=to_days(G.LAST_DATE), 'Y', 'N') LAST
//				FROM guest G
//				LEFT JOIN adv A ON A.ID = G.LAST_ADV_ID
//				WHERE G.ID=?`, id)
//	var guestId, lastUserId, lastAdvId int
//	var favorites, last string
//	err := row.Scan(&guestId, favorites, lastUserId, lastAdvId, last)
//	if err != nil {
//		return 0, "", 0, 0, "", err
//	}
//	return guestId, favorites, lastUserId, lastAdvId, last, nil
//}

func (gm Guest) AddGuest(guest entity.GuestDb) error {
	err := gm.chClient.Exec(gm.ctx, `INSERT INTO guest (
                   timestamp_x, favorites, events, sessions, hits, repair, session_id, date, url_from, url_to,
                   url_to_404, site_id, adv_id, referer1, referer2, referer3, user_id, user_auth, url, url_404, user_agent, ip,
                   cookie, language, adv_back, country_id, city_id, city_info, cookie_token) 
		VALUES (?,?,?,?,?,?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		guest.Favorites, guest.Events, guest.Sessions, guest.Hits, guest.Repair, guest.SessionId, guest.Date,
		guest.UrlFrom, guest.UrlTo, guest.UrlTo404, guest.SiteId, guest.AdvId, guest.Referer1, guest.Referer2,
		guest.Referer3, guest.UserId, guest.UserAuth, guest.Url, guest.Url404, guest.UserAgent, guest.Ip,
		guest.Cookie, guest.Language, guest.AdvBack, guest.CountryId, guest.CityId, guest.CityInfo, guest.Token,
	)
	if err != nil {
		return err
	}
	return nil
}

//func (gm Guest) AddGuest(guestDb entity.GuestDb) error {
//
//
//	return nil
//}

func (gm Guest) ExistsGuestByHash(token string) (bool, error) {
	row := gm.chClient.QueryRow(gm.ctx, `
				SELECT guest_hash
				FROM guest 				
				WHERE guest_hash=?`, token)
	var cookieToken string
	err := row.Scan(&cookieToken)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return len(cookieToken) > 0, nil
}

func (gm Guest) Find(filter filters.Filter) (entity.GuestDb, error) {
	return entity.GuestDb{}, nil
}

func (gm Guest) FindByHash(token string) ([]entity.GuestDb, error) {
	row := gm.chClient.QueryRow(gm.ctx, `
				SELECT * 
				FROM guest 				
				WHERE guest_hash=?`, token)
	var guestDb []entity.GuestDb
	err := row.Scan(&guestDb)
	if err != nil {
		return []entity.GuestDb{}, nil
	}

	return guestDb, nil
}
