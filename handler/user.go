package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/monk3yd/got/internal"
)

type UserHandler struct {
	db internal.DB
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return nil
}
