package app

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/monforje/dsl-edu-user/internal/mw"
	"github.com/monforje/dsl-edu-user/internal/repository/mongo"
	"github.com/monforje/dsl-edu-user/internal/service"
	"github.com/monforje/dsl-edu-user/internal/transport/http"
	"github.com/monforje/dsl-edu-user/pkg/config"
)

type App struct {
	e    *http.Endpoint
	s    *service.Service
	echo *echo.Echo
	cfg  *config.ConfigDatabase
	db   *mongo.Mongo
}

func New() (*App, error) {
	a := &App{}

	configType := "yaml"

	a.cfg = config.New(configType)

	db, err := mongo.New(a.cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка при инициализации MongoDB: %w", err)
	}
	a.db = db

	a.s = service.New(a.db)

	a.e = http.New(a.s)

	a.echo = echo.New()
	a.echo.HideBanner = true
	a.echo.HidePort = true

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)
	a.echo.POST("/user/exist", a.e.Exist)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("сервер запущен на :8080")

	if err := a.echo.Start(":8080"); err != nil {
		return fmt.Errorf("ошибка при запуске http сервера: %w", err)
	}
	return nil
}
