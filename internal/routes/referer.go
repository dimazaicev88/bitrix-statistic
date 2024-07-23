package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Referer struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewReferer(fbApp *fiber.App, ctx context.Context) *Referer {
	return &Referer{
		fbApp: fbApp,
		ctx:   ctx,
	}
}

func (p Referer) AddHandlers() {
	p.fbApp.Post("/v1/referer/filter", p.filter)
}

func (p Referer) filter(ctx *fiber.Ctx) error {
	return nil
}
