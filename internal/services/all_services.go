package services

import (
	"bitrix-statistic/internal/models"
	"context"
)

type AllServices struct {
	advServices     *AdvServices
	guestService    *GuestService
	sessionService  *SessionService
	statDayService  *StatDayService
	searcherService *SearcherService
	optionService   *OptionService
	hitService      *HitService
	refererService  *RefererService
}

func NewServices(ctx context.Context, allModels *models.Models) *AllServices {

	sessionService := NewSession(ctx, allModels),
		statDayService:  NewStatDay(ctx, allModels),
		searcherService: NewSearcher(ctx, allModels),
		hitService:      hitService,
		optionService:   NewOption(ctx, allModels),
		refererService:  NewReferer(ctx, allModels),

	return &AllServices{}
}
