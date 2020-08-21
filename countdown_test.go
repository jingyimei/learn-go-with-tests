package learn_go_with_tests

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountDownOperationSpy struct {
	Calls []string
}

func (c *CountDownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountDownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("test printing", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if spy.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spy.Calls)
		}

	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountDownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
