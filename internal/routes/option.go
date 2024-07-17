package routes

import (
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
)

type OptionHandlers struct {
	app         *fiber.App
	optionModel models.Option
}

func NewOptionHandlers(app *fiber.App, optionModel models.Option) OptionHandlers {
	return OptionHandlers{
		app:         app,
		optionModel: optionModel,
	}
}

func (hh OptionHandlers) AddHandlers() {
	hh.app.Post("/option/AddSession", hh.Add)
	hh.app.Post("/option/list", hh.List)
	hh.app.Delete("/option/delete/:id/", hh.DeleteById)
}

func (hh OptionHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}

func (hh OptionHandlers) Add(ctx *fiber.Ctx) error {
	return nil
}

func (hh OptionHandlers) List(ctx *fiber.Ctx) error {
	return nil
}
