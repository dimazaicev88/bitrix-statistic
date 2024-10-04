package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Option struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewOption(ctx context.Context, chClient driver.Conn) *Option {
	return &Option{ctx: ctx, chClient: chClient}
}

func (o Option) Find(name string) (entitydb.Option, error) {
	var optionRow entitydb.Option

	err := o.chClient.QueryRow(o.ctx, `select * from options where name=? `, name).ScanStruct(&optionRow)
	if err != nil {
		return entitydb.Option{}, err
	}
	return optionRow, nil
}

func (o Option) Set(option entitydb.Option) error {
	var tmpNum int8

	err := o.chClient.QueryRow(o.ctx, `select 1 from options where name=?`, option.Name).Scan(&tmpNum)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return o.Add(option)
	} else if err != nil {
		err := o.chClient.Exec(o.ctx, `ALTER TABLE options UPDATE value = ? WHERE name = ?`, option.Value, option.Name)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (o Option) Add(option entitydb.Option) error {
	err := o.chClient.Exec(o.ctx, `INSERT INTO options(name,value)VALUES(?, ?);`, option.Name, option.Value)
	if err != nil {
		return err
	}
	return nil
}
