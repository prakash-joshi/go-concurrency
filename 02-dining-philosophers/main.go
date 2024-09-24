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

// dine simulates the dining philosophers problem using goroutines and wait groups.
// It initializes a wait group to track when all philosophers have finished eating and another for seating.
// A map of forks (mutexes) is created to represent shared resources.
// For each philosopher, a goroutine is started to simulate the dining process with a call to `diningProblem`.
// The function waits for all philosophers to complete their dining before returning.
func dine() {
	// create a waitGroup for every one done eating
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// create a waitGroup for every one done seating
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

	// wait for the philosophers to finish dining
	wg.Wait()
}

// diningProblem simulates the actions of a single philosopher in the dining philosophers problem.
// Each philosopher will attempt to pick up the two forks they need to eat, one on their left and one on their right.
// The function ensures that mutual exclusion is maintained to avoid deadlock or starvation.
// Parameters:
// - philosophers: A reference to a Philosopher, representing the individual philosopher participating in the simulation.
// - wg: A pointer to sync.WaitGroup, used to track when all philosophers have finished their actions.
// - forks: A map where each fork is represented by a sync.Mutex, ensuring only one philosopher can hold a fork at any time.
// - seated: A pointer to sync.WaitGroup, used to synchronize when all philosophers are seated before they start dining.
func diningProblem(philosophers Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {

	defer wg.Done()

	// seat the philosophers at the table
	seated.Done()
	fmt.Printf("%s has been seated.\n", philosophers.name)
	seated.Wait()

	// philosopher eats
	for i := hunger; i >= 0; i-- {

		// get a lock on both forks

		// edge case for race condition
		// logical race condition here may cause a deadlock situation
		// race condition not guaranteed to happen every time
		// lock the lower numbered fork first to avoid the deadlock scenario
		if philosophers.leftFork > philosophers.rightFork {
			// for Aristotle the left fork is higher numbered than right fork so we lock the right fork first
			forks[philosophers.rightFork].Lock()
			forks[philosophers.leftFork].Lock()
			fmt.Printf("\t%s has picked up right fork.\n", philosophers.name)
			fmt.Printf("\t%s has picked up left fork.\n", philosophers.name)
		} else {
			// for all others left fork is lower numbered so it is locked first
			forks[philosophers.leftFork].Lock()
			forks[philosophers.rightFork].Lock()
			fmt.Printf("\t%s has picked up left fork.\n", philosophers.name)
			fmt.Printf("\t%s has picked up right fork.\n", philosophers.name)
		}

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
