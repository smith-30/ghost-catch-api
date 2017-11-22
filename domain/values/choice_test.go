package values_test

import (
	"project/ghost-catch-api/domain/values"
	"testing"
)

func TestChoice_StockAnswer(t *testing.T) {
	ch := values.NewChoice()

	ch.StockAnswer(1)
	ch.StockAnswer(1)
	ch.StockAnswer(1)
	t.Logf("%v\n", ch.Answers)
}
