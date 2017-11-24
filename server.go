package main

import (
	"project/ghost-catch-api/controllers/rest"
	"project/ghost-catch-api/controllers/ws"

	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// enable logger level info
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Test"},
	}))

	// Todo Auth jwt token when first access.

	/**

	  Static files

	*/
	e.Static("/", "static/html/top")
	e.Static("/img", "static/assets/img")
	e.Static("/games", "static/html/games")

	/**

	  API

	*/
	g := e.Group("/api/v1/games")
	g.GET("/ws/event", ws.Event)
	g.GET("/card", rest.Card)

	addr := os.Getenv("HTTP_PORT")
	if addr == "" {
		addr = "9090"
	}

	e.Logger.Fatal(e.Start(addr))
}
