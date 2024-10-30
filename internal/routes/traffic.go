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
	allService *services.AllServices
}

func NewTraffic(ctx context.Context, fbApp *fiber.App, allService *services.AllServices) *Traffic {
	return &Traffic{
		fbApp:      fbApp,
		ctx:        ctx,
		allService: allService,
	}
}

func (tr Traffic) AddHandlers() {
	tr.fbApp.Post("/api/v1/traffic/filter", tr.filter)
	tr.fbApp.Get("/api/v1/traffic/common", tr.commonValues)
	tr.fbApp.Get("/api/v1/traffic/daily", tr.dailyList)
	tr.fbApp.Get("/api/v1/traffic/phrase", tr.phraseList)
	tr.fbApp.Get("/api/v1/traffic/referer", tr.refererList)
	tr.fbApp.Get("/api/v1/traffic/sum/", tr.sumList)
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
