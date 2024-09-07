package routes

import (
	"bitrix-statistic/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
)

// Traffic для получения общих данных по посещаемости сайта.
type Traffic struct {
	fbApp      *fiber.App
	ctx        context.Context
	allService *services.AllService
}

func NewTraffic(ctx context.Context, fbApp *fiber.App, allService *services.AllService) *Traffic {
	return &Traffic{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (tr Traffic) AddHandlers() {
	tr.fbApp.Post("/v1/traffic/filter", tr.filter)
	tr.fbApp.Get("/v1/traffic/commonValues", tr.commonValues)
	tr.fbApp.Get("/v1/traffic/dailyList", tr.dailyList)
	tr.fbApp.Get("/v1/traffic/phraseList", tr.phraseList)
	tr.fbApp.Get("/v1/traffic/refererList", tr.refererList)
	tr.fbApp.Get("/v1/traffic/sumList", tr.sumList)
}

func (tr Traffic) filter(ctx *fiber.Ctx) error {
	return nil
}

func (tr Traffic) commonValues(ctx *fiber.Ctx) error {
	return nil
}

func (tr Traffic) dailyList(ctx *fiber.Ctx) error {
	return nil
}

func (tr Traffic) phraseList(ctx *fiber.Ctx) error {
	return nil
}

func (tr Traffic) refererList(ctx *fiber.Ctx) error {
	return nil
}

func (tr Traffic) sumList(ctx *fiber.Ctx) error {
	return nil
}
