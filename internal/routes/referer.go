package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Referer struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewReferer(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *Referer {
	return &Referer{
		fbApp:       fbApp,
		ctx:         ctx,
		allServices: allService,
	}
}

func (p Referer) AddHandlers() {
	p.fbApp.Post("/api/v1/referer/filter", p.filter)
}

func (p Referer) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := p.allServices.Referer.Find(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}
