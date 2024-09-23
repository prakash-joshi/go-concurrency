package main

import (
	"fmt"
	"sync"
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
	// create a waitGroup for every one done eating
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is the map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosophers Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {

	defer wg.Done()

	// seat the philosophers at the table
	seated.Done()
	fmt.Printf("%s has been seated.\n", philosophers.name)
	seated.Wait()

	// philosopher eats
	for i := hunger; i >= 0; i-- {

		// get a lock on both forks
		forks[philosophers.leftFork].Lock()
		forks[philosophers.rightFork].Lock()
		fmt.Printf("\t%s has picked up left fork.\n", philosophers.name)
		fmt.Printf("\t%s has picked up right fork.\n", philosophers.name)

		fmt.Printf("\t%s has both the forks and is eating.\n", philosophers.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosophers.name)
		time.Sleep(thinkTime)

		forks[philosophers.leftFork].Unlock()
		forks[philosophers.rightFork].Unlock()

		fmt.Printf("\t%s has put down both the forks.\n", philosophers.name)
	}

	fmt.Println(philosophers.name, "has finished eating.")
	fmt.Println(philosophers.name, "has left the table.")

}
