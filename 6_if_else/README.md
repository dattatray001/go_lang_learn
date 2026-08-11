# If/Else Conditionals

This directory covers conditional branch structures in Go using `if`, `else if`, and `else`.

## Key Concepts Covered

* **Basic If/Else**: Executing code blocks conditionally. Parentheses are not required around conditions, but curly braces `{}` are mandatory.
* **Logical Operators**: Combining multiple conditions using logical AND (`&&`) and logical OR (`||`).
* **Variable Declaration inside If statement**: You can declare and initialize local variables directly inside the `if` statement construct. These variables are only scoped to the `if` and `else` blocks (e.g., `if age := 20; age >= 18`).
* **No Ternary Operator**: Go does not have a ternary operator (like `cond ? expr1 : expr2`). You must use a standard `if/else` block.

## Code Example

```go
package ifelse

import "fmt"

func main() {
	age := 16

	// Basic if-else
	if age >= 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("person is not an adult")
	}

	// Else if chain
	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	// Initializing variable in condition
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command (ensure imports are correct and package is main if you execute):

```bash
go run ifesle.go
```
