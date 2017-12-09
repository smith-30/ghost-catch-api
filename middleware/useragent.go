package middleware

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/mssola/user_agent"
)

func UserAgent() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ua := user_agent.New(c.Request().UserAgent())
			c.Logger().Info(c.Request().UserAgent())
			c.Logger().Info(ua.Mobile())
			if ua.Mobile() {
				return nil
			}
			return errors.New("not permitted pc. only mobile is accessible.")
		}
	}
}
