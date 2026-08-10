package main

func main() {

	// ============================================================
	// 1. GET FILE INFORMATION
	// ============================================================

	// Open the file in read-only mode.
	// os.Open() returns:
	//   1. *os.File  -> file handle
	//   2. error     -> tells us whether the operation succeeded
	//
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// If the file doesn't exist or cannot be opened,
	// 	// err will contain the reason.
	// 	panic(err)
	// }
	//
	// // Always close a file after opening it.
	// // defer means f.Close() will execute when main() finishes.
	// defer f.Close()
	//
	// // Stat() gets metadata/information about the file.
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Name() returns the file name.
	// fmt.Println("file name:", fileInfo.Name())
	//
	// // IsDir() tells us whether this is a directory.
	// // false = file
	// // true  = directory
	// fmt.Println("file or folder:", fileInfo.IsDir())
	//
	// // Size() returns the file size in bytes.
	// fmt.Println("file size:", fileInfo.Size())
	//
	// // Mode() returns file permissions and file type information.
	// fmt.Println("file permission:", fileInfo.Mode())
	//
	// // ModTime() returns the last modified time.
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// ============================================================
	// 2. READ A FILE USING f.Read()
	// ============================================================

	// Open the file.
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Close the file automatically when main() finishes.
	// defer f.Close()
	//
	// // Create a byte slice with space for 12 bytes.
	// // The file content will be read into this buffer.
	// buf := make([]byte, 12)
	//
	// // Read() copies data from the file into buf.
	//
	// // d = number of bytes actually read
	// // err = error information
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Only process the bytes that were actually read.
	// for i := 0; i < d; i++ {
	// 	fmt.Println("data:", d, string(buf[i]))
	// }

	// ============================================================
	// 3. READ THE ENTIRE FILE USING os.ReadFile()
	// ============================================================

	// os.ReadFile() is the easiest way to read
	// the complete contents of a small file.
	//
	// It returns the file content as []byte.
	//
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Convert []byte to string before printing.
	// fmt.Println(string(data))

	// ============================================================
	// 4. READ DIRECTORY / FOLDER CONTENTS
	// ============================================================

	// Open the parent directory.
	//
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer dir.Close()
	//
	// // ReadDir(-1) reads all entries in the directory.
	// //
	// // Each entry is represented by os.DirEntry.
	// fileInfo, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Loop through every file/folder.
	// for _, fi := range fileInfo {
	//
	// 	// Name() gives the name of the entry.
	// 	// IsDir() tells us whether it is a directory.
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// ============================================================
	// 5. CREATE A FILE
	// ============================================================

	// os.Create() creates a new file.
	//
	// IMPORTANT:
	// If the file already exists, os.Create() truncates it
	// (removes its existing contents).
	//
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer f.Close()
	//
	// // WriteString() writes a string into the file.
	// f.WriteString("hi go")
	// f.WriteString("nice language")
	//
	// // We can also write []byte data.
	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	// ============================================================
	// 6. READ FROM ONE FILE AND WRITE TO ANOTHER
	//    USING STREAMING
	// ============================================================

	// Open the source file.
	//
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer sourceFile.Close()
	//
	// // Create the destination file.
	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer destFile.Close()
	//
	// // bufio.Reader provides convenient buffered reading.
	// reader := bufio.NewReader(sourceFile)
	//
	// // bufio.Writer provides buffered writing.
	// writer := bufio.NewWriter(destFile)
	//
	// // Keep reading until we reach the end of the file.
	// for {
	//
	// 	// Read one byte from the source file.
	// 	b, err := reader.ReadByte()
	//
	// 	if err != nil {
	//
	// 		// io.EOF means we reached the end of the file.
	// 		// EOF is not really an unexpected error here.
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	//
	// 		break
	// 	}
	//
	// 	// Write the byte to the destination file.
	// 	err = writer.WriteByte(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	//
	// // Buffered data may still be sitting in memory.
	// // Flush() makes sure it is actually written to the file.
	// writer.Flush()
	//
	// fmt.Println("written to new file successfully")

	// ============================================================
	// 7. DELETE A FILE
	// ============================================================

	// IMPORTANT:
	// We don't need to open a file before deleting it.
	// os.Remove() can directly delete the file.
	//
	// err := os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("file deleted successfully")
}
