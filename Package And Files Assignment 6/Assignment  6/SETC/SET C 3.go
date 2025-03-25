//3. WAP in Go language to add or append content at the end of a text file.

//Create a file named append_file.go with the following code:

package main

import (
	"fmt"
	"os"
)

func main() {
	// File name to append data
	fileName := "example.txt"

	// Open the file in append mode, create if it doesn't exist
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Text to append
	textToAppend := "This is new content being appended.\n"

	// Write to the file
	_, err = file.WriteString(textToAppend)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content successfully appended to", fileName)
}
/*Step 2: Create a Sample File (Optional)
If example.txt does not exist, the program will create it automatically.
If you want to create it manually, run:

echo "Initial content in the file." > example.txt
Step 3: Run the Program
Run the Go program using:

go run append_file.go
Step 4: Verify the Content
Check the content of example.txt:

sh
Copy
Edit
cat example.txt
Expected Output in Terminal
css
Copy
Edit
Content successfully appended to example.txt*/