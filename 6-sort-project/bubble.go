package main

import (
	"fmt"
)

func main() {
	// Take user input
	fmt.Println("Please enter in a sequence of up to 10 integers (example: 643 22 11 999 4 8 3 9 55 6): ")
	data := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		var input int
		fmt.Scan(&input)
		data = append(data, input)
	}

	// Sort the user input and print it to the console
	BubbleSort(&data)
	fmt.Printf("The sorted slice is: %v", data)
}

func Swap(numbers []int, i int) {
	temp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = temp
}

func BubbleSort(numbers *[]int) {
	for i := 0; i < len(*numbers)-1; i++ {
		swapped := false
		for j := 0; j < len(*numbers)-1; j++ {
			if (*numbers)[j+1] < (*numbers)[j] {
				Swap(*numbers, j)
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}
