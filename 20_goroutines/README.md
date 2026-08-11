# Goroutines and WaitGroups

This directory covers **Goroutines** and **WaitGroups** in Go. Goroutines are lightweight threads of execution managed by the Go runtime, enabling concurrent execution of code.

## Key Concepts Covered

* **Goroutines**: Started by prefixing a function call with the `go` keyword (e.g., `go task(i)`). This runs the function asynchronously in the background.
* **Synchronization**: Main goroutine does not wait for background goroutines to finish by default. We use synchronization primitives like `sync.WaitGroup` to coordinate their completion.
* **sync.WaitGroup Methods**:
  * `Add(int)`: Increments the counter of active tasks to wait for.
  * `Done()`: Decrements the counter when a task completes. Often used with `defer` to ensure it runs upon function exit.
  * `Wait()`: Blocks the calling thread (usually the main goroutine) until the WaitGroup counter becomes zero.

## Code Example

```go
package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	// Signal WaitGroup that this task is complete upon exit
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	// Start 11 concurrent tasks
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	// Block until all 11 tasks call Done()
	wg.Wait()
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run goroutine.go
```
