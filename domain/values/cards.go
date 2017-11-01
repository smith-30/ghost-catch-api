package values

//import "strconv"

var (
	Answers = []int{2, 3, 4, 5}
	Cards   = map[int]*Card{}
)

func init() {
	makeCards()
}

func makeCards() {
	// Todo 画像パスで十分なはず.
	for i, ans := range Answers {
		//name := "static/assets/img/cards/card_" + strconv.Itoa(i) + ".jpeg"
		name := "static/assets/img/sample/sample.jpeg"
		Cards[i] = NewCard(ans, name)
	}
}
