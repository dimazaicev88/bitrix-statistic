package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

type UserOnline struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllService
}

func NewUserOnline(ctx context.Context, fbApp *fiber.App, allService *services.AllService) *UserOnline {
	return &UserOnline{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (uo *UserOnline) AddHandlers() {
	uo.fbApp.Get("/v1/userOnline/filter", uo.filter)
	uo.fbApp.Get("/v1/userOnline/guestCount", uo.guestCount)
}

func (uo *UserOnline) guestCount(ctx *fiber.Ctx) error {
	return nil
}

func (uo *UserOnline) filter(ctx *fiber.Ctx) error {
	return nil
}
