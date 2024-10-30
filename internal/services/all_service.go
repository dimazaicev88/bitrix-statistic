package services

import (
	"bitrix-statistic/internal/models"
	"context"
)

type AllServices struct {
	Event     *EventService
	Guest     *GuestService
	Hit       *HitService
	Session   *SessionService
	Statistic *Statistic
	StopList  *StopListService
}

func NewAllServices(ctx context.Context, allModels *models.Models) *AllServices {

	hit := NewHit(ctx, allModels)
	event := NewEvent(ctx, allModels)
	guest := NewGuest(ctx, allModels)
	session := NewSession(ctx, allModels)
	statistic := NewStatistic()

	guest.SetHitService(hit)

	statistic.SetHitService(hit)
	statistic.SetGuestService(guest)
	statistic.SetSessionService(session)

	return &AllServices{
		Event:     event,
		Guest:     guest,
		Hit:       hit,
		Session:   session,
		Statistic: statistic,
		StopList:  nil,
	}
}
