package rest

import (
	"math"
	"math/rand"
	"net/http"
	"os"

	"time"

	"project/ghost-catch-api/domain/values"

	"github.com/labstack/echo"
)

func Card(c echo.Context) error {
	cardKey := genCardKey()
	c.Logger().Info("cardKey: ", cardKey)
	card, ok := values.Cards[cardKey]

	if !ok {
		c.Logger().Error("card not found. key: ", cardKey)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	file, err := os.Open(card.FileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer file.Close()

	fi, err := file.Stat() //FileInfo interface
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	data := make([]byte, fi.Size())
	file.Read(data)

	card.SetBase64Img(data)

	return c.JSON(http.StatusOK, card)
}

func genCardKey() int {
	rand.Seed(time.Now().UnixNano())
	cardKey := rand.Intn(len(values.Answers)) + 1 // not permit 0.
	time.Sleep(10 * time.Millisecond)
	_cardKey := rand.Intn(len(values.Answers))
	return int(math.Abs(float64(cardKey+_cardKey-25))) + 1
}
