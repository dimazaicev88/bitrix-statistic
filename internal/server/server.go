package server

import (
	"bitrix-statistic/internal/api"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Server struct {
	app     *fiber.App
	storage *storage.MysqlStorage
}

func NewServer(storage *storage.MysqlStorage) *Server {
	return &Server{
		storage: storage,
	}
}

func (a *Server) Start(port int) error {
	a.app = fiber.New()

	api.NewHitHandlers(a.app, models.NewHitModel(a.storage)).
		AddHandlers()

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
