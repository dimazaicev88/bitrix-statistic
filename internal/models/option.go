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

func (o Option) Find(name, siteId string) (entitydb.Option, error) {
	var optionRow entitydb.Option

	err := o.chClient.QueryRow(o.ctx, `select * from options where name=? and  siteId=?`, name, siteId).ScanStruct(&optionRow)
	if err != nil {
		return entitydb.Option{}, err
	}
	return optionRow, nil
}

func (o Option) Set(option entitydb.Option) error {
	var tmpNum int8

	err := o.chClient.QueryRow(o.ctx, `select 1 from options where name=? and  siteId=?;`).Scan(&tmpNum)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return o.Add(option)
	} else if err != nil {
		err := o.chClient.Exec(o.ctx, `ALTER TABLE options UPDATE options SET value=? and description=? where siteId=? and name=?`, option.Value, option.Description, option.SiteId, option.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o Option) Add(option entitydb.Option) error {
	err := o.chClient.Exec(o.ctx, `INSERT INTO options(name,value,description,siteId)VALUES(?, ?, ?, ?);`, option.Name, option.Value, option.Description, option.SiteId)
	if err != nil {
		return err
	}
	return nil
}
