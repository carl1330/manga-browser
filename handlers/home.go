package handlers

import (
	"github.com/carl1330/manga-browser/view/home"
	"github.com/labstack/echo/v4"
)

type Home struct {

}

func (h Home) HandleHome(c echo.Context) error {
	return render(c, home.Show())
}