package rest

import "testing"

func Test_genCardKey(t *testing.T) {
	key := genCardKey()
	t.Log(key)
}
