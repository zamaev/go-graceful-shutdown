package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Channel to receive OS shutdown signals (SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Create a cancellable context to handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the context is canceled on function exit

	// WaitGroup to track active goroutines and ensure they complete on shutdown
	var wg sync.WaitGroup

	// Start the job function in a separate goroutine
	wg.Add(1) // Increment the counter before starting the goroutine
	go func() {
		defer wg.Done() // Ensure the WaitGroup counter is decremented when goroutine exits

		for {
			select {
			case <-ctx.Done(): // If context is canceled, exit the goroutine
				return
			default:
				job() // Run the job if no shutdown signal is received
			}
		}
	}()

	// Block until a shutdown signal is received
	<-quit
	log.Println("Shutdown signal received. Gracefully shutting down...")

	// Cancel the context to signal all goroutines to stop
	cancel()

	// Wait for all goroutines to complete before exiting
	wg.Wait()
	log.Println("All goroutines completed. Exiting program.")
}

// job simulates a long-running task
func job() {
	log.Println("Start job", time.Now())
	time.Sleep(2 * time.Second) // Simulate work by sleeping for 2 seconds
	log.Println("Done job")
}
