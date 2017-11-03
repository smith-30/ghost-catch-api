package values

import "strconv"

var (
	Answers = []int{1, 2, 3, 4, 5}
	Cards   = map[int]*Card{}
)

func init() {
	makeCards()
}

func makeCards() {
	// Todo 画像パスで十分なはず.
	for i, ans := range Answers {
		i += 1
		//name := "static/assets/img/cards/card_" + strconv.Itoa(i) + ".jpeg"
		name := "static/assets/img/sample/card_" + strconv.Itoa(i) + ".jpg"
		//name := "static/assets/img/sample/sample_01.png"
		Cards[i] = NewCard(ans, name)
	}
}
