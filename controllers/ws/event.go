package ws

import (
	"context"
	"net/http"
	"project/ghost-catch-api/domain/values"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func Event(c echo.Context) error {
	t := c.Request().Header.Get("Sec-WebSocket-Protocol")
	if t == "" {
		c.Logger().Error("Not found custom header")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ans, err := strconv.Atoi(t)
	if err != nil {
		c.Logger().Error("Sec-WebSocket-Protocol is invalid", ans)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		ch := values.NewChoice()

		for {
			select {
			case <-ctx.Done():
				c.Logger().Info("game is timeout. close WebSocket connection")
				return
			default:
				// Read
				err = websocket.JSON.Receive(ws, ch)
				if err != nil {
					c.Logger().Error(err)
					if ch.Failed > 10 {
						c.Logger().Info("fail count is over. close WebSocket connection")
						return
					}
					ch.Failed += 1
				}

				if ch.Number == ans {
					c.Logger().Info("success")
					// Write
					err := websocket.JSON.Send(ws, values.NewResult(true))
					if err != nil {
						c.Logger().Error(err)
					}
				}
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
