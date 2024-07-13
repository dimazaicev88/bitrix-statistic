package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type CountryHandlers struct {
	fbApp        *fiber.App
	countryModel models.CountryModel
}

func NewCountryHandlers(fbApp *fiber.App, countryModel models.CountryModel) CountryHandlers {
	return CountryHandlers{
		fbApp:        fbApp,
		countryModel: countryModel,
	}
}

func (ch CountryHandlers) AddHandlers() {
	ch.fbApp.Post("/country/filter", ch.Filter)
	ch.fbApp.Post("/country/graph", ch.Filter) //Возвращает данные необходимые для построения графика и круговой диаграммы посещаемости в разрезе по странам.
	ch.fbApp.Delete("/country/delete/:id/", ch.DeleteById)
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
