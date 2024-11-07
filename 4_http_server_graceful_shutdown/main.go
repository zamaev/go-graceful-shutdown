package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a cancellable context that triggers on SIGINT or SIGTERM for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop() // Ensure resources are freed when the function exits

	// Configure the main HTTP router (mux)
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start job", time.Now())
		time.Sleep(5 * time.Second)    // Simulate a long-running task
		_, _ = w.Write([]byte("test")) // Send response to the client
		log.Println("Done job")
	})

	// Set up the HTTP
	server := &http.Server{
		Addr:    ":8333",
		Handler: mux,
	}

	// Run the server in a separate goroutine to avoid blocking
	go func() {
		log.Println("Server started on :8333")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server.ListenAndServe error: %v", err)
		}
	}()

	// Wait for termination signal to gracefully shutdown the server
	<-ctx.Done()
	log.Println("Shutdown signal received, shutting down server...")

	// Create a new context with timeout for server shutdown to ensure it completes gracefully
	// Note:
	// - Using the old context (`ctx`) would cause an immediate shutdown because it's already canceled by the signal.
	// - Using a context without a timeout could cause the shutdown to hang indefinitely if there are long-lived connections that don't close.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shut down the server, waiting for active connections to close
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped gracefully")
}
