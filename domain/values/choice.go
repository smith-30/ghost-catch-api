package values

type (
	Choice struct {
		Number int `json:"number"` // json field must start uppercase.
		Score  int `json:"-"`      // total score that the user operates.
		Failed int `json:"-"`
	}
)

func NewChoice() *Choice {
	return &Choice{}
}
