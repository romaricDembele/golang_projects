package main

import "fmt"

func main() {
	fmt.Println("Please enter a floating point number: ")
	var input float32
	fmt.Scan(&input)
	var trunc_input int = int(input)
	fmt.Printf("The truncated version of your floating point number is: %d", trunc_input)
}