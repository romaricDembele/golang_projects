package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	// 6. The philosopher is numbered using id
	id                              int
	left_chopstick, right_chopstick *Chopstick
}

func (philosopher *Philosopher) Eat(wg *sync.WaitGroup, start_channel chan bool, done_channel chan bool) {

	<-start_channel

	zero_or_one := rand.Intn(1)

	// 2. The philosopher eat only 3 times
	for i := 0; i < 3; i++ {

		// 3. The philosopher pick up the chopsticks in any order
		if zero_or_one == 0 {
			philosopher.left_chopstick.Lock()
			philosopher.right_chopstick.Lock()
		} else {
			philosopher.right_chopstick.Lock()
			philosopher.left_chopstick.Lock()
		}

		// 7. When the philosopher starts eating, it prints "starting to eat"
		fmt.Printf("Starting to eat %v\n", philosopher.id)

		// 8. When the philosopher finishes eating, it prints "finishing eating"
		fmt.Printf("Finishing eating %v\n", philosopher.id)

		philosopher.right_chopstick.Unlock()
		philosopher.left_chopstick.Unlock()
	}

	done_channel <- true

	wg.Done()
}

func host(wg *sync.WaitGroup, start_channel chan bool, done_channel chan bool) {
	start_channel <- true
	start_channel <- true

	for i := 0; i < 5; i++ {
		<-done_channel
		start_channel <- true
	}

	wg.Done()
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{id: i + 1, left_chopstick: chopsticks[i], right_chopstick: chopsticks[(i+1)%5]}
	}

	wg := sync.WaitGroup{}
	wg.Add(6)

	start_channel := make(chan bool, 2)
	done_channel := make(chan bool, 2)

	// 4. The host give permission to eat to the philosopher
	go host(&wg, start_channel, done_channel)

	// 1. 5 philosophers running on different goroutines
	for i := 0; i < 5; i++ {
		go philosophers[i].Eat(&wg, start_channel, done_channel)
	}

	wg.Wait()

}
