package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// HitHandlers Получения данных по хитами посетителей.
type HitHandlers struct {
	fbApp      *fiber.App
	allService *services.AllServices
	ctx        context.Context
}

func NewHit(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) HitHandlers {
	return HitHandlers{
		fbApp:      fbApp,
		allService: allService,
		ctx:        ctx,
	}
}

func (hh HitHandlers) AddHandlers() {
	hh.fbApp.Post("/v1/hit/filter", hh.filter)
	hh.fbApp.Delete("/v1/hit/:uuid/", hh.findById)
}

func (hh HitHandlers) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := hh.allService.Hit.Find(filter)
	if err != nil {
		ctx.Status(502)
		return err
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString(string(resultJson))
}

func (hh HitHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}

func (hh HitHandlers) findById(ctx *fiber.Ctx) error {
	return nil
}
