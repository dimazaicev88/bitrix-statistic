package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type UserOnline struct {
	fbApp *fiber.App
	ctx   context.Context
}

func NewUserOnline(fbApp *fiber.App, ctx context.Context) *UserOnline {
	return &UserOnline{
		fbApp: fbApp,
		ctx:   ctx,
	}
}

func (uo *UserOnline) RegRoutes() {
	uo.fbApp.Get("/v1/userOnline/filter", uo.filter)
	uo.fbApp.Get("/v1/userOnline/guestCount", uo.guestCount)
}

func (uo *UserOnline) guestCount(ctx *fiber.Ctx) error {
	return nil
}

func (uo *UserOnline) filter(ctx *fiber.Ctx) error {
	return nil
}
