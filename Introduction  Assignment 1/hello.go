package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() 

{
	// Create a new reader
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	fmt.Print("Enter your name: ")

	// Read user input
	name, _ := reader.ReadString('\n')

	// Print a greeting message
	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
}
