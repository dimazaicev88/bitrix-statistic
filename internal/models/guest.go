package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Guest struct {
	ctx          context.Context
	chClient     driver.Conn
	sessionModel *Session
}

func NewGuest(ctx context.Context, chClient driver.Conn) *Guest {
	return &Guest{
		ctx:          ctx,
		chClient:     chClient,
		sessionModel: NewSession(ctx, chClient),
	}
}

//func (gm GuestStat) FindLastById(id int) (int, string, int, int, string, error) {
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

func (gm Guest) Add(guest entitydb.Guest) error {
	return gm.chClient.Exec(gm.ctx,
		`INSERT INTO guest (uuid,guest_hash,user_agent,ip, x_forwarded_for,date_create) 
			   VALUES (generateUUIDv7(), ?, ?, ?, ?, now())`,
		guest.GuestHash, guest.UserAgent, guest.Ip, guest.XForwardedFor,
	)
}

func (gm Guest) ExistsByHash(token string) (bool, error) {
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

func (gm Guest) Find(filter filters.Filter) ([]entitydb.GuestStat, error) {
	return []entitydb.GuestStat{}, nil
}

func (gm Guest) FindByHash(token string) ([]entitydb.GuestStat, error) {
	row := gm.chClient.QueryRow(gm.ctx, `
				SELECT * 
				FROM guest 				
				WHERE guest_hash=?`, token)
	var guest []entitydb.GuestStat
	err := row.Scan(&guest)
	if err != nil {
		return []entitydb.GuestStat{}, nil
	}

	return guest, nil
}

func (gm Guest) FindByUuid(uuid uuid.UUID) (entitydb.GuestStat, error) {
	var hit entitydb.GuestStat
	err := gm.chClient.QueryRow(gm.ctx, `select * from guest where uuid=?`, uuid.String()).Scan(&hit)
	if err != nil {
		return entitydb.GuestStat{}, err
	}
	return hit, nil

}
