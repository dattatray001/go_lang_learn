# Channels

This directory covers **Channels** in Go. Channels are the pipes that connect concurrent goroutines, allowing them to send and receive values and synchronize execution.

## Key Concepts Covered

* **Unbuffered Channels**: Created using `make(chan Type)`. By default, sends and receives block until both the sender and receiver are ready. This acts as a synchronization mechanism.
* **Buffered Channels**: Created with a capacity (e.g., `make(chan Type, capacity)`). Sends only block when the buffer is full, and receives only block when the buffer is empty.
* **Channel Directions**: You can specify if a channel is send-only (`chan<- Type`) or receive-only (`<-chan Type`) in function parameters to provide compile-time type safety.
* **Closing Channels**: Senders can call `close(chan)` to signal that no more values will be sent. Receivers can test if a channel is closed using the multi-value receive expression (e.g., `val, ok := <-chan`).
* **Range over Channels**: You can iterate over values received from a channel using `for val := range chan`. The loop terminates automatically when the channel is closed.
* **Select Statement**: The `select` statement lets a goroutine wait on multiple channel operations. It blocks until one of its cases can run, then executes that case.

## Code Example

```go
package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	// Read until channel is closed
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 1. Basic Channels and Select
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { chan1 <- 10 }()
	go func() { chan2 <- "pong" }()

	for i := 0; i < 2; i++ {
		select {
		case val := <-chan1:
			fmt.Println("Received from chan1:", val)
		case val := <-chan2:
			fmt.Println("Received from chan2:", val)
		}
	}

	// 2. Buffered Channel and Range
	emailChan := make(chan string, 5)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 3; i++ {
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}
	close(emailChan) // Close channel to terminate range loop in goroutine

	<-done // Block until goroutine signals completion
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run channels.go
```
