package repository

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/uptrace/go-clickhouse/ch"
)

type Hit struct {
	chClient *ch.DB
}

func NewHit(chClient *ch.DB) *Hit {
	return &Hit{chClient: chClient}
}

func (hm Hit) AddHit(ctx context.Context, hit models.Hit) error {
	_, err := hm.chClient.NewInsert().Model(&hit).Exec(ctx)
	return err
}
