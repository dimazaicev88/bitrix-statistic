package api

import (
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
)

type StatisticHandlers struct {
	app         *fiber.App
	optionModel models.StatisticModel
}

func NewStatisticHandlers(app *fiber.App, optionModel models.OptionModel) OptionHandlers {
	return OptionHandlers{
		app:         app,
		optionModel: optionModel,
	}
}

func (sh StatisticHandlers) AddHandlers() {
	sh.app.Post("/statistic/add", sh.Add)
}

func (sh StatisticHandlers) Add(ctx *fiber.Ctx) error {
	return nil
}
