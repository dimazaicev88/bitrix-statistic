package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

// HitHandlers Получения данных по хитами посетителей.
type HitHandlers struct {
	fbApp      *fiber.App
	allService *services.AllService
	ctx        context.Context
}

func NewHit(ctx context.Context, fbApp *fiber.App, allService *services.AllService) HitHandlers {
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
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := hh.allService.Hit.Find(filter)
	if err != nil {
		ctx.Status(502)
		return err
	}

	json, err := jsoniter.MarshalToString(result)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString(json)
}

func (hh HitHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}

func (hh HitHandlers) findById(ctx *fiber.Ctx) error {
	return nil
}
