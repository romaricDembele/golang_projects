package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Println("Enter a name: ")
	fmt.Scan(&name)
	fmt.Println("Enter an address: ")
	fmt.Scan(&address)

	person := map[string]string{
		"name":    name,
		"address": address}

	person_json, err := json.Marshal(person)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The JSON object representing the entered data is: %s", person_json)
}
