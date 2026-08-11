# For Loops

This directory covers loop structures in Go. In Go, `for` is the **only** looping construct. There are no `while` or `do-while` loops, but `for` can be configured to mimic them.

## Key Concepts Covered

* **Single Condition Loop (While-like)**: A loop with a single condition that executes as long as the condition is true (e.g., `for i <= 3`).
* **Infinite Loop**: A loop without any condition (e.g., `for {}`). It runs indefinitely unless stopped by a `break` statement or return.
* **Classic For Loop**: A loop containing initialization, condition, and post-iteration statements (e.g., `for i := 0; i <= 3; i++`).
* **Loop Control Statements**:
  * `break`: Exits the loop immediately.
  * `continue`: Skips the rest of the current iteration and starts the next one.
* **Range Loop over Integer**: Starting in Go 1.22, you can iterate over a range of integers directly (e.g., `for i := range 11`).

## Code Example

```go
package main

import "fmt"

func main() {
	// while loop style
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// Go 1.22 integer range loop
	for i := range 11 {
		fmt.Println(i)
	}
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run for.go
```
