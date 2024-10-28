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

// Event Для работы с событиями.
type Event struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewEvent(ctx context.Context, app *fiber.App, allServices *services.AllServices) *Event {
	return &Event{
		ctx:         ctx,
		fbApp:       app,
		allServices: allServices,
	}
}

// TODO добавить методы AddByEvents AddCurrent DecodeGID GetGID GetListByGuest
func (e Event) AddHandlers() {
	e.fbApp.Post("/api/v1/event/filter", e.Filter)
	e.fbApp.Delete("/api/v1/event/:uuid/", e.DeleteById)
	e.fbApp.Post("/api/v1/event/set/", e.Add)
}

func (e Event) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := e.allServices.Event.Find(filter) //TODO добавить парсинг
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (e Event) DeleteById(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) Add(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) findAll(ctx *fiber.Ctx) error {
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

	allHits, err := e.allServices.Event.FindAll(uint32(skip), uint32(limit))
	if err != nil {
		return ctx.JSON(dto.Response{
			Result: nil,
			Error:  err.Error(),
			Total:  0,
		})
	}
	return ctx.JSON(dto.Response{
		Result: e.allServices.Event.ConvertToJSONListEvents(allHits),
		Total:  1,
	})
}
