//go:generate oapi-codegen --config=spec-config/models.yaml "../../api-spec/api.yaml"
//go:generate oapi-codegen --config=spec-config/server.yaml "../../api-spec/api.yaml"
package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type api struct {
}

func NewServer() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/", homePage)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	RegisterHandlers(e, &api{})
	return e
}

func homePage(c echo.Context) error {
	name := "TCP"
	meaning := "Transmission Control Protocol"
	category := "Computing"

	resp := AbbreviationResponse{
		Name:     &name,
		Meaning:  &meaning,
		Category: &category,
	}

	return c.JSON(http.StatusOK, resp)
}

func (a *api) GetAbbreviation(ctx echo.Context, abbr string) error {
	name := abbr
	meaning := "Transmission Control Protocol"
	category := "Computing"

	resp := AbbreviationResponse{Name: &name, Meaning: &meaning, Category: &category}

	return ctx.JSON(http.StatusOK, resp)
}
