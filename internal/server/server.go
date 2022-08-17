package server

import (
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/routes"
	"bitrix-statistic/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"strconv"
)

type Server struct {
	app     *fiber.App
	storage storage.Storage
}

func NewServer(storage storage.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (a *Server) Start(port int) error {
	a.app = fiber.New(fiber.Config{
		Views: html.New("./web/html", ".html"),
	})
	a.app.Static("/assets", "./web/assets")
	routes.NewMainPageHandlers(a.app).AddHandler()
	routes.NewHitHandlers(a.app, models.NewHitModel(a.storage)).AddHandlers()
	routes.NewCityHandlers(a.app, models.NewCityModel(a.storage)).AddHandlers()
	routes.NewSessionHandlers(a.app, models.NewSessionModel(a.storage)).AddHandlers()

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
