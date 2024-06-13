package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type CityHandlers struct {
	fbApp     *fiber.App
	cityModel models.CityModel
}

func NewCityHandlers(fbApp *fiber.App, cityModel models.CityModel) CityHandlers {
	return CityHandlers{
		fbApp:     fbApp,
		cityModel: cityModel,
	}
}

func (ch CityHandlers) AddHandlers() {
	ch.fbApp.Post("/city/filter", ch.Filter)
	ch.fbApp.Delete("/city/delete/:id/", ch.DeleteById)
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
