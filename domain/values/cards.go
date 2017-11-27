package values

import "strconv"

const (
	black = iota + 1
	blue
	green
	red
	yellow
)

var (
	Answers = []int{
		black, black, blue, blue, green,
		green, red, red, yellow, yellow,
	}
	Cards = map[int]*Card{}
)

func init() {
	makeCards()
}

func makeCards() {
	// Todo 画像パスで十分なはず.
	for i, ans := range Answers {
		i += 1
		name := "static/assets/img/cards/" + strconv.Itoa(i) + ".png"
		Cards[i] = NewCard(ans, name)
	}
}
