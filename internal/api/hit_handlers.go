package api

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type HitHandlers struct {
	app      *fiber.App
	hitModel models.HitModel
}

func NewHitHandlers(app *fiber.App, hitModel models.HitModel) HitHandlers {
	return HitHandlers{
		app:      app,
		hitModel: hitModel,
	}
}

func (hh HitHandlers) AddHandlers() {
	hh.app.Post("/hit/filter", hh.Filter)
	hh.app.Delete("/hit/delete/:id/", hh.DeleteById)
}

func (hh HitHandlers) Filter(ctx *fiber.Ctx) error {
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

func (hh HitHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
