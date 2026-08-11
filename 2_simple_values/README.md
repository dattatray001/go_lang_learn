# Simple Values

This directory demonstrates Go's built-in basic data types and simple values, including integers, floats, strings, booleans, and basic operations.

## Key Concepts Covered

* **Strings**: Sequences of characters, represented using double quotes (e.g., `"hello golang"`). Can be concatenated using the `+` operator.
* **Integers**: Whole numbers (e.g., `1`, `2`). Go supports basic arithmetic operations like addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
* **Floats**: Decimal numbers (e.g., `10.5`, `7.0`). Division between floats yields a float value (e.g., `7.0 / 3.0`).
* **Booleans**: Truth values, which can be `true` or `false`.

## Code Example

```go
package main

import "fmt"

func main() {
	// int
	fmt.Println(1 + 1)
	
	// string
	fmt.Println("hello golang")
	
	// bool
	fmt.Println(true)
	fmt.Println(false)
	
	// floats
	fmt.Println(10.5)
	fmt.Println(7.0 / 3.0)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run main.go
```
