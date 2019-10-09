package controller

import (
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/K8s/handles"
)

func CustomerController(e *echo.Echo) {
	publicRoute := e.Group("/v1/test")
	publicRoute.GET("/", handles.TestHandles)
}