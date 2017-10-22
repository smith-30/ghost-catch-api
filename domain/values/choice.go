package values

type (
	Choice struct {
		Number int `json:"number"` // json field must start uppercase.
		Failed int `json:"-"`
	}
)

func NewChoice() *Choice {
	return &Choice{}
}
