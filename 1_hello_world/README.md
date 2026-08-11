# Hello World

This directory contains a basic "Hello World" program in Go. It represents the starting point of learning the Go language, demonstrating how to write, compile, and run a simple program.

## Key Concepts Covered

* **Package Declaration**: Every Go file starts with a `package` declaration. The `package main` tells the Go compiler that this file should be compiled as an executable program rather than a shared library.
* **Import Statement**: The `import "fmt"` statement imports the standard format package, which contains functions for formatting text, including printing to the console.
* **Main Function**: The `func main()` is the entry point of the executable program.
* **Printing to Console**: `fmt.Println` is used to print messages to the standard output, followed by a newline.

## Code Example

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run main.go
```
