package rest

import (
	"math/rand"
	"net/http"
	"os"

	"time"

	"project/ghost-catch-api/domain/values"

	"github.com/labstack/echo"
)

func Card(c echo.Context) error {

	rand.Seed(time.Now().UnixNano())
	cardKey := rand.Intn(len(values.Answers))
	card, ok := values.Cards[cardKey]

	if !ok {
		c.Logger().Error("card not found.")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	file, _ := os.Open(card.FileName)
	defer file.Close()

	fi, err := file.Stat() //FileInfo interface
	if err != nil {
		c.Logger().Error("%v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	data := make([]byte, fi.Size())
	file.Read(data)

	card.SetBase64Img(data)

	return c.JSON(http.StatusOK, card)
}
