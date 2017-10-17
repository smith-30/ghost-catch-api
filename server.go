package main

import (
	"project/ghost-catch-api/controllers/ws/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/", hello)
	//e.POST("/", hello)

	g := e.Group("/games")
	g.GET("/ws/touch", handler.Touch)
	//g.POST("/play", hello)

	e.Logger.Fatal(e.Start(":1323"))
}
