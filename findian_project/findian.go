package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Please enter some text: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.ToLower(input)

	i_index := strings.Index(input, "i")
	a_index := strings.Index(input, "a")
	n_index := strings.LastIndex(input, "n")

	if i_index == 0 && n_index == len(input)-1 && a_index != -1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}