package services

import (
	"bitrix-statistic/internal/models"
	"context"
)

type AllService struct {
	Adv        *AdvServices
	Country    *CountryServices
	Event      *EventService
	Guest      *GuestService
	Hit        *HitService
	Option     *OptionService
	Path       *PathService
	Page       *PageService
	PageAdv    *PageAdvService
	Phrase     *PhraseService
	Referer    *RefererService
	Searcher   *SearcherService
	Session    *SessionService
	StatDay    *StatDayService
	Statistic  *StatisticService
	StopList   *StopListService
	Traffic    *TrafficService
	UserOnline *UserOnlineService
}

func NewAllServices(ctx context.Context, allModels *models.Models) *AllService {

	adv := NewAdv(ctx, allModels)
	hit := NewHit(ctx, allModels)
	option := NewOption(ctx, allModels)
	event := NewEvent(ctx, allModels)
	guest := NewGuest(ctx, allModels)
	path := NewPath(ctx, allModels)
	pathCache := NewPathCacheService(ctx, allModels)
	pathAdv := NewPathAdvService(ctx, allModels)
	page := NewPage(ctx, allModels)
	pageAdv := NewPageAdvService(ctx, allModels)
	phrase := NewPhraseService(ctx, allModels)
	referer := NewReferer(ctx, allModels)
	searcher := NewSearcher(ctx, allModels)
	session := NewSession(ctx, allModels)
	statDay := NewStatDay(ctx, allModels)
	statistic := NewStatistic()

	adv.SetHitService(hit)
	adv.SetOptionService(option)

	guest.SetHitService(hit)
	guest.SetAdvService(adv)

	path.SetPathCacheService(pathCache)
	path.SetPathAdvService(pathAdv)
	path.SetOptionService(option)
	path.SetPageService(page)
	path.SetPageAdvService(pageAdv)

	statistic.SetHitService(hit)
	statistic.SetAdvServices(adv)
	statistic.SetGuestService(guest)
	statistic.SetPathService(path)
	statistic.SetSessionService(session)
	statistic.SetStatDayService(statDay)
	statistic.SetSearcherService(searcher)
	statistic.SetOptionService(option)
	statistic.SetRefererService(referer)

	return &AllService{
		Adv:        adv,
		Country:    nil,
		Event:      event,
		Guest:      guest,
		Hit:        hit,
		Option:     option,
		Path:       path,
		Page:       page,
		PageAdv:    pageAdv,
		Phrase:     phrase,
		Referer:    referer,
		Searcher:   searcher,
		Session:    session,
		StatDay:    statDay,
		Statistic:  statistic,
		StopList:   nil,
		Traffic:    nil,
		UserOnline: nil,
	}
}
