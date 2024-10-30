package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Models struct {
	Browser    *Browser
	Event      *Event
	Guest      *Guest
	Hit        *Hit
	Session    *Session
	StopList   *StopList
	Traffic    *Traffic
	UserOnline *UserOnline
}

func NewModels(ctx context.Context, chClient driver.Conn) *Models {
	return &Models{
		Browser:    NewBrowser(ctx, chClient),
		Event:      NewEvent(ctx, chClient),
		Guest:      NewGuest(ctx, chClient),
		Hit:        NewHit(ctx, chClient),
		Session:    NewSession(ctx, chClient),
		StopList:   NewStopList(ctx, chClient),
		Traffic:    NewTraffic(ctx, chClient),
		UserOnline: NewUserOnline(ctx, chClient),
	}
}
