package main

import (
	"project/ghost-catch-api/controllers/rest"
	"project/ghost-catch-api/controllers/ws"

	"golang.org/x/crypto/acme/autocert"

	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	env := os.Getenv("ENV")
	if env == "prod" {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(os.Getenv("HOST"))
	}
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

	addr := os.Getenv("BIND_HTTP")
	if addr == "" {
		addr = ":9090"
	}

	if env == "prod" {
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.Start(addr))
	}
}
