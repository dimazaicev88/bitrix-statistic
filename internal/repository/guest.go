package repository

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/uptrace/go-clickhouse/ch"
)

type Guest struct {
	chClient *ch.DB
}

func NewGuest(chClient *ch.DB) *Guest {
	return &Guest{
		chClient: chClient,
	}
}

func (gm Guest) Add(ctx context.Context, guest models.Guest) error {
	_, err := gm.chClient.NewInsert().Model(&guest).Exec(ctx)
	return err
}

func (gm Guest) FindByHash(ctx context.Context, hash string) (models.Guest, error) {
	var guestDb models.Guest
	err := gm.chClient.NewSelect().Model(&guestDb).Where("guestHash = ?", hash).Scan(ctx)
	return guestDb, err
}
