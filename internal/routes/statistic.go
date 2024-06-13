package routes

import (
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
)

type StatisticHandlers struct {
	app         *fiber.App
	optionModel models.StatisticModel
}

func NewStatisticRoutes(app *fiber.App, optionModel models.OptionModel) OptionHandlers {
	return OptionHandlers{
		app:         app,
		optionModel: optionModel,
	}
}

func (sh StatisticHandlers) RegRoutes() {
	sh.app.Post("/statistic/add", sh.Add)
}

func (sh StatisticHandlers) Add(ctx *fiber.Ctx) error {
	//var runImport entity.StatData
	//if err := ctx.BodyParser(&runImport); err != nil {
	//	return c.JSON(entity.AnswerResult{
	//		Err: err.Error(),
	//	})
	//}
	return nil
}
