package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

type Referer struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllServices
}

func NewReferer(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *Referer {
	return &Referer{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (p Referer) AddHandlers() {
	p.fbApp.Post("/v1/referer/filter", p.filter)
}

func (p Referer) filter(ctx *fiber.Ctx) error {
	return nil
}
