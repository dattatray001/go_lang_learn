# Closures

This directory covers **Closures** in Go. Go supports anonymous functions, which can form closures. A closure is a function value that references variables from outside its body.

## Key Concepts Covered

* **State Preservation**: The closure function preserves and can modify the state of variables in its surrounding scope even after the outer function has finished executing.
* **Isolation**: Each time the outer function is called, it creates a new, isolated set of state variables for the returned closure.

## Code Example

```go
package main

import "fmt"

// counter returns a closure that encloses and increments 'count'
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	// First closure instance
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	// Second closure instance (gets its own independent count)
	newIncrement := counter()
	fmt.Println(newIncrement()) // 1
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run closures.go
```
