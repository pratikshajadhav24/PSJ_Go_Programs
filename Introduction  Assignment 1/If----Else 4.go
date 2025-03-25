 //  Check if a Number is a Prime Number 

 package main

 import "fmt"
 
 func main() {
	 var num int
	 fmt.Print("Enter a number: ")
	 fmt.Scanln(&num)
 
	 if num < 2 {
		 fmt.Println("The number is not a prime number.")
	 } else {
		 isPrime := true
		 for i := 2; i*i <= num; i++ {
			 if num%i == 0 {
				 isPrime = false
				 break
			 }
		 }
		 if isPrime {
			 fmt.Println("The number is a prime number.")
		 } else {
			 fmt.Println("The number is not a prime number.")
		 }
	 }
 }
 