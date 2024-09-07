package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type CountryHandlers struct {
	fbApp       *fiber.App
	allServices *services.AllService
	ctx         context.Context
}

func NewCountry(ctx context.Context, fbApp *fiber.App, allServices *services.AllService) *CountryHandlers {
	return &CountryHandlers{
		fbApp:       fbApp,
		allServices: allServices,
		ctx:         ctx,
	}
}

func (ch CountryHandlers) AddHandlers() {
	ch.fbApp.Post("/v1/country/filter", ch.Filter)
	ch.fbApp.Post("/v1/country/graph", ch.Filter) //Возвращает данные необходимые для построения графика и круговой диаграммы посещаемости в разрезе по странам.
	ch.fbApp.Delete("/v1/country/:uuid/", ch.DeleteById)
}

func (ch CountryHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := ch.allServices.Country.Find(filter)
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
