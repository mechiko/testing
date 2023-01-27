package db

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServiceRoutes struct {
	// serviceUseCase entity.ServiceUseCase
}

// http://localhost:3600/v1/db/name
func NewServiceRoutes(handler *echo.Group) {
	r := &ServiceRoutes{
		// serviceUseCase: usecase.New(application.Get()),
	}

	h := handler.Group("/db")
	{
		h.GET("/test", r.Test)
	}
}

func (r *ServiceRoutes) Test(c echo.Context) error {

	c.String(http.StatusOK, "TESTING....")
	return nil
}
