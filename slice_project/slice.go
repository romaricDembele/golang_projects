package main

import (
	"fmt"
	"sort"
)

func main(){
	numbers_slice := make([]int, 0, 3)

	for {
		// Taking input from the user.
		fmt.Println("Please, enter an integer:")
		var input int
		_ , err := fmt.Scan(&input)

		// The program quit the existing loop when anything instead of an integer is entered (including the character 'X').
		if err != nil {
			break
		}

		// Append the user input to the slice; sort and print the slice.
		numbers_slice = append(numbers_slice, input)
		sort.Ints(numbers_slice)
		fmt.Println("The sorted slice is:")
		fmt.Println(numbers_slice)
	}
}