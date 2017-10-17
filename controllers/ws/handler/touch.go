package handler

import (

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type Option struct {
	Id int `json:"id"`
}

type msg struct {
	Num int
}

func Touch(c echo.Context) error {
	u := websocket.Upgrader{}
	ws, err := u.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Read
		o := new(Option)
		err := ws.ReadJSON(o)
		if err != nil {
			c.Logger().Error(err)
		}

		if err = ws.WriteJSON(o); err != nil {
			c.Logger().Error(err)
		}
	}
}