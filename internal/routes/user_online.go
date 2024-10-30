package routes

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserOnline struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllServices
}

func NewUserOnline(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *UserOnline {
	return &UserOnline{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (uo UserOnline) AddHandlers() {
	uo.fbApp.Get("/api/v1/userOnline/filter", uo.filter)
	uo.fbApp.Get("/api/v1/userOnline/count", uo.guestCount)
}

func (uo UserOnline) guestCount(ctx *fiber.Ctx) error {
	return nil
}

func (uo UserOnline) filter(ctx *fiber.Ctx) error {
	return nil
}

func (uo UserOnline) findAll(ctx *fiber.Ctx) error {
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

	allHits, err := uo.allService.UserOnline.FindAll(uint32(skip), uint32(limit))
	if err != nil {
		return ctx.JSON(dto.Response{
			Result: nil,
			Error:  err.Error(),
			Total:  0,
		})
	}
	return ctx.JSON(dto.Response{
		Result: uo.allService.UserOnline.ConvertToJSONListUserOnline(allHits),
		Total:  1,
	})
}
