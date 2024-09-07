package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdvHandlers struct {
	fbApp       *fiber.App
	ctx         context.Context
	allServices *services.AllService
}

func NewAdv(ctx context.Context, fbApp *fiber.App, allServices *services.AllService) *AdvHandlers {
	return &AdvHandlers{
		fbApp:       fbApp,
		ctx:         ctx,
		allServices: allServices,
	}
}

func (ah AdvHandlers) AddHandlers() {
	ah.fbApp.Post("/v1/adv/filter", ah.Filter)
	ah.fbApp.Get("/v1/adv/:uuid/", ah.FindByUuid)
	ah.fbApp.Post("/v1/adv/event/filter", ah.FilterEvent)
	ah.fbApp.Delete("/v1/adv/delete/:uuid/", ah.DeleteByUuid)
}

func (ah AdvHandlers) Filter(ctx *fiber.Ctx) error {
	return nil
}

func (ah AdvHandlers) DeleteByUuid(ctx *fiber.Ctx) error {
	advUuid := ctx.Params("uuid", "")
	if len(advUuid) > 0 {
		bytes, err := uuid.FromBytes([]byte(advUuid))
		if err = ah.allServices.Adv.DeleteByUuid(bytes); err != nil {
			return err
		}
	}
	return nil
}

func (ah AdvHandlers) FilterEvent(ctx *fiber.Ctx) error {
	return nil
}

func (ah AdvHandlers) FindByUuid(ctx *fiber.Ctx) error {
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
