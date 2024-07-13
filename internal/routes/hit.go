package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

//получения данных по хитами посетителей.

type HitHandlers struct {
	fbApp    *fiber.App
	hitModel models.HitModel
}

func NewHitHandlers(fbApp *fiber.App, hitModel models.HitModel) HitHandlers {
	return HitHandlers{
		fbApp:    fbApp,
		hitModel: hitModel,
	}
}

func (hh HitHandlers) AddHandlers() {
	hh.fbApp.Post("/v1/hit/filter", hh.filter)
	hh.fbApp.Delete("/v1/hit/delete/:id/", hh.DeleteById)
}

func (hh HitHandlers) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := hh.hitModel.Find(filter)
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

func (hh HitHandlers) filterBitrix(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := hh.hitModel.Find2(filter)
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
