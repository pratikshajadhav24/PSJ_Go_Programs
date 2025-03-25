// 2. WAP in Go language to print file information.

Create a file named file_info.go with the following code:

Code for file_info.go
go
Copy
Edit
package main

import (
	"fmt"
	"os"
)

func main() {
	// Specify the file name (change it to an existing file)
	fileName := "example.txt"

	// Get file information
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print file details
	fmt.Println("File Information:")
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
}
Step 2: Create a Sample File
Before running the program, create a sample file example.txt in the same directory.

sh
Copy
Edit
echo "Hello, Go!" > example.txt
Step 3: Run the Program
Run the Go program with:

sh
Copy
Edit
go run file_info.go