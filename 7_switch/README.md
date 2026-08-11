# Switch Statement

This directory demonstrates `switch` statements in Go, which provide a cleaner way to express multi-way conditionals.

## Key Concepts Covered

* **Basic Switch**: Compares a value against multiple cases. Unlike languages like C, Java, or JavaScript, Go has implicit `break` for cases; it does not fall through automatically.
* **Multiple Expressions in Case**: You can list multiple comma-separated values in a single `case` statement (e.g., `case time.Saturday, time.Sunday:`).
* **Type Switch**: A specialized switch statement that compares types rather than values. It uses `i.(type)` syntax inside an interface check to inspect the dynamic type of an interface variable.

## Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// Multiple-condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// Type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}
	whoAmI(55)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run swich.go
```
