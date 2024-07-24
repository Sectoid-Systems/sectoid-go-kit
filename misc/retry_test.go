package misc

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestRetry tests the Retry function.
func TestRetry(t *testing.T) {
	tests := []struct {
		name      string
		timeouts  []time.Duration
		f         func(ctx context.Context) error
		expectErr bool
	}{
		{
			name:     "Success on first try",
			timeouts: []time.Duration{time.Second},
			f: func(ctx context.Context) error {
				return nil
			},
			expectErr: false,
		},
		{
			name:     "Success on second try",
			timeouts: []time.Duration{time.Second, time.Second},
			f: func() func(ctx context.Context) error {
				attempts := 0
				return func(ctx context.Context) error {
					attempts++
					if attempts == 1 {
						return errors.New("retry error")
					}
					return nil
				}
			}(),
			expectErr: false,
		},
		{
			name:     "All retries fail",
			timeouts: []time.Duration{time.Second, time.Second},
			f: func(ctx context.Context) error {
				return errors.New("retry error")
			},
			expectErr: true,
		},
		{
			name:     "Timeout before function completion",
			timeouts: []time.Duration{time.Millisecond},
			f: func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(2 * time.Millisecond):
					return nil
				}
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Retry(context.Background(), tt.f, tt.timeouts)
			if (err != nil) != tt.expectErr {
				t.Errorf("Retry() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// TestRetry_CancelContext tests the Retry function when the context is canceled.
func TestRetry_CancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Immediately cancel the context

	err := Retry(ctx, func(ctx context.Context) error {
		return errors.New("should not be called")
	}, []time.Duration{time.Second})

	if err == nil {
		t.Error("expected error due to canceled context, got nil")
	}
}
