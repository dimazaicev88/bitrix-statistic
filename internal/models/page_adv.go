package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"database/sql"
	"errors"
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
		`INSERT INTO page_adv (date_stat, adv_uuid, counter_back, enter_counter_back, exit_counter_back)
		       VALUES (curdate(),?,?,?,?)`,
		pageAdv.AdvUuid, pageAdv.CounterBack, pageAdv.EnterCounterBack, pageAdv.ExitCounterBack,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return err
	}
	return nil
}
