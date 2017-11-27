package main

import (
	"project/ghost-catch-api/controllers/rest"
	"project/ghost-catch-api/controllers/ws"

	"golang.org/x/crypto/acme/autocert"

	"os"

	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	revision = "unknown"
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
	// prod かつ uaがpcの場合は404にする

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

	fmt.Println(`
|¯¯¯¯¯¯¯|° /¯¯,¯¯\  |¯¯¯|_|¯¯'|  /¯¯¯/\__) °|¯¯¯|_|¯¯¯| |\¯¯¯(\_/
|¯¯|__|¯¯|' |\____ /|' |\______/| |\     \/¯¯¯) |___|¯|___| \/
 ¯¯|__|¯¯   \|___ |/ ° \|_____|/  \|¯¯¯¯¯¯|  |___|¯|___| |¯¯¯¯¯¯¯|
                                              ¯¯¯¯¯¯'      ¯¯¯¯¯¯¯'`)
	fmt.Printf("\n\n revision at ")
	fmt.Printf("\x1b[32m%s\x1b[0m\n\n", revision)

	if env == "prod" {
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		defaultAddr := ":9090"
		e.Logger.Fatal(e.Start(defaultAddr))
	}
}
