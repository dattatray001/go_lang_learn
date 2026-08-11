# Range (Iterating Data Structures)

This directory demonstrates the `range` keyword in Go, which provides a clean syntax to iterate over elements in various data structures (slices, arrays, maps, and strings).

## Key Concepts Covered

* **Slices and Arrays**: Iterating yields both the index and the element value (e.g., `for index, value := range slice`).
* **Maps**: Iterating yields key-value pairs (e.g., `for key, value := range m`). You can iterate over keys only by omitting the second variable (e.g., `for key := range m`).
* **Strings**: Iterating yields the starting byte index and the Unicode code point (rune) of each character. This is particularly useful for handling multi-byte characters.

## Code Example

```go
package main

import "fmt"

func main() {
	// Iterating over a slice
	nums := []int{6, 7, 8}
	for index, num := range nums {
		fmt.Println("Index:", index, "Value:", num)
	}

	// Iterating over a map
	m := map[string]string{"fname": "john", "lname": "doe"}
	for key, val := range m {
		fmt.Println(key, "->", val)
	}

	// Iterating over map keys only
	for key := range m {
		fmt.Println("Key:", key)
	}

	// Iterating over a string
	for index, c := range "golang" {
		fmt.Println("Byte Index:", index, "Character:", string(c))
	}
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run range.go
```
