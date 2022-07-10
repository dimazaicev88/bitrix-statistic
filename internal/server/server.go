package server

import (
	"bitrix-statistic/internal/api"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type Server struct {
	app     *fiber.App
	storage *sqlx.DB
}

func NewServer() *Server {
	return &Server{}
}

func (a *Server) Start(port int) error {
	a.app = fiber.New()

	api.NewHitHandlers(a.app).AddHandlers()

	err := a.app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		return err
	}
	return nil
}

func (a *Server) Stop() error {
	err := a.app.Shutdown()
	if err != nil {
		return err
	}
	return nil
}
