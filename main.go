package main

import (
	"github.com/carl1330/manga-browser/handlers"
	"github.com/labstack/echo/v4"
)


func main() {
	app := echo.New()

	homeHandler := handlers.Home{}
	settingsHandler := handlers.Settings{}
	mangaHandler := handlers.Manga{}
	app.Static("/static", "assets")
	app.GET("/", homeHandler.HandleHome)
	app.GET("/settings", settingsHandler.HandleSettings)
	app.POST("/api/manga", mangaHandler.PostManga)

	app.Start(":3000")
}