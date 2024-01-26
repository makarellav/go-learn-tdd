package mocking

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if sleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", sleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		sleeperOperations := &SpyCountdownOperations{}

		Countdown(sleeperOperations, sleeperOperations)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(sleeperOperations.Calls, want) {
			t.Errorf("wanted calls %v, got %v", want, sleeperOperations.Calls)
		}
	})
}

func TestConfigurableSleeper_Sleep(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
