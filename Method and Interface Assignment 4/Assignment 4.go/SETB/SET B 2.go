// 2. Write a program in go language to demonstrate working type switch in interface.
package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type is int with value %d\n", v)
	case string:
		fmt.Printf("Type is string with value %s\n", v)
	case float64:
		fmt.Printf("Type is float64 with value %.2f\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType(42)
	checkType("Hello, Go!")
	checkType(3.14)
}
