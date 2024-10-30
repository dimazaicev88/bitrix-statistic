package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Session struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewSession(ctx context.Context, chClient driver.Conn) *Session {
	return &Session{ctx: ctx, chClient: chClient}
}

func (s Session) Add(session entitydb.Session) error {
	return s.chClient.Exec(s.ctx, `INSERT INTO session (uuid, guestUuid, dateAdd, phpSessionId) VALUES (?,?,?,?)`,
		session.Uuid, session.GuestUuid, session.PhpSessionId)
}

func (s Session) ExistsByPhpSession(session string) (bool, error) {
	var count uint8
	row := s.chClient.QueryRow(s.ctx, `select 1 as cnt from session where phpSessionId=?`, session)

	err := row.Scan(&count)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return false, err
	}
	return count > 0, nil
}

func (s Session) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	var sessionDb entitydb.Session
	err := s.chClient.QueryRow(s.ctx, `select * from session where phpSessionId=?`, phpSessionId).ScanStruct(&sessionDb)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.Session{}, err
	}
	return sessionDb, nil
}
