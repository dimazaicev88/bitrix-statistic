package server

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"strconv"
)

type Server struct {
	app      *fiber.App
	chClient driver.Conn
	ctx      context.Context
}

func NewServer(ctx context.Context, chClient driver.Conn) *Server {
	return &Server{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (a *Server) Start(port int) error {
	a.app = fiber.New(fiber.Config{
		Views: html.New("./web/html", ".html"),
	})
	a.app.Static("/assets", "./web/assets")

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
