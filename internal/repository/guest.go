package repository

import (
	"bitrix-statistic/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Guest struct {
	chClient driver.Conn
}

func NewGuest(chClient driver.Conn) *Guest {
	return &Guest{
		chClient: chClient,
	}
}

func (gm Guest) Add(ctx context.Context, guest models.Guest) error {
	return gm.chClient.Exec(
		ctx,
		`INSERT INTO guest(guestHash, dateAdd) VALUES (?,?)`,
		guest.GuestHash, guest.DateAdd,
	)
}

func (gm Guest) FindByHash(ctx context.Context, hash string) (models.Guest, error) {
	var guest models.Guest
	err := gm.chClient.QueryRow(
		ctx, `select guestHash, dateAdd from guest where guestHash=?`,
		hash).ScanStruct(&guest)
	if !errors.Is(err, sql.ErrNoRows) {
		return models.Guest{}, err
	}
	return guest, nil
}
