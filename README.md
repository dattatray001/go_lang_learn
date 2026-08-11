# Go Learning Repository

Welcome to the Go Learning Repository! This repository contains a curated set of Go examples, code structures, and explanations designed to help you learn the Go programming language step-by-step, starting from absolute basics up to concurrency and file operations.

## Repository Index

Each folder below contains a Go source file representing a topic and an explanatory `README.md` outlining the concepts, syntax, and instructions on how to run it.

| # | Topic | Description | Link |
|---|---|---|---|
| 1 | **Hello World** | Standard starting program showing package structure and standard output. | [Read README](./1_hello_world/README.md) |
| 2 | **Simple Values** | Basic built-in Go types: integer, string, boolean, float64, and arithmetic operations. | [Read README](./2_simple_values/README.md) |
| 3 | **Variables** | Explicit type declarations, type inference, shorthand syntax `:=`, and zero values. | [Read README](./3_variables/README.md) |
| 4 | **Constants** | Declaring immutable values and grouping constants in parenthesized blocks. | [Read README](./4_constants/README.md) |
| 5 | **For Loops** | Iterating using `for` (the only loop construct in Go), including infinite, conditional, classic, and range loops. | [Read README](./5_for/README.md) |
| 6 | **If/Else Conditionals** | Conditional branching, logical operators, inline variable initialization, and the absence of ternary operators. | [Read README](./6_if_else/README.md) |
| 7 | **Switch Statements** | Multi-way branching, matching multiple values in a single case, and type switches on interface values. | [Read README](./7_switch/README.md) |
| 8 | **Arrays** | Fixed-size sequences, length queries, zero initialization, and multi-dimensional matrices. | [Read README](./8_arrays/README.md) |
| 9 | **Slices** | Dynamic sequences in Go, length vs. capacity, pre-allocation with `make`, appending, copying, and slice bounds operators. | [Read README](./9_slices/README.md) |
| 10 | **Maps** | Hash tables / key-value dictionary structures, setting/getting, comma-ok idiom, and deletion/clearing. | [Read README](./10_maps/README.md) |
| 11 | **Range** | Iterating over slices, arrays, maps, and strings using the `range` keyword. | [Read README](./11_range/README.md) |
| 12 | **Functions** | Declaring parameters, multiple return values, ignoring returns with `_`, and functions as first-class citizens. | [Read README](./12_functions/README.md) |
| 13 | **Variadic Functions** | Functions taking any number of arguments, slice representation, and unpacking slices using `...`. | [Read README](./13_variadic_functions/README.md) |
| 14 | **Closures** | Implementing stateful, anonymous nested functions that reference outer scope variables. | [Read README](./14_closures/README.md) |
| 15 | **Pointers** | Passing arguments by reference instead of copy using pointer types (`*T`), address-of (`&`), and dereferencing (`*`). | [Read README](./15_pointers/README.md) |
| 16 | **Structs** | Grouping fields, composition/embedding nested structs, value vs. pointer receiver methods, and constructors. | [Read README](./16_structs/README.md) |
| 17 | **Interfaces** | Polymorphism in Go, implicit implementation, and applying the Open-Closed Principle for mocking and testing. | [Read README](./17_interfaces/README.md) |
| 18 | **Enums** | Creating enumerated types using custom types and grouped constants. | [Read README](./18_enums/README.md) |
| 19 | **Generics** | Writing reusable, type-parameterized functions and structs with constraints like `any` and `comparable`. | [Read README](./19_generics/README.md) |
| 20 | **Goroutines & WaitGroups** | Asynchronous task execution using the `go` keyword and managing synchronization with `sync.WaitGroup`. | [Read README](./20_goroutines/README.md) |
| 21 | **Channels** | Communicating across goroutines, unbuffered vs. buffered channels, channel directions, closing, and `select` blocks. | [Read README](./21_channels/README.md) |
| 22 | **Mutex** | Managing concurrent access to shared resources to prevent race conditions using `sync.Mutex`. | [Read README](./22_mutex/README.md) |
| 23 | **File Operations** | Reading, writing, streaming, extracting metadata, and deleting files and directories using standard library operations. | [Read README](./23_files/README.md) |
| 24 | **Packages & Modules** | Organizing codebase with Go packages, dependency management with `go.mod`/`go.sum`, and external library installation. | [Read README](./24_packages/README.md) |

## How to Run the Code

To run any of the code examples:
1. Ensure you have Go installed on your system. You can verify this by running:
   ```bash
   go version
   ```
2. Navigate to the desired folder in your terminal:
   ```bash
   cd <folder_name>
   ```
3. Run the Go source file:
   ```bash
   go run <filename>.go
   ```
