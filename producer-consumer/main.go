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
				color.Red("Error closing channel!", err)
			}
		}
	}

	// printout the ending message
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification
	// try to make pizzas

	for {
		// try to make a pizza
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
		// decision
	}
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {

		color.Magenta("Received order #%d!\n", pizzaNumber)

		random := rand.Intn(12) + 1
		msg := ""
		success := false
		delay := rand.Intn(5) + 1

		color.Yellow("Making Pizza #%d, It will take %d seconds....\n", pizzaNumber, delay)
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

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
