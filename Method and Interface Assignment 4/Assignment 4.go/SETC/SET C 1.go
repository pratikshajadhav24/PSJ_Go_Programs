// 1. Write a program in go language to create an interface and display it’s values with the help of type assertion.

package main

import "fmt"

func display(i interface{}) {
	val, ok := i.(string)
	if ok {
		fmt.Println("Value is:", val)
	} else {
		fmt.Println("Not a string")
	}
}

func main() {
	var data interface{} = "Hello, Interface!"
	display(data)

	data = 123
	display(data)
}
