package main

import (
	"fmt"
)

func GenDisplaceFn(a float64, v_o float64, s_o float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v_o * t) + s_o
	}
}

func main() {
	fmt.Println("Please enter 3 float numbers representing the acceleration, the initial velocity and the initial displacement (example: 3.4 5 77.2):")
	inputs := make([]float64, 0, 3)
	for i := 0; i < 3; i++ {
		var input float64
		fmt.Scan(&input)
		inputs = append(inputs, input)
	}

	fn := GenDisplaceFn(inputs[0], inputs[1], inputs[2])

	fmt.Println("Please enter a float number representing the time:")
	second_input := 0
	fmt.Scan(&second_input)

	fmt.Printf("The corresponding displacement value is: %v", fn(float64(second_input)))

}
