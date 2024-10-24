package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type PageAdv struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPageAdv(ctx context.Context, chClient driver.Conn) *PageAdv {
	return &PageAdv{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (pa PageAdv) Add(pageAdv entitydb.PageAdv) error {
	err := pa.chClient.Exec(pa.ctx,
		`INSERT INTO page_adv (dateStat, advUuid, counterBack, enterCounterBack, exitCounterBack)
		       VALUES (curdate(),?,?,?,?)`,
		pageAdv.AdvUuid, pageAdv.CounterBack, pageAdv.EnterCounterBack, pageAdv.ExitCounterBack,
	)
	if err != nil {
		return err
	}
	return nil
}
