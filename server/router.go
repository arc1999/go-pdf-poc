package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-pdf-poc/controller"
)

// NewRouter --
func NewRouter() (e *echo.Echo) {
	e = echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization", "auth_token"},
	}))
	ctrl := new(controller.PrescriptionController)

	group := e.Group("/api/v1/dt/prescription")
	{
		group.POST("/create", ctrl.Create())
		group.PUT("/update", ctrl.Update())
		group.GET("/draft/:patientId", ctrl.GetByPatientID())
	}
	return e
}
