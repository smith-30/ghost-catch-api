package values

type (
	Choice struct {
		Number  int   `json:"number"` // json field must start uppercase.
		Failed  int   `json:"-"`
		Answers []int `json:"-"`
	}
)

func NewChoice() *Choice {
	return &Choice{}
}

func (c *Choice) StockAnswer(answer int) {
	var exists bool

	for _, a := range c.Answers {
		if a == answer {
			exists = true
		}
	}

	if !exists {
		c.Answers = append(c.Answers, answer)
	}
}
