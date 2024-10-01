package routes

import (
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// HitHandlers Получения данных по хитами посетителей.
type HitHandlers struct {
	fbApp      *fiber.App
	allService *services.AllServices
	ctx        context.Context
}

func NewHit(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) HitHandlers {
	return HitHandlers{
		fbApp:      fbApp,
		allService: allService,
		ctx:        ctx,
	}
}

func (hh HitHandlers) AddHandlers() {
	hh.fbApp.Get("/v1/hit/findAll", hh.findAll)
	hh.fbApp.Post("/v1/hit/filter", hh.filter)
	hh.fbApp.Delete("/v1/hit/:uuid/", hh.findById)
}

func (hh HitHandlers) filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	result, err := hh.allService.Hit.Find(filter)
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

func (hh HitHandlers) findById(ctx *fiber.Ctx) error {
	return nil
}

func (hh HitHandlers) findAll(ctx *fiber.Ctx) error {
	skip, err := strconv.Atoi(ctx.Params("skip", "0"))
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	limit, err := strconv.Atoi(ctx.Params("limit", "0"))
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	allHits, err := hh.allService.Hit.FindAll(uint32(skip), uint32(limit))
	if err != nil {
		return ctx.JSON(entityjson.Response{
			Result: nil,
			Error:  err.Error(),
			Total:  0,
		})
	}
	return ctx.JSON(entityjson.Response{
		Result: hh.allService.Hit.ConvertToJSONListHits(allHits),
		Total:  1,
	})
}

func (hh HitHandlers) name() {

}
