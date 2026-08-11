# Constants

This directory demonstrates how to declare and use constants in Go. Constants are immutable values that cannot be reassigned once defined.

## Key Concepts Covered

* **Constant Declaration**: Declared using the `const` keyword.
* **Immutability**: Unlike variables, constants cannot be modified after declaration (e.g., trying to reassign a `const` causes a compilation error).
* **Cannot Use Shorthand Syntax**: The shorthand syntax `:=` is only for variables, not constants.
* **Grouped Declaration**: Multiple constants can be grouped together in a parenthesized block for cleaner code.

## Code Example

```go
package constants

import "fmt"

func main() {
	// Grouped constant declaration
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run constant.go
```
