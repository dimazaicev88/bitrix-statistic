package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

// получения данных по посетителям сайта.

type GuestRoutes struct {
	fbApp      *fiber.App
	guestModel models.Guest
}

func NewGuestRoutes(fbApp *fiber.App, guestModel models.Guest) GuestRoutes {
	return GuestRoutes{
		fbApp:      fbApp,
		guestModel: guestModel,
	}
}

func (hh GuestRoutes) AddHandlers() {
	hh.fbApp.Post("/v1/guest/filter", hh.filter)
	hh.fbApp.Delete("/v1/guest/delete/:id/", hh.DeleteById)
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
