package tools

import (
	"github.com/lthibault/jitterbug/v2"
	"time"
)

func NewTicker() (<-chan time.Time, func()) {
	t := jitterbug.New(
		time.Minute*5,
		&jitterbug.Norm{Stdev: time.Minute * 1},
	)

	c := make(chan time.Time) // We make a proxy channel, to have more control over what we do

	go func() {
		// Relay events
		for tick := range t.C {
			c <- tick
		}
	}()

	go func() {
		// Emit one event at start
		c <- time.Now()
	}()

	stopFunc := func() {
		t.Stop() // as many things, jitterbug is a little stupid, and stopping it would cause it to wait until the next tick to actually stop.
		// Good thing, golang (as of 1.19) just throws away all goroutines when the main thread exits, so we don't really have to worry about that
		close(c)
	}

	return c, stopFunc
}
