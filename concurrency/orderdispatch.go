package main

import (
	. "fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

/* Restricting Channel Direction by using <- after keyword chan! The location
* of the arrow specifies the direction of the channel. When the arrow follows
* the chan keyword, as it does in Listing 14-23, then the channel can be used
* only to send. The channel can be used to receive only if the arrow precedes
* the chan keyword (<-chan, for example). Attempting to receive from a
* send-only (and vice versa) channel is a compile-time error. */
func DispatchOrders(channel chan<- DispatchNotification) {
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(5) + 5
	Printf("Order count: %d\n", orderCount)
	custLastIndex, productsLastIndex := len(Customers)-1, len(ProductList)-1
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(custLastIndex)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(productsLastIndex)],
		}
		//if i == 1 {
		//	notification := <-channel
		//	Printf("Reading Customer %s from the channel to which this goroutine "+
		//		" has just sent\n", notification.Customer)
		//	/* It is easy to spot this issue in the example code, but I usually
		//	* make this mistake when an if statement is used to conditionally
		//	* send additional values through the channel. The result, however,
		//	* is that the function receives the message it has just sent,
		//	* removing it from the channel. */
		//}
		time.Sleep(time.Millisecond * 750)
	}
	Println("Closing  the DispatchNotification Channel")
	close(channel)
}
