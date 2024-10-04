package routes

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type CountryHandlers struct {
	fbApp       *fiber.App
	allServices *services.AllServices
	ctx         context.Context
}

func NewCountry(ctx context.Context, fbApp *fiber.App, allServices *services.AllServices) *CountryHandlers {
	return &CountryHandlers{
		fbApp:       fbApp,
		allServices: allServices,
		ctx:         ctx,
	}
}

func (ch CountryHandlers) AddHandlers() {
	ch.fbApp.Post("/api/v1/country/filter", ch.Filter)
	ch.fbApp.Post("/api/v1/country/graph", ch.Filter) //Возвращает данные необходимые для построения графика и круговой диаграммы посещаемости в разрезе по странам.
	ch.fbApp.Delete("/api/v1/country/:uuid/", ch.DeleteById)
}

func (ch CountryHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := ch.allServices.Country.Find(filter)
	if err != nil {
		ctx.Status(502)
		return err
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString(string(resultJson))
}

func (ch CountryHandlers) DeleteById(ctx *fiber.Ctx) error {

	return nil
}
