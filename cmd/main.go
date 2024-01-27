package main

import (
	"github.com/labstack/echo/v4"
	"github.com/monk3yd/got/handler"
	"github.com/monk3yd/got/internal"
)

func main() {
	port := internal.GetEnv("PORT")

	app := echo.New()

	app.GET("/", rootHandler)

	userHandler := handler.UserHandler{}
	app.GET("/users", userHandler.HandleUserShow)

	app.Start(":" + port)
}

func rootHandler(ctx echo.Context) error {
	return ctx.String(200, "root")
}
