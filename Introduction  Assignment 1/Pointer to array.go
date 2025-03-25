package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	var p *int = &arr[0]
	fmt.Println("First element:", *p)
	p = &arr[1]
	fmt.Println("Second element:", *p)
}
