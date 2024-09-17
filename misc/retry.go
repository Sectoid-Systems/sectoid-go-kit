package misc

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type RetryPrinter func(v ...any)
type RetryLogger func(format string, v ...any)

// Retry runs the provided function len(timeouts) times, passing a context with each timeout as a parameter.
// Waits at least until the timeout before moving on to the next attempt. Returns result of last call.
func Retry(ctx context.Context, f func(ctx context.Context) error, timeouts []time.Duration) error {
	return RetryWithNoise[*any](ctx, f, nil, timeouts)
}

func RetryWithNoise[T RetryPrinter | RetryLogger | Nilable](ctx context.Context, f func(ctx context.Context) error, out T, timeouts []time.Duration) error {
	for i, timeout := range timeouts {
		start := time.Now()
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)

		err := f(ctxWithTimeout)
		cancel() // ensure the cancel is called at the end of each iteration

		if err == nil {
			return nil // successful attempt
		}

		if IsNotNil(out) {
			switch f := any(out).(type) {
			case RetryPrinter:
				f(fmt.Sprintf("attempt %d failed: %v", i+1, err))
			case RetryLogger:
				f("attempt %d failed: %v", i+1, err)
			}
		}

		if i == len(timeouts)-1 {
			return err // last attempt
		}

		time.Sleep(time.Until(start.Add(timeout)))
	}

	return errors.New("unexpected: retry completed without returning") // should never happen
}
