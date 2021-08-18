package service

import (
	"github.com/Mike-OxHuge/yandex-hackaton/pkg/store"
	"github.com/labstack/echo"
)

type Service struct {
	Store *store.Store
}

func New(store *store.Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) Accesible(c echo.Context) error {
	return c.NoContent(200)
}
