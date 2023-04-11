package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type EchoServer struct {
	*echo.Echo
	ctx  context.Context
	db   *gorm.DB
	port string
}

func (es *EchoServer) configure() {
	es.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} : ${method} -> ${uri}, status=${status} ::${error}\n",
		CustomTimeFormat: "15:04:05.00000",
	}))

	es.Use(middleware.Recover())

	es.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
}

func (es *EchoServer) Run() {

	fmt.Println(es.port)
	es.Logger.Fatal(es.Start(":" + es.port))
}

func NewEchoServer(ctx context.Context, db *gorm.DB, appPort string) Server {
	if appPort == "" {
		appPort = "8080"
	}

	server := &EchoServer{
		echo.New(),
		ctx,
		db,
		appPort,
	}

	server.configure()
	server.routes()

	return server
}
