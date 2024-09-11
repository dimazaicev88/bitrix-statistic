package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// GuestRoutes Получения данных по посетителям сайта.
type GuestRoutes struct {
	fbApp      *fiber.App
	allService *services.AllServices
	ctx        context.Context
}

func NewGuest(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) GuestRoutes {
	return GuestRoutes{
		fbApp:      fbApp,
		allService: allService,
		ctx:        ctx,
	}
}

func (hh GuestRoutes) AddHandlers() {
	hh.fbApp.Post("/v1/guest/filter", hh.filter)
	hh.fbApp.Get("/v1/guest/:uuid", hh.findById)
}

func (hh GuestRoutes) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := hh.allService.Guest.Find(filter)
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

func (hh GuestRoutes) DeleteById(ctx *fiber.Ctx) error {

	return nil
}

func (hh GuestRoutes) findById(ctx *fiber.Ctx) error {
	return nil
}
