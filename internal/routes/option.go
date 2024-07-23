package routes

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/gofiber/fiber/v2"
)

type OptionHandlers struct {
	app         *fiber.App
	optionModel models.Option
	ctx         context.Context
}

func NewOptionHandlers(app *fiber.App, optionModel models.Option) OptionHandlers {
	return OptionHandlers{
		app:         app,
		optionModel: optionModel,
	}
}

func (oh OptionHandlers) AddHandlers() {
	oh.app.Post("/v1/option/add", oh.add)
	oh.app.Get("/v1/option/list", oh.list)
	oh.app.Delete("/v1/option/:uuid/", oh.deleteByUuid)
}

func (oh OptionHandlers) deleteByUuid(ctx *fiber.Ctx) error {

	return nil
}

func (oh OptionHandlers) add(ctx *fiber.Ctx) error {
	return nil
}

func (oh OptionHandlers) list(ctx *fiber.Ctx) error {
	return nil
}
