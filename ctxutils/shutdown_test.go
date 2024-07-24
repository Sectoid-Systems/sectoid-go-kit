package ctxutils

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"
)

// sendSignal is a helper function to send a signal after a delay.
func sendSignal(t *testing.T, sig os.Signal, delay time.Duration) {
	t.Helper()
	time.Sleep(delay)
	if err := syscall.Kill(syscall.Getpid(), sig.(syscall.Signal)); err != nil {
		t.Errorf("failed to send signal: %v", err)
	}
}

func TestWaitForShutdown_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	shutdownCalled := make(chan bool, 1)

	shutdown := func() error {
		shutdownCalled <- true
		return nil
	}

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	WaitForShutdown(ctx, shutdown, cancel)

	select {
	case <-shutdownCalled:
		// Expected behavior
	case <-time.After(2 * time.Second):
		t.Error("expected shutdown to be called on context cancel")
	}
}

func TestWaitForShutdown_SignalReceived(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	shutdownCalled := make(chan bool, 1)

	shutdown := func() error {
		shutdownCalled <- true
		return nil
	}

	go sendSignal(t, syscall.SIGTERM, 1*time.Second)

	WaitForShutdown(ctx, shutdown, cancel)

	select {
	case <-shutdownCalled:
		// Expected behavior
	case <-time.After(2 * time.Second):
		t.Error("expected shutdown to be called on signal received")
	}
}

func TestWaitForShutdown_ShutdownError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	expectedErr := errors.New("shutdown error")
	shutdownCalled := make(chan bool, 1)

	shutdown := func() error {
		shutdownCalled <- true
		return expectedErr
	}

	go sendSignal(t, syscall.SIGTERM, 1*time.Second)

	// Capture log output
	var logBuf strings.Builder
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	WaitForShutdown(ctx, shutdown, cancel)

	select {
	case <-shutdownCalled:
		if !strings.Contains(logBuf.String(), expectedErr.Error()) {
			t.Errorf("expected log message to contain %q, got %q", expectedErr.Error(), logBuf.String())
		}
	case <-time.After(2 * time.Second):
		t.Error("expected shutdown to be called on signal received")
	}
}
