package services

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StatDayService struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewStatDayService(ctx context.Context, chClient driver.Conn) *StatDayService {
	return &StatDayService{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (sds StatDayService) Update() {

}
