# Go Graceful Shutdown Examples

This repository contains example code in Go for implementing graceful shutdowns in applications. Each example demonstrates different approaches for handling OS signals, managing contexts, and ensuring the orderly completion of long-running tasks.

## Repository Structure

- **[1_basic_signal_handling](./1_basic_signal_handling/main.go)**
  This example shows a basic approach to handling termination signals (SIGINT, SIGTERM) using a channel (`os.Signal`). The main loop runs a task (`job`) until a termination signal is received, at which point the application exits gracefully.

- **[2_signal_with_context_and_waitgroup](./2_signal_with_context_and_waitgroup/main.go)**
  This example adds a `context.WithCancel` and `sync.WaitGroup` for more controlled shutdowns. The context allows for safely canceling the execution of goroutines, while the `WaitGroup` ensures that all goroutines finish before the program exits.

- **[3_context_with_signal_notify](./3_context_with_signal_notify/main.go)**
  This example uses `signal.NotifyContext`, which automatically creates a context that responds to termination signals (SIGINT, SIGTERM). This is a more concise approach to graceful shutdown, as the context automatically releases resources when done and doesn't require an explicit `cancel` call.

- **[4_http_server_graceful_shutdown](./4_http_server_graceful_shutdown/main.go)**
  This example demonstrates graceful shutdown of an HTTP server. Upon receiving a termination signal, the server stops smoothly, completing active connections and freeing resources. A timeout is used for the shutdown context to prevent indefinite hanging in case of long-lived connections.

## Running Each Example

1. Run the Go program:
```
go run ./1_basic_signal_handling
go run ./2_signal_with_context_and_waitgroup
go run ./3_context_with_signal_notify
go run ./4_http_server_graceful_shutdown
```
2. To test graceful shutdown, send an interrupt signal (e.g., Ctrl+C in the terminal)
