package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// SearcherHit Для получения данных о хитах поисковых систем (про индексированных страниц).
type SearcherHit struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewSearcherHit(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *SearcherHit {
	return &SearcherHit{
		fbApp:       fbApp,
		ctx:         ctx,
		allServices: allService,
	}
}

func (p SearcherHit) AddHandlers() {
	p.fbApp.Post("/api/v1/searcherHit/filter", p.filter)
}

func (p SearcherHit) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := p.allServices.Searcher.FindDynamicList(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}
