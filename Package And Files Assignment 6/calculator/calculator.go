// 2. WAP in Go language using user defined package calculator that performs one calculator operation as per the0user's choice.



// Create the package: Create a folder named calculator and inside it create a file calculator.go with the following code:
// calculator/calculator.go

package calculator

import "errors"

func Add(a, b float64) float64 {
    return a + b
}

func Subtract(a, b float64) float64 {
    return a - b
}

func Multiply(a, b float64) float64 {
    return a * b
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero is not allowed")
    }
    return a / b, nil
}
