// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var UseMiddlewars = []func() echo.MiddlewareFunc{
	CustomeContext, // This middleware should be registered before any other middleware.
	echoMiddleware.Logger,
	echoMiddleware.Recover,
}

// https://pkg.go.dev/github.com/labstack/echo/v4#readme-official-middleware-repositories
func Init(e *echo.Echo) {
	for _, middleware := range UseMiddlewars {
		e.Use(middleware())
	}

	// swagger Gzip middleware
	e.Use(echoMiddleware.GzipWithConfig(echoMiddleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

}
