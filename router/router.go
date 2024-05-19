package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	v0 "github.com/xiexianbin/go-echo-demo/api/v0"
)

func Init(e *echo.Echo) {

	// Static resource dir
	e.Static("/static", "static")

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// v0 API for demo
	gv0 := e.Group("/v0")
	{
		gv0.GET("/foo/", v0.GetFoo1)
		gv0.GET("/foo/:id", v0.GetFoo2)
		gv0.POST("/foo", v0.SaveFoo1)
		gv0.POST("/foo2", v0.SavetFoo2)
	}

}
