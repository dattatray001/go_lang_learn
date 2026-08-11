# Arrays

This directory introduces arrays in Go. An array is a numbered sequence of elements of a specific, fixed length.

## Key Concepts Covered

* **Fixed Size**: Array size is part of its type (e.g., `[4]int` is different from `[5]int`). Arrays cannot grow or shrink.
* **Zero Initialization**: By default, elements in a new array are initialized to their zero values (`0` for integers, `""` for strings, `false` for booleans).
* **Length**: You can query the length of an array using the built-in `len()` function.
* **Single-line Declaration**: Initialize arrays inline with values using shorthand syntax (e.g., `nums := [3]int{1, 2, 3}`).
* **Multi-dimensional Arrays**: Support nested array structures, such as a 2D array representing a matrix (e.g., `[2][2]int`).
* **Characteristics**:
  * Predictable sizing.
  * Memory optimization (stored contiguously).
  * Constant time $O(1)$ access.

## Code Example

```go
package main

import "fmt"

func main() {
	// Zero-initialized array
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums) // [1 0 0 0]
	fmt.Println(len(nums)) // 4

	// Inline declaration
	vals := [3]int{1, 2, 3}
	fmt.Println(vals)

	// 2D Array
	matrix := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(matrix)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run arrays.go
```
