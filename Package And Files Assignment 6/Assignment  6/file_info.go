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
