# Slices

This directory covers **Slices** in Go. Slices are wrappers around arrays that provide a dynamic, flexible interface to sequences of data. Unlike arrays, slices can grow and shrink dynamically, and are the most commonly used sequence structure in Go.

## Key Concepts Covered

* **Uninitialized Slices**: An uninitialized slice has the value `nil` and a length of 0.
* **Length vs. Capacity**:
  * **Length** (`len`): The number of elements currently in the slice.
  * **Capacity** (`cap`): The maximum number of elements the slice can hold before reallocating memory.
* **Slice Creation**:
  * Using `make` to pre-allocate size and capacity (e.g., `make([]int, length, capacity)`).
  * Direct literals (e.g., `nums := []int{}`).
* **Appending Elements**: The built-in `append` function adds elements to the end of a slice. If capacity is exceeded, Go allocates a new larger backing array.
* **Copying Slices**: The built-in `copy` function copies elements from a source slice to a destination slice. The destination slice must be pre-allocated with the desired size.
* **Slice Operator**: Extract subsets of a slice using `[start:end]` (non-inclusive index of the end element).
* **Equality**: Slices cannot be compared directly using `==`. Use `slices.Equal` from the standard `slices` package (available in Go 1.21+).

## Code Example

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Creating slices
	var nums []int // nil slice
	fmt.Println(nums == nil) // true

	// Pre-allocating capacity
	nums2 := make([]int, 0, 5)
	fmt.Println(cap(nums2)) // 5

	// Appending
	nums2 = append(nums2, 1, 2)
	fmt.Println(nums2) // [1 2]

	// Copying
	nums3 := make([]int, len(nums2))
	copy(nums3, nums2)

	// Slice operator
	fullList := []int{1, 2, 3, 4, 5}
	fmt.Println(fullList[1:4]) // [2 3 4]

	// Slice comparisons
	fmt.Println(slices.Equal(nums2, nums3)) // true
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run slices.go
```
