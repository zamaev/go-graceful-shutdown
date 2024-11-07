package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Create a cancellable context to stop the server upon receiving SIGINT and SIGTERM signals
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop() // Ensure resources are freed when the context is no longer needed

	// WaitGroup to track active goroutines and wait for them to finish before shutting down
	wg := sync.WaitGroup{}
	wg.Add(1) // Increment the WaitGroup counter for the goroutine running the job

	// Launch the job in a separate goroutine to allow it to run concurrently
	go func() {
		defer wg.Done() // Mark this goroutine as done once it exits
		for {
			select {
			case <-ctx.Done(): // Listen for shutdown signal in the context
				return // Exit the goroutine when a shutdown signal is received
			default:
				job() // Execute the job if no shutdown signal is received
			}
		}
	}()

	// Wait for all goroutines tracked by the WaitGroup to complete
	wg.Wait()
	log.Println("All goroutines have completed, shutting down gracefully")
}

// job simulates a long-running task
func job() {
	log.Println("Start job", time.Now())
	time.Sleep(2 * time.Second)
	log.Println("Done job")
}
