package main

import (
	. "fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	//Printf("in CalcStoreTotal ... data<ProductData<ProductGroup<Product>>>: %#v\n", data["Soccer"][2])
	var storeTotal float64
	Println("\nReturning Results from Goroutines")
	/* Channels are strongly typed, which means that they will carry values of
	* a specified type or interface. The type for a channel is the chan
	* keyword, followed by the type the channel will carry */
	Println("\nUsing a Buffered Channel - buffer size 2")
	// See explanation below this title in the book
	var channel chan float64 = make(chan float64, 2)
	var count int
	for cat, group := range data {
		Printf("\nCreating Additional Goroutine for %s. Channel: %v Num: %d\n", cat, channel, count)
		count++
		go group.TotalPrice(cat, channel)
	}
	// Wait for the goroutines :) :)...
	//time.Sleep(1 * time.Second) // sleep for one second
	for i := 0; i < len(data); i++ {
		lch, cch := len(channel), cap(channel)
		Printf("Waiting for a result FROM the channel:"+
			" %v Num: %d with len(%d) and cap(%d)\n",
			channel, i, lch, cch)
		value := <-channel
		Printf("Received a result FROM the channel:"+
			" %v Num: %d Value: %.2f with len(%d) and cap(%d)\n",
			channel, i, value, lch, cch)
		storeTotal += value
		time.Sleep(time.Second)
	}
	Printf("Total: %s\n", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(cat string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		Printf("%s product: '%s'\n", cat, p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}
	Printf("%s subtotal: %s\n", cat, ToCurrency(total))
	Printf("Sending a Result for %s TO the channel...\n", cat)
	resultChannel <- total
	Printf("Sent a Result for %s TO the channel...\n", cat)
}
