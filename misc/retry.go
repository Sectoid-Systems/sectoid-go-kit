package misc

import (
	"context"
	"errors"
	"time"
)

// Retry runs the provided function len(timeouts) times, passing a context with each timeout as a parameter.
// Waits at least until the timeout before moving on to the next attempt. Returns result of last call.
func Retry(ctx context.Context, f func(ctx context.Context) error, timeouts []time.Duration) error {
	for i, timeout := range timeouts {
		start := time.Now()
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)

		err := f(ctxWithTimeout)
		cancel() // ensure the cancel is called at the end of each iteration

		if err == nil {
			return nil // successful attempt
		}

		if i == len(timeouts)-1 {
			return err // last attempt
		}

		time.Sleep(time.Until(start.Add(timeout)))
	}

	return errors.New("unexpected: retry completed without returning") // should never happen
}
