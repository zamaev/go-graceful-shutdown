package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a channel to receive OS signals for termination (SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Main loop that runs until a termination signal is received
	for {
		select {
		case <-quit: // Exit on receiving a termination signal
			log.Println("Gracefully shutting down...")
			return
		default:
			job() // Run the job if no shutdown signal is received
		}
	}
}

// job simulates a long-running task
func job() {
	log.Println("Start job", time.Now())
	time.Sleep(2 * time.Second)
	log.Println("Done job")
}
