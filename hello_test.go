package learn_go_with_tests

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got string, want string) {
		if got != want {
			t.Errorf("got '%q', but want '%q'", got, want)
		}
	}
	t.Run("test input string", func(t *testing.T) {
		got := hello("may", "")
		want := "Hello may"

		assertMessage(t, got, want)
	})

	t.Run("test empty input string", func(t *testing.T) {
		got := hello("", "")
		want := "Hello world"
		assertMessage(t, got, want)
	})

	t.Run("test specifying English", func(t *testing.T) {
		got := hello("amy", "English")
		want := "Hello amy"
		assertMessage(t, got, want)
	})

	t.Run("test specifying Spanish", func(t *testing.T) {
		got := hello("amy", "Spanish")
		want := "Hola amy"
		assertMessage(t, got, want)
	})

	t.Run("test specifying French", func(t *testing.T) {
		got := hello("amy", "French")
		want := "Bonjour amy"
		assertMessage(t, got, want)
	})
}
