package main

const NumberOfPizzas = 10

var (
	pizzasMade   int
	pizzasFailed int
	totalPizzas  int
)

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random value generator

	// print out a message

	// create a producer

	// run the producer in background

	// create and run the consumer

	// printout the ending message
}
