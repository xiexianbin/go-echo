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

package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	v0 "github.com/xiexianbin/go-echo-demo/api/v0"
	_ "github.com/xiexianbin/go-echo-demo/docs"
)

func Init(e *echo.Echo) {

	// Static resource dir
	e.Static("/static", "static")

	// swagger
	url := echoSwagger.URL("swagger.yaml")
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))

	// v0 API for demo
	apiv0 := e.Group("/api/v0")
	{
		apiv0.GET("/helloworld", v0.Helloworld)
		apiv0.GET("/foo", v0.GetFoo1)
		apiv0.GET("/foo/:id", v0.GetFoo2)
		apiv0.POST("/foo", v0.SaveFoo1)
		apiv0.POST("/foo2", v0.SavetFoo2)
	}
}
