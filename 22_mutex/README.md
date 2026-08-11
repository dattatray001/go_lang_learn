# Mutex (Mutual Exclusion)

This directory covers **Mutexes** (`sync.Mutex`) in Go. In concurrent programming, a mutex is used to protect shared data from being accessed/modified by multiple goroutines simultaneously, preventing race conditions.

## Key Concepts Covered

* **Race Condition**: Occurs when multiple goroutines try to access and modify the same memory location concurrently.
* **Mutual Exclusion**: We can use a `sync.Mutex` to lock code blocks (critical sections) so that only one goroutine can execute them at a time.
* **Lock and Unlock**:
  * `mu.Lock()`: Acquires the lock. If another goroutine holds the lock, it blocks until the lock is released.
  * `mu.Unlock()`: Releases the lock. Usually paired with `defer` to guarantee unlocking even if the function exits early or panics.

## Code Example

```go
package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // Mutex guarding the views count
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Always unlock
		wg.Done()
	}()

	p.mu.Lock() // Lock access before modifying shared resource
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	// Start 100 goroutines to increment views concurrently
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println("Total Views (should be exactly 100):", myPost.views)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run mutex.go
```
