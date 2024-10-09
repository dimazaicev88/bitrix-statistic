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

func (gm Guest) Add(guest entitydb.Guest) error {
	return gm.chClient.Exec(gm.ctx,
		`INSERT INTO guest (uuid, date_add, favorites,repair) VALUES (?,?,?,?)`,
		guest.Uuid, guest.DateAdd, guest.Favorites, guest.Repair,
	)
}

func (gm Guest) Find(filter filters.Filter) ([]entitydb.Guest, error) {
	return []entitydb.Guest{}, nil
}

func (gm Guest) FindByUuid(uuid uuid.UUID) (entitydb.Guest, error) {
	var hit entitydb.Guest
	err := gm.chClient.QueryRow(gm.ctx, `select uuid, date_add, favorites, repair from guest where uuid=?`,
		uuid).ScanStruct(&hit)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.Guest{}, err
	}
	return hit, nil
}
