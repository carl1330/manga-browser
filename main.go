package main

import (
	"github.com/carl1330/manga-browser/handlers"
	"github.com/labstack/echo/v4"
)


func main() {
	app := echo.New()

	homeHandler := handlers.Home{}
	settingsHandler := handlers.Settings{}
	app.Static("/static", "assets")
	app.GET("/", homeHandler.HandleHome)
	app.GET("/settings", settingsHandler.HandleSettings)

	app.Start(":3000")
}