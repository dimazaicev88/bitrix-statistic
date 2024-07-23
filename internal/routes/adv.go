package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

type AdvHandlers struct {
	fbApp      *fiber.App
	ctx        context.Context
	advService services.AdvServices
}

func NewAdvHandlers(advService services.AdvServices, fbApp *fiber.App, ctx context.Context) *AdvHandlers {
	return &AdvHandlers{
		fbApp:      fbApp,
		ctx:        ctx,
		advService: advService,
	}
}

func (ah AdvHandlers) AddHandlers() {
	ah.fbApp.Post("/adv/filter", ah.Filter)
	ah.fbApp.Get("/adv/:uuid/", ah.FindByUuid)
	ah.fbApp.Post("/adv/event/filter", ah.FilterEvent)
	ah.fbApp.Delete("/adv/delete/:uuid/", ah.DeleteByUuid)
}

func (ah AdvHandlers) Filter(ctx *fiber.Ctx) error {
	return nil
}

func (ah AdvHandlers) DeleteByUuid(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")
	if len(uuid) > 0 {
		ah.advService.DeleteByUuid(uuid)
	}
	return nil
}

func (ah AdvHandlers) FilterEvent(ctx *fiber.Ctx) error {
	return nil
}

func (ah AdvHandlers) FindByUuid(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")
	if len(uuid) > 0 {
		adv, err := ah.advService.FindByUuid(uuid)
		if err != nil {
			return err
		}

		return ctx.JSON(adv)
	}
	return nil
}
