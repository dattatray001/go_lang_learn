# Functions

This directory covers defining and using **Functions** in Go. Go functions are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.

## Key Concepts Covered

* **Basic Function Declaration**: Declaring parameters and return types. Parameters of the same type can share the type declaration (e.g., `func add(a, b int)`).
* **Multiple Return Values**: A function can return more than one value (e.g., `(string, string, bool)`). This is a common pattern in Go, especially for returning a result along with an error.
* **Blank Identifier**: Use `_` to ignore specific returned values when calling a function.
* **Anonymous Functions**: Functions defined inline without a name.
* **First-Class Functions**:
  * Passing functions as arguments to other functions.
  * Returning functions from other functions.

## Code Example

```go
package main

import "fmt"

// Share parameter types
func add(a, b int) int {
	return a + b
}

// Multiple return values
func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

// Returning a function
func processIt() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

func main() {
	// Call function with multiple returns
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	// Call basic function
	result := add(3, 5)
	fmt.Println("Result:", result)

	// Using functions as first-class objects
	fn := processIt()
	fmt.Println("Returned function output:", fn(6))
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run funcations.go
```
