package main

import (
	"go-learn-tdd/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 3 * time.Second, SleepFunc: time.Sleep}

	mocking.Countdown(os.Stdout, sleeper)
}
