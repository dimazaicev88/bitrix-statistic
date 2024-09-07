package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// SearcherHit Для получения данных о хитах поисковых систем (про индексированных страниц).
type SearcherHit struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllService
}

func NewSearcherHit(ctx context.Context, fbApp *fiber.App, allService *services.AllService) *SearcherHit {
	return &SearcherHit{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (p SearcherHit) AddHandlers() {
	p.fbApp.Post("/v1/searcherHit/filter", p.filter)
}

func (p SearcherHit) filter(ctx *fiber.Ctx) error {
	return nil
}
