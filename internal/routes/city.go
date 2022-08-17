package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type CityHandlers struct {
	app       *fiber.App
	cityModel models.CityModel
}

func NewCityHandlers(app *fiber.App, cityModel models.CityModel) CityHandlers {
	return CityHandlers{
		app:       app,
		cityModel: cityModel,
	}
}

func (ch CityHandlers) AddHandlers() {
	ch.app.Post("/city/filter", ch.Filter)
	ch.app.Delete("/city/delete/:id/", ch.DeleteById)
}

func (ch CityHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := ch.cityModel.Find(filter)
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

func (ch CityHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
