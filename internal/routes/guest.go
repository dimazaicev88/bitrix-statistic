package routes

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GuestRoutes Получения данных по посетителям сайта.
type GuestRoutes struct {
	fbApp       *fiber.App
	allServices *services.AllServices
	ctx         context.Context
}

func NewGuest(ctx context.Context, fbApp *fiber.App, allServices *services.AllServices) GuestRoutes {
	return GuestRoutes{
		fbApp:       fbApp,
		allServices: allServices,
		ctx:         ctx,
	}
}

func (hh GuestRoutes) AddHandlers() {
	hh.fbApp.Post("/api/v1/guest/filter", hh.filter)
	hh.fbApp.Get("/api/v1/guest/:uuid", hh.findById)
}

func (hh GuestRoutes) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := hh.allServices.Guest.Find(filter)
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

func (hh GuestRoutes) DeleteById(ctx *fiber.Ctx) error {

	return nil
}

func (hh GuestRoutes) findById(ctx *fiber.Ctx) error {
	return nil
}

func (hh GuestRoutes) findAll(ctx *fiber.Ctx) error {
	skip, err := strconv.Atoi(ctx.Params("skip", "0"))
	if err != nil {
		return ctx.JSON(map[string]any{
			"error": err.Error(),
		})
	}
	limit, err := strconv.Atoi(ctx.Params("limit", "0"))
	if err != nil {
		return ctx.JSON(map[string]any{
			"error": err.Error(),
		})
	}

	allHits, err := hh.allServices.Guest.FindAll(uint32(skip), uint32(limit))
	if err != nil {
		return ctx.JSON(dto.Response{
			Result: nil,
			Error:  err.Error(),
			Total:  0,
		})
	}
	return ctx.JSON(dto.Response{
		Result: hh.allServices.Guest.ConvertToJSONListGuest(allHits),
		Total:  1,
	})
}
