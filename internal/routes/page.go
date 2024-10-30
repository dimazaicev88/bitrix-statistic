package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// Page Получения данных о посещенных страницах сайта.
type Page struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewPage(ctx context.Context, app *fiber.App, allService *services.AllServices) *Page {
	return &Page{
		fbApp:       app,
		ctx:         ctx,
		allServices: allService,
	}
}

func (oh Page) AddHandlers() {
	oh.fbApp.Post("/api/v1/page/filter", oh.filter)
	oh.fbApp.Post("/api/v1/page/dynamic/filter", oh.dynamicFilter)
}

func (oh Page) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := oh.allServices.Page.Filter(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (oh Page) dynamicFilter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := oh.allServices.Page.DynamicList(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}
