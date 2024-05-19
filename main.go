// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/labstack/echo/v4"

	"github.com/xiexianbin/go-echo-demo/pkg/custommiddleware"
	"github.com/xiexianbin/go-echo-demo/router"
)

// @title go-echo-demo Swagger API
// @version 1.0
// @description This is a sample server go-echo demo server.
// @termsOfService https://github.com/xiexianbin/go-echo-demo

// @contact.name API Support
// @contact.url http://www.xiexianbin.cn
// @contact.email me@xiexianbin.cn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:1323
// @BasePath /api
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	custommiddleware.Init(e)

	// Routes
	router.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
