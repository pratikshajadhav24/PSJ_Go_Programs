// 3. Write a program in go language to demonstrate working embedded interfaces.

package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type Device struct{}

func (Device) Print() {
	fmt.Println("Printing document...")
}

func (Device) Scan() {
	fmt.Println("Scanning document...")
}

func main() {
	var mfd MultiFunctionDevice = Device{}
	mfd.Print()
	mfd.Scan()
}
