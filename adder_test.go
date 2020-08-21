package learn_go_with_tests

import "testing"

func TestAdd(t *testing.T) {
	i := 1
	j := 2

	sum := add(i, j)
	if sum != 3 {
		t.Errorf("wrong! expected %d, get %d", 3, sum)
	}
}
