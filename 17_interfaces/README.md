# Interfaces

This directory demonstrates **Interfaces** in Go. Interfaces are collections of method signatures that define the behavior of objects.

## Key Concepts Covered

* **Implicit Implementation**: In Go, interfaces are satisfied implicitly. A concrete type does not need to declare that it "implements" an interface using any keyword. It simply must implement all the methods specified by the interface.
* **Polymorphism and Mocking**: Interfaces allow writing modular code that adheres to the **Open-Closed Principle** (open for extension, closed for modification). You can substitute real implementations with mock implementations (e.g., a `fakepayment` for testing) without changing the consumer code.

## Code Example

```go
package main

import "fmt"

// Define the interface
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

// Struct that consumes the interface
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// Implementation 1: paypal
type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {}

// Implementation 2: fakepayment (for tests)
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway")
}

func (f fakepayment) refund(amount float32, account string) {}

func main() {
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run interface.go
```
