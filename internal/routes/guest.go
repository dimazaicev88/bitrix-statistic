package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

// GuestRoutes Получения данных по посетителям сайта.
type GuestRoutes struct {
	fbApp      *fiber.App
	guestModel models.Guest
	ctx        context.Context
}

func NewGuest(fbApp *fiber.App, guestModel models.Guest) GuestRoutes {
	return GuestRoutes{
		fbApp:      fbApp,
		guestModel: guestModel,
	}
}

func (hh GuestRoutes) AddHandlers() {
	hh.fbApp.Post("/v1/guest/filter", hh.filter)
	hh.fbApp.Get("/v1/guest/:uuid", hh.findById)
}

func (hh GuestRoutes) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := hh.guestModel.Find(filter)
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

func (hh GuestRoutes) DeleteById(ctx *fiber.Ctx) error {

	return nil
}

func (hh GuestRoutes) findById(ctx *fiber.Ctx) error {
	return nil
}
