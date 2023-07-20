package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	// Read the file name from the user.
	var input string
	fmt.Println("Please enter the file name you want to read (example : file.txt): ")
	fmt.Scan(&input)

	// Open the file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}

	// Read each line of the file, create a struct and add it to the slice of names.
	scanner := bufio.NewScanner(file)
	var names []name

	for scanner.Scan() {
		name_text := strings.Split(scanner.Text(), " ")
		name_map := name{fname: name_text[0], lname: name_text[1]}
		names = append(names, name_map)
	}
	if error := scanner.Err(); err != nil {
		fmt.Println(error)
	}

	// Iterate through the slice of structs and print the first ans last names found in each struct
	for _, person := range names {
		fmt.Println(person.fname, person.lname)
	}

	file.Close()
}
