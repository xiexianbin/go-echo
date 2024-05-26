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

package v0

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/xiexianbin/go-echo-demo/pkg/app"
	"github.com/xiexianbin/go-echo-demo/pkg/ecode"
)

// ref https://github.com/swaggo/swag/blob/master/example/celler/controller/accounts.go

// Foo foo struct for demo
// ref https://github.com/swaggo/swag/blob/master/README.md#user-defined-structure-with-an-array-type
type Foo struct {
	Name  string `json,xml,form,query:"name" example:"account name"`
	Email string `json,xml,form,query:"email" example:"account email"`
}

// Helloworld godoc
//
//	@Summary	helloworld example
//	@Schemes
//	@Description	say helloworld
//	@Tags			foo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	app.Response{data=string}	"Helloworld"
//	@Router			/v0/helloworld [get]
func Helloworld(c echo.Context) error {
	ctx := c.(app.CustomeContext)
	return ctx.WithCodeOK().WithData("helloworld").RJSON()
}

// GetFoo1 example
//
//	@Summary		Read foo by query id
//	@Description	https://echo.labstack.com/docs/quick-start#query-parameters
//	@Tags			foo
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int				true	"Foo Id"
//	@Success		200	{object}	string			"success"
//	@Failure		400	{object}	app.Response	"We need ID!!"
//	@Failure		404	{object}	app.Response	"Can not find ID"
//	@Router			/v0/foo [get]
func GetFoo1(c echo.Context) error {
	id := c.QueryParam("id")
	return c.String(http.StatusOK, "id: "+id)
}

// GetFoo2 example
//
//	@Summary		Read foo by url id
//	@Description	https://echo.labstack.com/docs/quick-start#path-parameters
//	@Tags			foo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int							true	"Foo Id"
//	@Success		200	{object}	app.Response{data=string}	"success"
//	@Failure		400	{object}	app.Response				"We need ID!!"
//	@Failure		404	{object}	app.Response				"Can not find ID"
//	@Router			/v0/foo/{id} [get]
func GetFoo2(c echo.Context) error {
	ctx := c.(app.CustomeContext)
	id := ctx.Param("id")
	return ctx.WithCode(ecode.CodeOK).WithData(id).RJSON()
}

// SaveFoo1 example
//
//	@Summary		save foo by application/form-data
//	@Description	https://echo.labstack.com/docs/quick-start#form-applicationx-www-form-urlencoded
//	@Description	https://echo.labstack.com/docs/quick-start#form-multipartform-data
//	@Description	curl -F "name=xiexianbin" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/foo
//	@Tags			foo
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string						true	"Foo Id"
//	@Param			avatar	formData	file						true	"Foo Id"
//	@Success		200		{object}	app.Response{data=string}	"success"
//	@Failure		400		{object}	app.Response				"We need ID!!"
//	@Failure		404		{object}	app.Response				"Can not find ID"
//	@Router			/v0/foo [post]
func SaveFoo1(c echo.Context) error {
	name := c.FormValue("name")

	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	return c.HTML(http.StatusOK, "<b>name: "+name+"</b>")
}

// SavetFoo2 example
//
//	@Summary		save foo by json
//	@Description	...
//	@Tags			foo
//	@Accept			json
//	@Produce		json
//	@Param			foo	body		Foo				true	"Foo Data"
//	@Success		200	{object}	string			"success"
//	@Failure		400	{object}	app.Response	"We need ID!!"
//	@Failure		404	{object}	app.Response	"Can not find ID"
//	@Router			/v0/foo2 [post]
func SavetFoo2(c echo.Context) error {
	f := new(Foo)
	if err := c.Bind(f); err != nil {
		return err
	}
	// return c.XML(http.StatusCreated, u)
	return c.JSON(http.StatusCreated, f)
}

// DeleteFoo example
//
//	@Summary		delete foo by path id
//	@Description	...
//	@Tags			foo
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int				true	"Foo Id"
//	@Success		200	{object}	app.Response	"success"
//	@Failure		400	{object}	app.Response	"We need ID!!"
//	@Failure		404	{object}	app.Response	"Can not find ID"
//	@Router			/v0/foo/{id} [get]
func DeleteFoo(c echo.Context) error {
	return nil
}
