package main

import (
	"fmt"
	"time"
)

// struct storing philosophers data
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of philosophers
var philosophers = []Philosopher{
	{name: "Aristotle", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Diogenes", leftFork: 1, rightFork: 2},
	{name: "Plato", leftFork: 2, rightFork: 3},
	{name: "Zeno", leftFork: 3, rightFork: 4},
}

// predefined variables
var (
	hunger    = 3               // how may times a person aet at once
	eatTime   = 1 * time.Second // time taken to eat once
	thinkTime = 3 * time.Second // time taken between eating to think
	sleepTime = 1 * time.Second // time taken to pause before eating again
)

func main() {
	// print a welcome message
	fmt.Println("Dining Philosophers Problem.")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")
}

func dine() {

}
