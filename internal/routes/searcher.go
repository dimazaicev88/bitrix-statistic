package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Searcher Для работы с поисковыми системами.
type Searcher struct {
	fbApp           *fiber.App
	ctx             context.Context
	searcherService *services.SearcherService
}

func NewSearcher(ctx context.Context, app *fiber.App, searcherService *services.SearcherService) *Searcher {
	return &Searcher{
		fbApp:           app,
		ctx:             ctx,
		searcherService: searcherService,
	}
}

func (p Searcher) AddHandlers() {
	p.fbApp.Post("/v1/searcher/filter", p.filter)
	p.fbApp.Post("/v1/searcher/:uuid", p.findById)
	p.fbApp.Get("/v1/searcher/filterDomainList", p.findDomainList)
	p.fbApp.Get("/v1/searcher/filterDynamicList", p.findDynamicList)
}

func (p Searcher) filter(ctx *fiber.Ctx) error {
	return nil
}

func (p Searcher) findById(ctx *fiber.Ctx) error {
	return nil
}

func (p Searcher) findDomainList(ctx *fiber.Ctx) error {
	return nil
}

func (p Searcher) findDynamicList(ctx *fiber.Ctx) error {
	return nil
}
