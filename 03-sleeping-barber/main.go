// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
// open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
// waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
// sleep until a new client arrives. So the rules are as follows:
//
//   - if there are no customers, the barber falls asleep in the chair
//   - a customer must wake the barber if he is asleep
//   - if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
//     sits in an empty chair if it's available
//   - when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
//     and falls asleep if there are none
//   - shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is
//     empty
//   - after the shop is closed and there are no clients left in the waiting area, the barber
//     goes home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Dijkstra.
//
// The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
// semaphores (mutexes) is not needed.
package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/fatih/color"
)

var (
	seatingCapacity = 10
	arrivalRate     = 100
	cutDuration     = 1000 * time.Millisecond
	timeOpen        = 10 * time.Second
)

func main() {
	// seed our random number generator
	rand.NewSource(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientsChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientsChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	shop.addBarber("Pappu")
	shop.addBarber("Kallu")
	shop.addBarber("Gullu")
	shop.addBarber("Dholu")
	shop.addBarber("Bholu")

	// start the barbershop as a goroutine
	shopClosed := make(chan bool)
	shopClosing := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		shopClosed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average client arrival rate
			randomMilliSeconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliSeconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed

	// time.Sleep(5 * time.Second)
	<-shopClosed
}
