package values

type (
	Result struct {
		Value bool `json:"val"`
	}
)

func NewResult(val bool) *Result {
	return &Result{
		Value: val,
	}
}
