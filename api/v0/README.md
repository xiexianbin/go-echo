v0 for study

## route

- https://echo.labstack.com/docs/quick-start#routing

```
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```

## Content

```
// https://echo.labstack.com/docs/context
package v0

import "github.com/labstack/echo"

// Define a custom context
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

// Create a middleware to extend default context
// e := echo.New()
// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		cc := &CustomContext{c}
// 		return next(cc)
// 	}
// })

// Use in handler
// e.GET("/", func(c echo.Context) error {
// 	cc := c.(*CustomContext)
// 	cc.Foo()
// 	cc.Bar()
// 	return cc.String(200, "OK")
// })
```

## Customization

- https://echo.labstack.com/docs/customization
  - debug/logger and so on

## Cookies

- https://echo.labstack.com/docs/cookies

## Error Handling

- https://echo.labstack.com/docs/error-handling

## Start Server

- https://echo.labstack.com/docs/start-server

## IP Address

- [With proxies using X-Forwarded-For header](https://echo.labstack.com/docs/ip-address#case-2-with-proxies-using-x-forwarded-for-header)

## Request

- https://echo.labstack.com/docs/request
- [Validate Data](https://echo.labstack.com/docs/request#validate-data)
  - https://github.com/go-playground/validator

## Response

- https://echo.labstack.com/docs/response
- Hook
  - Before Response
  - After Response

## Routing

- https://echo.labstack.com/docs/routing

## Static Files

- https://echo.labstack.com/docs/static-files

## Templates

- https://echo.labstack.com/docs/templates

## Testing

- https://echo.labstack.com/docs/testing

## Middleware

- https://echo.labstack.com/docs/quick-start#middleware
- https://echo.labstack.com/docs/category/middleware
  - [Casbin Auth](https://echo.labstack.com/docs/middleware/casbin-auth)
  - [Jaeger](https://echo.labstack.com/docs/middleware/jaeger)
  - [Logger](https://echo.labstack.com/docs/middleware/logger#examples) for zap
  - [Prometheus](https://echo.labstack.com/docs/middleware/prometheus)
  - [Rate Limiter](https://echo.labstack.com/docs/middleware/rate-limiter)
  - [Recover](https://echo.labstack.com/docs/middleware/recover)
  - [Request ID](https://echo.labstack.com/docs/middleware/request-id)
  - [Session](https://echo.labstack.com/docs/middleware/session)
  - [Trailing Slash](https://echo.labstack.com/docs/middleware/trailing-slash)

```
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Group level middleware
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string, c echo.Context)(error, bool) {
  if username == "u1" && password == "pwd" {
    return nil, true
  }
  rturn nil, false
}))

// Route level middleware
track := func(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // ...
    return next(c)
  }
}
e.GET("/users", func(c echo.Context) error {
  return c.String(http.StatusOK, "ok")
}, track)
```

## WebSocket

- https://echo.labstack.com/docs/cookbook/websocket
