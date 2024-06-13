package routes

import (
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
)

type Statistic struct {
	fbApp          *fiber.App
	statisticModel models.StatisticModel
}

func NewStatistic(fbApp *fiber.App) Statistic {
	return Statistic{
		fbApp: fbApp,
	}
}

func (sh Statistic) RegRoutes() {
	sh.fbApp.Post("/statistic/add", sh.Add)
}

func (sh Statistic) Add(ctx *fiber.Ctx) error {
	//var runImport entity.StatData
	//if err := ctx.BodyParser(&runImport); err != nil {
	//	return c.JSON(entity.AnswerResult{
	//		Err: err.Error(),
	//	})
	//}
	return nil
}
