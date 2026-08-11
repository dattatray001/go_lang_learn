# Structs

This directory covers **Structs** in Go. Structs are typed collections of fields, useful for grouping data together to form records.

## Key Concepts Covered

* **Struct Definition**: Declaring fields with names and types (e.g., `type customer struct`).
* **Composition/Embedding**: Go supports struct composition instead of classical inheritance. You can embed one struct inside another (e.g., embedding `customer` in `order`), giving the outer struct direct access to the fields of the inner struct.
* **Initialization**: Structs can be initialized using field names, or sequentially (omitting field names). Unassigned fields receive their type's zero value.
* **Methods/Receiver Functions**: Functions attached to a specific struct type.
  * **Value Receivers** (`(o order)`): Operates on a copy of the struct. Cannot mutate its fields.
  * **Pointer Receivers** (`(o *order)`): Operates on the original struct. Can mutate its fields.
* **Constructor Functions**: Conventional helper functions to initialize and return a pointer to a struct (e.g., `func newOrder(...) *order`).
* **Anonymous Structs**: Structs declared without a defined type name, useful for one-off data shapes.

## Code Example

```go
package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  // Embedding customer struct
}

// Pointer receiver method (mutates the receiver)
func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	// Initializing a struct with embedded struct
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	// Modifying nested struct fields
	newOrder.customer.name = "robin"
	fmt.Println(newOrder)

	// Value changes using pointer receiver
	newOrder.changeStatus("confirmed")
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run struct.go
```
