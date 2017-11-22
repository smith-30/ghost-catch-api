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

	c.Logger().Info("answer is receive: ", ans)

	r := values.NewResult()

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
						err := websocket.JSON.Send(ws, r)
						if err != nil {
							c.Logger().Error(err)
						}
						return
					}
					ch.Failed += 1
				}

				if ch.Number == ans {
					// Todo 成功メッセージを送ったらクライアント側から切断されるか確かめる
					score := len(ch.Answers) + 1
					r.SetSuccess(score)

					c.Logger().Info("answer is valid.")
					// Write
					err := websocket.JSON.Send(ws, r)
					if err != nil {
						c.Logger().Error(err)
					}
				} else {
					ch.StockAnswer(ch.Number)
					r.Score = len(ch.Answers)
					err := websocket.JSON.Send(ws, r)
					if err != nil {
						c.Logger().Error(err)
					}
					c.Logger().Info("receive wrong answer: ", ch.Number)
				}
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
