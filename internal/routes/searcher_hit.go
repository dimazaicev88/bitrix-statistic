package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// SearcherHit Для получения данных о хитах поисковых систем (про индексированных страниц).
type SearcherHit struct {
	fbApp           *fiber.App
	ctx             context.Context
	searcherService *services.SearcherService
}

func NewSearcherHit(ctx context.Context, fbApp *fiber.App, searcherService *services.SearcherService) *SearcherHit {
	return &SearcherHit{
		fbApp:           fbApp,
		ctx:             ctx,
		searcherService: searcherService,
	}
}

func (p SearcherHit) AddHandlers() {
	p.fbApp.Post("/v1/searcherHit/filter", p.filter)
}

func (p SearcherHit) filter(ctx *fiber.Ctx) error {
	return nil
}
