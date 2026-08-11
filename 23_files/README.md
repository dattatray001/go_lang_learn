# File Operations

This directory covers various file-system operations in Go. It shows how to read, write, stream, retrieve metadata, and delete files and directories using the built-in standard library packages (`os`, `bufio`, `io`).

## Key Concepts Covered

1. **Getting File Metadata**: Use `os.Open()` and `f.Stat()` to query properties like filename, permissions, size, modification time, and check if the path is a directory.
2. **Reading Files**:
   * **Buffered Chunk Reading**: Reading chunk by chunk into a pre-allocated byte slice buffer using `f.Read()`.
   * **Full Reading**: Using `os.ReadFile()` for a convenient one-step read of small files.
3. **Reading Directories**: Using `os.Open` on a directory path and running `dir.ReadDir(-1)` to list all file entries.
4. **Creating and Writing Files**: Using `os.Create()` (which truncates if the file already exists), followed by `f.WriteString()` or `f.Write()`.
5. **Streaming Files**: Copying content byte-by-byte or block-by-block using buffered readers and writers (`bufio.NewReader`, `bufio.NewWriter`). Handle termination when the error message is `"EOF"` (End of File). Call `writer.Flush()` to empty buffered memory to disk.
6. **Deleting Files**: Using `os.Remove(filename)` directly.

## Code Example

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Create and Write to a file
	f, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	f.WriteString("Hello, Go Files!")
	f.Close()

	// 2. Read the entire file
	content, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File content:", string(content))

	// 3. File metadata
	fInfo, err := os.Stat("example.txt")
	if err == nil {
		fmt.Println("Name:", fInfo.Name())
		fmt.Println("Size (bytes):", fInfo.Size())
	}

	// 4. Delete file
	os.Remove("example.txt")
}
```

## How to Run

To run the program, navigate to this directory and use the `go run` command:

```bash
go run file.go
```
