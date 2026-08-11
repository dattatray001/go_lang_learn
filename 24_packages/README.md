# Go Modules and Packages

This directory demonstrates how to organize your Go code into **Packages** and manage dependencies using **Go Modules**.

## Key Concepts Covered

1. **Go Module**: A collection of related packages versioned together. Initialized at the root of a project using `go mod init <module_path>` which generates a `go.mod` file.
2. **Go Packages**: Used to group related source files together. Every Go source file must declare a package (e.g. `package auth` or `package user`).
3. **Internal Imports**: Importing packages from the same module. The import path starts with the module path defined in `go.mod` followed by the subfolder path (e.g. `import "github.com/dattatary001/myApp/auth"`).
4. **Visibility Rules (Exported vs Unexported)**:
   * Identifiers (functions, structs, variables, fields) that start with a **Capital letter** are exported (public) and visible outside their package (e.g., `GetSession()`, `User.Email`).
   * Identifiers starting with a **lowercase letter** are unexported (private) and only accessible within the package they are defined (e.g., `extractSession()`).
5. **External Dependencies**: Installing third-party packages using `go get <package_url>` (e.g., `go get github.com/fatih/color`). This adds dependencies to the `go.mod` file and records checksums in `go.sum` for build verification.
6. **Dependency Cleanup**: Running `go mod tidy` synchronizes the `go.mod` and `go.sum` files to match the imports in the source code, downloading missing modules and removing unused ones.

## Project Structure

```text
24_packages/
├── go.mod                     # Module definition and dependencies
├── go.sum                     # Cryptographic checksums of dependencies
├── main.go                    # Main application entry point
├── auth/                      # Custom package 'auth'
│   ├── credentials.go
│   └── session.go
└── user/                      # Custom package 'user'
    └── user.go
```

## How to Run

Navigate to this directory and compile/run using:

```bash
go run main.go
```
