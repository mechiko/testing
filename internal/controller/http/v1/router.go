// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"fmt"

	v1db "testing/internal/controller/http/v1/db"

	"github.com/labstack/echo/v4"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a Example service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// func NewRouter(handler *echo.Echo, uCase usecase.ExampleUseCase) {
// 	// Options
// 	// handler.Use(gin.Logger())
// 	// handler.Use(gin.Recovery())

// 	// K8s probe
// 	handler.GET("/healthz", func(c echo.Context) error {
// 		fmt.Println("GET(/healthz, func(c echo.Context)")
// 		return c.String(200, "ok")
// 	})

// 	// Routers
// 	h := handler.Group("/v1")
// 	{
// 		newExampleRoutes(h, uCase)
// 	}

// }

func NewRouterSimple(handler *echo.Echo) {
	// Options
	// handler.Use(gin.Logger())
	// handler.Use(gin.Recovery())

	// K8s probe
	handler.GET("/healthz", func(c echo.Context) error {
		fmt.Println("GET(/healthz, func(c echo.Context)")
		return c.String(200, "ok")
	})

	// Routers
	// http://localhost:3600/v1/db/name
	h := handler.Group("/v1")

	v1db.NewServiceRoutes(h)
}
