package main

import (
	"github.com/labstack/echo/v4"
	"github.com/monk3yd/got/internal"
)

func main() {
	port := internal.GetEnv("PORT")

	app := echo.New()
	// userHandler := handler.NewUserHandler()

	app.GET("/", handleRoot)

	app.Start(":" + port)
}

func handleRoot(ctx echo.Context) error {
	return ctx.String(200, "Hello World")
}
