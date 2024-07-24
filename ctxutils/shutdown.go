package ctxutils

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ShutdownFunc defines the function signature for shutdown procedures.
type ShutdownFunc func() error

// WaitForShutdown waits for an interrupt or context cancellation signal to perform a shutdown.
func WaitForShutdown(ctx context.Context, shutdown ShutdownFunc, cancel context.CancelFunc) {
	defer cancel()

	// Create a channel to listen for interrupt signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigChan)
	defer close(sigChan)

	select {
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down...")
	case sig := <-sigChan:
		log.Printf("Received signal: %s, shutting down...", sig)
	}

	// Execute the shutdown function.
	if err := shutdown(); err != nil {
		log.Printf("Shutdown error: %s", err)
	}

	log.Println("Bye!")
}
