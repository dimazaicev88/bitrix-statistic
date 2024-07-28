package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Models struct {
	AdvModel   *Adv
	Browser    *Browser
	City       *City
	Country    *Country
	Day        *Day
	Event      *Event
	Guest      *Guest
	Hit        *Hit
	Option     *Option
	Page       *Page
	Path       *Path
	Phrase     *Phrase
	Referer    *Referer
	Searcher   *Searcher
	Session    *Session
	StopList   *StopList
	Traffic    *Traffic
	UserOnline *UserOnline
}

func NewModels(ctx context.Context, chClient driver.Conn) *Models {
	return &Models{
		AdvModel:   NewAdv(ctx, chClient),
		Browser:    NewBrowser(ctx, chClient),
		City:       NewCity(ctx, chClient),
		Country:    NewCountry(ctx, chClient),
		Day:        NewDay(ctx, chClient),
		Event:      NewEvent(ctx, chClient),
		Guest:      NewGuest(ctx, chClient),
		Hit:        NewHit(ctx, chClient),
		Option:     NewOption(ctx, chClient),
		Page:       NewPage(ctx, chClient),
		Path:       NewPath(ctx, chClient),
		Phrase:     NewPhrase(ctx, chClient),
		Referer:    NewReferer(ctx, chClient),
		Searcher:   NewSearcher(ctx, chClient),
		Session:    NewSession(ctx, chClient),
		StopList:   NewStopList(ctx, chClient),
		Traffic:    NewTraffic(ctx, chClient),
		UserOnline: NewUserOnline(ctx, chClient),
	}
}
