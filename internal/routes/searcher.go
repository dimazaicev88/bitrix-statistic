package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// Searcher Для работы с поисковыми системами.
type Searcher struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewSearcher(ctx context.Context, app *fiber.App, allService *services.AllServices) *Searcher {
	return &Searcher{
		fbApp:       app,
		ctx:         ctx,
		allServices: allService,
	}
}

func (p Searcher) AddHandlers() {
	p.fbApp.Post("/api/v1/searcher/filter", p.filter)
	p.fbApp.Post("/api/v1/searcher/:uuid", p.findById)
	p.fbApp.Get("/api/v1/searcher/domain/filter/", p.findDomainList)
	p.fbApp.Get("/api/v1/searcher/dynamic/filter", p.findDynamicList)
}

func (p Searcher) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := p.allServices.Searcher.Find(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (p Searcher) findById(ctx *fiber.Ctx) error {
	return nil
}

func (p Searcher) findDomainList(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := p.allServices.Searcher.FindDomainList(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (p Searcher) findDynamicList(ctx *fiber.Ctx) error {
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
