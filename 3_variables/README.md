# Variables

This directory covers how variables are declared and initialized in Go. Go is statically typed, meaning variables must have a defined type at compile time. However, Go provides multiple ways to declare variables, including type inference and shorthand syntax.

## Key Concepts Covered

* **Explicit Type Declaration**: Declaring variables with the `var` keyword followed by the name, type, and optional value (e.g., `var name string = "golang"`).
* **Type Inference**: Declaring variables with `var` and letting the compiler infer the type based on the initialization value (e.g., `var name = "golang"`).
* **Shorthand Syntax**: Inside functions, the `:=` operator can be used to declare and initialize a variable without the `var` keyword or the type (e.g., `name := "golang"`).
* **Zero Value**: Declaring a variable without a value initializes it to its type's "zero value" (e.g., `""` for string, `0` for numbers, `false` for booleans).

## Code Example

```go
package variables

func main() {
	// Explicit declaration
	// var name string = "golang"

	// Type inference
	// var name = "golang"
	// var isAdult bool = true
	// var age int = 30

	// Shorthand syntax (only allowed inside functions)
	// name := "golang"

	// Declaration without immediate initialization (initializes to zero value)
	// var name string
	// name = "golang"

	// Float types
	// var price float32 = 50.5
	// var price = 50.5
	// price := 50.5
}
```

## How to Run

Since the package is declared as `variables` and the main body is currently commented out, this file serves as a reference. To run it, you would need to change the package to `package main`, uncomment the statements, import `"fmt"`, and use:

```bash
go run variable.go
```
