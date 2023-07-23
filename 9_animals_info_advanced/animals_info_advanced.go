package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorBlue  = "\033[34m"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (cow Cow) Eat() {
	fmt.Println(string(colorGreen), cow.food, string(colorReset))
}

func (cow Cow) Move() {
	fmt.Println(string(colorGreen), cow.locomotion, string(colorReset))
}

func (cow Cow) Speak() {
	fmt.Println(string(colorGreen), cow.noise, string(colorReset))
}

func (bird Bird) Eat() {
	fmt.Println(string(colorGreen), bird.food, string(colorReset))
}

func (bird Bird) Move() {
	fmt.Println(string(colorGreen), bird.locomotion, string(colorReset))
}

func (bird Bird) Speak() {
	fmt.Println(string(colorGreen), bird.noise, string(colorReset))
}

func (snake Snake) Eat() {
	fmt.Println(string(colorGreen), snake.food, string(colorReset))
}

func (snake Snake) Move() {
	fmt.Println(string(colorGreen), snake.locomotion, string(colorReset))
}

func (snake Snake) Speak() {
	fmt.Println(string(colorGreen), snake.noise, string(colorReset))
}

func CheckAnimal(animal Animal, animal_name string) (Animal, bool) {
	var found bool
	var found_animal Animal

	switch s_animal := animal.(type) {
	case Cow:
		if s_animal.name == animal_name {
			found_animal = s_animal
			found = true
			break
		}
		found = false
	case Bird:
		if s_animal.name == animal_name {
			found_animal = s_animal
			found = true
			break
		}
		found = false
	case Snake:
		if s_animal.name == animal_name {
			found_animal = s_animal
			found = true
			break
		}
		found = false
	}
	return found_animal, found
}

func RetrieveAnimal(animals *map[string][]Animal, animal_name string) (Animal, bool) {
	var found_animal Animal
	var found bool

	for _, animal_type_animals := range *animals {
		for _, animal := range animal_type_animals {
			found_animal, found = CheckAnimal(animal, animal_name)
		}
	}
	return found_animal, found
}

func AddAnimalToDataBase(animals *map[string][]Animal, animal_name string, animal_type_name string) {
	switch animal_type_name {
	case "cow":
		cow := Cow{name: animal_name, food: "grass", locomotion: "walk", noise: "moo"}
		(*animals)["cows"] = append((*animals)["cows"], cow)
	case "bird":
		bird := Bird{name: animal_name, food: "worms", locomotion: "fly", noise: "peep"}
		(*animals)["birds"] = append((*animals)["birds"], bird)
	case "snake":
		snake := Snake{name: animal_name, food: "mice", locomotion: "slither", noise: "hsss"}
		(*animals)["snakes"] = append((*animals)["snakes"], snake)
	default:
		fmt.Println(string(colorRed), "The type of the animal is not correct.\nPlease choose between \"cow\", \"bird\", and \"snake\".", string(colorReset))
	}
}

func main() {
	// Print app usage instructions
	fmt.Println(string(colorBlue), "INSTRUCTIONS: \n The app allow you to make requests. You have two types of requests.", string(colorReset))
	fmt.Println(string(colorBlue), "The first type allow you to create a new animal. Request structure: newanimal [animal_name] [animal_type].\n You choose the animal name you want but you have only three choices for the animal type: \"cow\", \"bird\" or \"snake\".\n EXAMPLE: newanimal eagle bird", string(colorReset))
	fmt.Println(string(colorBlue), "The second request type allow you to get information of an animal. Request structure: query [animal_name] [action]. \n You choose an animal name that you previously created but you have three choices for the action: \"eat\", \"move\" or \"speak\".\n EXAMPLE: query eagle eat\n", string(colorReset))

	// Animals database
	animals := map[string][]Animal{}

	// Take user input and process the request
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " ")
		if inputs[0] == "newanimal" {
			// Create an animal and add it to the database.
			AddAnimalToDataBase(&animals, inputs[1], inputs[2])
			fmt.Println(string(colorGreen), "Created it!", string(colorReset))

		} else if inputs[0] == "query" {
			animal, status := RetrieveAnimal(&animals, inputs[1])
			if status == true {
				switch inputs[2] {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println(string(colorRed), "The chosen action is incorrect.\n Please choose between \"eat\", \"move\", and \"speak\".", string(colorReset))
				}
			} else {
				println(string(colorRed), "Animal not found.\n Please correct the animal name.", string(colorReset))
			}

		} else {
			fmt.Println(string(colorRed), "The command name is incorrect.\n Please choose between \"newanimal\" and \"query\".", string(colorReset))
		}
		fmt.Printf("> ")
	}

}
