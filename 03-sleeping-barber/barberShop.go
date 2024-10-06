package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		// isSleeping keeps a track if the barber is awake or sleeping
		isSleeping := false
		color.Cyan("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there is no client waiting, the barber goes to sleep.
			if len(shop.ClientChan) == 0 {
				color.Red("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			// if the channel has some value isShopOpen value will be true
			// if the channel is closed the isShopOpen value will be false
			// use isShopOpen value to determine if the shop is open or closed
			client, isShopOpen := <-shop.ClientChan

			if isShopOpen {
				if isSleeping {
					color.Red("%s wakes %s up.", client, barber)
					isSleeping = false
				}

				// cut the hairs
				shop.cutHair(barber, client)

			} else {

				// the shop is closed send the barber home
				shop.sendBarberHome(barber)
				return

			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {

	color.Cyan("%s is cutting %s's hair", barber, client)
	time.Sleep(cutDuration)
	color.Cyan("%s is finished cutting %s's hair", barber, client)

}

func (shop *BarberShop) sendBarberHome(barber string) {

	color.Magenta("%s is going home", barber)
	shop.BarbersDoneChan <- true
}
