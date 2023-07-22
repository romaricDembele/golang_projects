package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(">")
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " ")
		switch {
		case inputs[0] == "cow" && inputs[1] == "eat":
			cow.Eat()
		case inputs[0] == "cow" && inputs[1] == "move":
			cow.Move()
		case inputs[0] == "cow" && inputs[1] == "speak":
			cow.Speak()
		case inputs[0] == "bird" && inputs[1] == "eat":
			bird.Eat()
		case inputs[0] == "bird" && inputs[1] == "move":
			bird.Move()
		case inputs[0] == "bird" && inputs[1] == "speak":
			bird.Speak()
		case inputs[0] == "snake" && inputs[1] == "eat":
			snake.Eat()
		case inputs[0] == "snake" && inputs[1] == "move":
			snake.Move()
		case inputs[0] == "snake" && inputs[1] == "speak":
			snake.Speak()
		default:
			fmt.Println("The entered choices are not correct. Valid example: cow eat.")
		}
		fmt.Println(">")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
