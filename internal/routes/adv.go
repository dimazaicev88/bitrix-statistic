package routes

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
)

type AdvHandlers struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllServices
}

func NewAdv(ctx context.Context, fbApp *fiber.App, allServices *services.AllServices) *AdvHandlers {
	return &AdvHandlers{
		fbApp:       fbApp,
		ctx:         ctx,
		allServices: allServices,
	}
}

func (ah AdvHandlers) AddHandlers() {
	ah.fbApp.Post("/api/v1/adv/filter", ah.filter)
	ah.fbApp.Post("/api/v1/adv/dynamic/filter", ah.filterDynamic)
	ah.fbApp.Post("/api/v1/adv/simple/filter", ah.filterSimple)

	ah.fbApp.Get("/api/v1/adv/:uuid/", ah.findByUuid)
	ah.fbApp.Post("/api/v1/adv/event/filter", ah.filterEvent)
	ah.fbApp.Delete("/api/v1/adv/delete/:uuid/", ah.deleteByUuid)
}

func (ah AdvHandlers) filter(ctx *fiber.Ctx) error {
	return nil
}

func (ah AdvHandlers) deleteByUuid(ctx *fiber.Ctx) error {
	advUuid := ctx.Params("uuid", "")
	if len(advUuid) > 0 {
		bytes, err := uuid.FromBytes([]byte(advUuid))
		if err = ah.allServices.Adv.Delete(bytes); err != nil {
			return err
		}
	}
	return nil
}

func (ah AdvHandlers) filterEvent(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := ah.allServices.Adv.GetEventList(filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (ah AdvHandlers) findByUuid(ctx *fiber.Ctx) error {
	advUuid := ctx.Params("uuid", "")
	if len(advUuid) > 0 {
		bytes, err := uuid.FromBytes([]byte(advUuid))
		if err != nil {
			return err
		}
		adv, err := ah.allServices.Adv.FindByUuid(bytes)
		if err != nil {
			return err
		}

		return ctx.JSON(adv)
	}
	return nil
}

func (ah AdvHandlers) filterDynamic(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := ah.allServices.Adv.GetDynamicList(filter, true) //TODO добавить парсинг
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}

func (ah AdvHandlers) findAll(ctx *fiber.Ctx) error {
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

	allHits, err := ah.allServices.Adv.FindAll(uint32(skip), uint32(limit))
	if err != nil {
		return ctx.JSON(dto.Response{
			Result: nil,
			Error:  err.Error(),
			Total:  0,
		})
	}
	return ctx.JSON(dto.Response{
		Result: ah.allServices.Adv.ConvertToJSONListAdv(allHits),
		Total:  1,
	})
}

func (ah AdvHandlers) filterSimple(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := ah.allServices.Adv.GetSimpleList(filter) //TODO добавить парсинг
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.SendString(string(resultJson))
}
