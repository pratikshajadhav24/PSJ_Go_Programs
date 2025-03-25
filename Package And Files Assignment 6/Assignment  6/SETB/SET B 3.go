// 3. Write a function in Go language to find the square of a number and write a benchmark for it.

Create a file named square.go with the following code:

Code for square.go

package main

// Square function returns the square of a number.
func Square(n int) int {
    return n * n
}
Step 2: Create the Benchmark Test
Go provides the testing package to write benchmark tests. Create a file named square_test.go with the following code:

Code for square_test.go

package main

import (
    "testing"
)

// Benchmark for the Square function
func BenchmarkSquare(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Square(10) // Testing with a fixed number
    }
}
Step 3: Run the Benchmark Test
Open a terminal in the same directory.

Run the benchmark using:

go test -bench=.