package routes

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type SessionHandlers struct {
	app          *fiber.App
	sessionModel models.SessionModel
}

func NewSessionHandlers(app *fiber.App, sessionModel models.SessionModel) SessionHandlers {
	return SessionHandlers{
		app:          app,
		sessionModel: sessionModel,
	}
}

func (sh SessionHandlers) AddHandlers() {
	sh.app.Post("/session/filter", sh.Filter)
	sh.app.Post("/session/add", sh.AddSession)
	sh.app.Delete("/session/delete/:id/", sh.DeleteById)
}

func (sh SessionHandlers) Filter(ctx *fiber.Ctx) error {
	var filter filters.Filter
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &filter)
	if err != nil {
		ctx.Status(502)
		return err
	}
	err, result := sh.sessionModel.Find(filter)
	if err != nil {
		ctx.Status(502)
		return err
	}

	json, err := jsoniter.MarshalToString(result)
	if err != nil {
		ctx.Status(502)
		return err
	}
	return ctx.SendString(json)
}

func (sh SessionHandlers) DeleteById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", -1)
	if err != nil {
		return err
	}
	sh.sessionModel.DeleteById(id)
	return nil
}

func (sh SessionHandlers) AddSession(ctx *fiber.Ctx) error {
	start := time.Now()
	var session entity.SessionJson
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &session)
	if err != nil {
		ctx.Status(502)
		return err
	}
	stop := time.Now().Sub(start)
	fmt.Println(stop)
	//err = sh.sessionModel.AddSession(session)
	//if err != nil {
	//	ctx.Status(502)
	//	return err
	//}

	return nil
}

func (sh SessionHandlers) AddSessionData(ctx *fiber.Ctx) error {
	start := time.Now()
	var session entity.SessionData
	body := ctx.Body()
	err := jsoniter.Unmarshal(body, &session)
	if err != nil {
		ctx.Status(502)
		return err
	}
	stop := time.Now().Sub(start)
	fmt.Println(stop)
	//err = sh.sessionModel.AddSession(session)
	//if err != nil {
	//	ctx.Status(502)
	//	return err
	//}

	return nil
}
