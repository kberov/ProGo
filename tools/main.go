// The main package
package main

import "fmt"

func main() {
	PrintHello()
	limit := 10
	for i := 0; i < limit; i++ {
		// i = i
		PrintHello()   //print again
		PrintNumber(i) //a number
	}
}

func PrintHello() {
	fmt.Println("Hello, Go")
}

// PrintNumber writes a number using the fmt.Println function
func PrintNumber(number int) {
	fmt.Println(number)
}
