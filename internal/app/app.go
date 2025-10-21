package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	log.Println("сервер запущен на :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := a.echo.Start(":8080"); err != nil {
			log.Println("HTTP сервер остановлен...")
			log.Println(err)
		}
	}()

	<-quit
	log.Println("завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.db.Close(ctx); err != nil {
		log.Printf("ошибка при отключении Mongo: %v", err)
	}

	if err := a.echo.Shutdown(ctx); err != nil {
		log.Printf("ошибка при завершении Echo: %v", err)
	}

	log.Println("сервер корректно завершил работу")
	return nil
}
