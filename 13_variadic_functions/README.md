# Variadic Functions

This directory covers **Variadic Functions** in Go. A variadic function is a function that can accept a variable number of arguments of a specific type.

## Key Concepts Covered

* **Variadic Parameter**: Declared using the ellipsis `...` before the type (e.g., `nums ...int`).
* **Slice Representation**: Inside the function, the variadic parameter is treated as a slice of the specified type (e.g., `nums` behaves as `[]int`).
* **Slice Unpacking**: If you have a slice of elements and want to pass it as individual arguments to a variadic function, suffix the slice variable name with `...` (e.g., `sum(nums...)`).

## Code Example

```go
package main

import "fmt"

// Variadic function that sums any number of integers
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Call variadic function with individual arguments
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4))

	// Call variadic function by unpacking a slice
	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run vardic.go
```
