# Maps

This directory demonstrates how to use **Maps** in Go. Maps are Go's built-in associative data type, representing hash tables, dictionaries, or key-value stores.

## Key Concepts Covered

* **Map Creation**: Created using `make(map[keyType]valueType)` or using a map literal (e.g., `m := map[string]int{"price": 40}`).
* **Adding and Modifying Elements**: Assign values using key brackets (e.g., `m["name"] = "golang"`).
* **Missing Keys and Zero Values**: Reading a key that doesn't exist in the map returns the zero-value of the value type (e.g., `0` for int, `""` for string).
* **Comma-Ok Idiom**: Used to distinguish between a key that doesn't exist vs. a key that is present with a value equal to the zero-value. It returns a secondary boolean (`value, exists := m[key]`).
* **Deleting Elements**: The built-in `delete(map, key)` function removes a key-value pair.
* **Clearing Maps**: The `clear(map)` built-in function (available in Go 1.21+) removes all key-value pairs from a map.
* **Equality**: Maps cannot be compared using `==`. Use `maps.Equal` from the standard `maps` library to compare two maps.

## Code Example

```go
package main

import (
	"fmt"
	"maps"
)

func main() {
	// Creating a map
	m := make(map[string]int)

	// Setting values
	m["age"] = 30
	m["price"] = 50

	// Getting values
	fmt.Println(m["age"])

	// Check if a key exists
	v, ok := m["phones"]
	if ok {
		fmt.Println("Key exists with value:", v)
	} else {
		fmt.Println("Key does not exist")
	}

	// Deleting a key
	delete(m, "price")

	// Comparing maps
	m1 := map[string]int{"price": 40}
	m2 := map[string]int{"price": 40}
	fmt.Println(maps.Equal(m1, m2)) // true
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run maps.go
```
