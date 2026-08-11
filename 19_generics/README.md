# Generics

This directory demonstrates **Generics** in Go. Introduced in Go 1.18, generics allow you to write functions and structs that work with any of a set of types, parameterized by a type list.

## Key Concepts Covered

* **Type Parameters**: Declared in square brackets `[...]` after the function or struct name.
* **Constraints**: 
  * Built-in interfaces that restrict which types can be used.
  * Common constraints include `any` (allowing any type) and `comparable` (restricting to types that support the `==` and `!=` operators).
* **Generic Functions**: Functions that accept slices or values of a generic type parameter (e.g., `func printSlice[T comparable](items []T)`).
* **Generic Structs**: Struct definitions that hold generic fields (e.g., `type stack[T any] struct { elements []T }`).

## Code Example

```go
package main

import "fmt"

// Generic function with two type parameters
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// Generic struct representing a stack (LIFO)
type stack[T any] struct {
	elements []T
}

func main() {
	// Using the generic struct
	myStack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println(myStack)

	// Using the generic function
	vals := []bool{true, false, true}
	printSlice(vals, "john")
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run generices.go
```
