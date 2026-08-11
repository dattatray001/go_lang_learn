# Enums (Enumerated Types)

This directory demonstrates how to implement **Enums** in Go. Since Go does not have a native `enum` keyword, enums are represented using custom types and grouped constants.

## Key Concepts Covered

* **Custom Type Alias**: Defining a new type based on a primitive type like `string` or `int` (e.g., `type OrderStatus string`).
* **Grouped Constant Constants**: Declaring a set of predefined values of the custom type using `const` syntax. Subsequent values can omit the type declaration as they inherit it.
* **Type Safety**: Using custom types instead of raw strings or integers ensures that functions only accept values from the defined enumeration.

## Code Example

```go
package main

import "fmt"

// Define a custom type OrderStatus
type OrderStatus string

// Define constants of the OrderStatus type
const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	// Prepared is a valid OrderStatus
	changeOrderStatus(Prepared)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run enum.go
```
