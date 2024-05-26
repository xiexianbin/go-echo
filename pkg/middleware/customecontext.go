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

// https://echo.labstack.com/docs/context#extending
package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/xiexianbin/go-echo-demo/pkg/app"
)

// CustomeContextConfig defines the config for CustomeContext middleware.
type CustomeContextConfig struct {
	// Skipper defines a function to skip middleware.
	Skipper echoMiddleware.Skipper
}

// DefaultCustomeContextConfig is the default CustomeContext middleware config.
var DefaultCustomeContextConfig = CustomeContextConfig{
	Skipper: echoMiddleware.DefaultSkipper,
}

// CustomeContext returns a middleware which use CustomeContext.
func CustomeContext() echo.MiddlewareFunc {
	return CustomeContextWithConfig(DefaultCustomeContextConfig)
}

// CustomeContextWithConfig return CustomeContext middleware with config.
// See: `CustomeContext()`.
func CustomeContextWithConfig(config CustomeContextConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultCustomeContextConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := app.CustomeContext{
				Context: c,
			}
			return next(cc)
		}
	}
}
