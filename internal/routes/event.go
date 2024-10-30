package routes

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
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

func (e Event) AddHandlers() {
	e.fbApp.Post("/api/v1/event/filter", e.filterEvent)
	e.fbApp.Delete("/api/v1/event/:uuid/", e.Delete)
	e.fbApp.Get("/api/v1/event/gid/", e.getGid)
	e.fbApp.Post("/api/v1/event/gid/decode", e.getDecodeGid)
	e.fbApp.Post("/api/v1/event/add/", e.Add)
	e.fbApp.Post("/api/v1/event/add/by/events/", e.addByEvents)
	e.fbApp.Post("/api/v1/event/add/current/", e.addCurrent)
	e.fbApp.Post("/api/v1/event/handlers/", e.handlerList)
	e.fbApp.Post("/api/v1/event/filter/by/guests", e.handlerList)

	e.fbApp.Post("/api/v1/event/type/filter", e.filterEventType)
	e.fbApp.Delete("/api/v1/event/type/:uuid/", e.eventTypeDeleteByUuid)
	e.fbApp.Post("/api/v1/event/type/condition/set/", e.conditionSet)
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
func (e Event) filterEvent(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := e.allServices.Adv.GetEventList(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (e Event) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) Add(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) addByEvents(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) addCurrent(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) getGid(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) getDecodeGid(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) handlerList(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) eventTypeDeleteByUuid(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) conditionSet(ctx *fiber.Ctx) error {
	return nil
}

func (e Event) filterEventType(ctx *fiber.Ctx) error {
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
