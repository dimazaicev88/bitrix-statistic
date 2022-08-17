package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type CountryHandlers struct {
	app          *fiber.App
	countryModel models.CountryModel
}

func NewCountryHandlers(app *fiber.App, countryModel models.CountryModel) CountryHandlers {
	return CountryHandlers{
		app:          app,
		countryModel: countryModel,
	}
}

func (ch CountryHandlers) AddHandlers() {
	ch.app.Post("/country/filter", ch.Filter)
	ch.app.Delete("/country/delete/:id/", ch.DeleteById)
}

func (ch CountryHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := ch.countryModel.Find(filter)
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

func (ch CountryHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
