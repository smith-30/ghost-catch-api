package values

type (
	Result struct {
		Result bool `json:"result"`
		Score  int  `json:"score"`
	}
)

func NewResult() *Result {
	return &Result{}
}

func (r *Result) SetSuccess(s int) {
	r.Score = s
	r.Result = true
}
