package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tranhuy-dev/K8s/controller"
)

func main() {
	fmt.Print("Hello world")
	initFrameword()
}

func initFrameword(){
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, //1KB
	}))
	controller.CustomerController(e)
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	fmt.Printf("Server listening at 3000")
	err := e.Start(":" + "3000")
	if err != nil {
		fmt.Println(err)
	}
}