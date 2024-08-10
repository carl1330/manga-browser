package handlers

import (
	"github.com/carl1330/manga-browser/view/settings"
	"github.com/labstack/echo/v4"
)

type Settings struct {

}

func (s Settings) HandleSettings(c echo.Context) error {
	return render(c, settings.Show())
}

