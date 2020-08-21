package learn_go_with_tests

import "testing"

func TestRepeat(t *testing.T) {
	got := repeat("s")
	if got != "sssss" {
		t.Errorf("wrong! get %q but want %q", got, "sssss")
	}
}