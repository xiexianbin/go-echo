package custommiddleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// https://pkg.go.dev/github.com/labstack/echo/v4#readme-official-middleware-repositories
func Init(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// swagger Gzip middleware
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

}
