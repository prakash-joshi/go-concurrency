package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var (
	pizzasMade   int
	pizzasFailed int
	totalPizzas  int
)

// Producer is a struct that holds two channels
// data: information for a given pizza order
// quit: to handle end of processing
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a struct that defines an order
// pizzaNumber: describes order no
// message: message describing what happened to the order
// success: indicates if the order was successfully completed.
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random value generator
	rand.NewSource(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in background
	go pizzeria(pizzaJob)

	// create and run the consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is angry!!!")
			}
		} else {
			color.Cyan("Done making pizzas!!!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// printout the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the Day!")
	color.Cyan("We made %d pizzas, failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, totalPizzas)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day....")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a great day!")
	}

}

// pizzeria is a goroutine that runs in the background
// calls makePizza to make one order at a time
// rune until it receives something on quit channel.
// the quit channel does not receives anything until the consumer executes the Close() methods
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something.
	for {
		// try to make a pizza
		currentPizza := makePizza(i)

		// decision
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we send something to the data channel -- a chan PizzaOrder)
			case pizzaMaker.data <- *currentPizza:

			// we want to quit, so send pizzMaker.quit to the quitChan (a chan error)
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

// makePizza attempts to make pizza
// generate random number from 1-12 and decide wether the pizza will be made or not
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {

		color.Magenta("Received order #%d!\n", pizzaNumber)

		random := rand.Intn(12) + 1
		msg := ""
		success := false
		delay := rand.Intn(5) + 1

		color.Yellow("Making Pizza #%d, It will take %d seconds....\n", pizzaNumber, delay)

		// delay for a bit
		time.Sleep(time.Duration(time.Duration(delay) * time.Second))

		if random <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
			pizzasFailed++
		} else if random < 5 {
			msg = fmt.Sprintf("*** The pizza #%d got burnt in oven due to excess heat", pizzaNumber)
			pizzasFailed++
		} else {
			msg = fmt.Sprintf("Pizza Order #%d is ready!", pizzaNumber)
			pizzasMade++
			success = true
		}
		totalPizzas++

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

// Close method is to close a channel when you are done with it
// i.e, something is pushed to the quit channel
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
