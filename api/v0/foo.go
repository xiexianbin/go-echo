package v0

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// 请求参数 /foo?id=id1
// g.GET("/foo")
// https://echo.labstack.com/docs/quick-start#query-parameters
func GetFoo1(c echo.Context) error {
	id := c.QueryParam("id")
	return c.String(http.StatusOK, "id: "+id)
}

// URL 路径参数 g.GET("/foo/:id")
// https://echo.labstack.com/docs/quick-start#path-parameters
func GetFoo2(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// 表单 application/form-data
// POST /foo
// curl -F "name=xiexianbin" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/foo
// https://echo.labstack.com/docs/quick-start#form-applicationx-www-form-urlencoded
// https://echo.labstack.com/docs/quick-start#form-multipartform-data
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

type Foo struct {
	Name  string `json,xml,form,query:"name"`
	Email string `json,xml,form,query:"email"`
}

// POST /foo2
func SavetFoo2(c echo.Context) error {
	u := new(Foo)
	if err := c.Bind(u); err != nil {
		return err
	}
	// return c.XML(http.StatusCreated, u)
	return c.JSON(http.StatusCreated, u)
}

func DeleteFoo(c echo.Context) error {
	return nil
}
