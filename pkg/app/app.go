package app

import (
	"github.com/Mike-OxHuge/yandex-hackaton/pkg/service"
	"github.com/Mike-OxHuge/yandex-hackaton/pkg/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	store := store.New()
	s := service.New(store)
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/accesible", s.Accesible)
	e.Logger.Fatal(e.Start(":8080"))
}
