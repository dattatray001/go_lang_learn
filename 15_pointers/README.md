# Pointers

This directory demonstrates **Pointers** in Go. Pointers allow you to pass references to values and records within your program rather than copying the values.

## Key Concepts Covered

* **Pass by Value (Default)**: Go is a pass-by-value language. When you pass a variable to a function, it copies the value. Any modification inside the function does not affect the original variable.
* **Pass by Reference**: Using pointers, you can pass a reference (memory address) to a variable. Modifying the dereferenced pointer affects the original variable.
* **Pointer Syntax**:
  * `*T`: The type of a pointer pointing to a value of type `T`.
  * `&val`: The address-of operator, which returns the memory address of a variable `val`.
  * `*ptr`: The dereference operator, which accesses or modifies the value at the memory address stored in the pointer `ptr`.

## Code Example

```go
package main

import "fmt"

// changeNum accepts an integer pointer (by reference)
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum:", *num)
}

func main() {
	num := 1

	// Pass the address of 'num'
	changeNum(&num)

	// 'num' has been modified in the main function as well
	fmt.Println("After changeNum in main:", num) // 5
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run pointer.go
```
