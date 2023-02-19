package main

import (
	. "fmt"
	"time"
)

/*
Table 14-1. Putting Goroutines and Channels in Context
Question Answer
What are they? Goroutines are lightweight threads created and managed by the Go runtime.
Channels are pipes that carry values of a specific type.
Why are they useful?  Goroutines allow functions to be executed concurrently, without needing to deal
with the complications of operating system threads. Channels allow goroutines to
produce results asynchronously.
How are they used?
Goroutines are created using the go keyword. Channels are defined as data types.
Are there any pitfalls Care must be taken to manage the direction of channels. Goroutines that share data
or limitations?
require additional features, which are described in Chapter 14.
Are there any
alternatives?
Goroutines and channels are the built-in Go concurrency features, but some
applications can rely on a single thread of execution, which is created by default to
execute the main function.

Table 14-2. Chapter Summary
Problem Solution Listing
Execute a function asynchronouslyCreate a goroutine7
Produce a result from a function executed
asynchronouslyUse a channel10, 15, 16, 22–26
Send and receive values using a channelUse arrow expressions11–13
Indicate that no further values will be sent over Use the close function
a channel17–20
Enumerate the values received from a channel
Use a for loop with the range keyword 21
Send or receive values using multiple channels Use a select statement
27–32
*/
func main() {
	Println("Using Goroutines and Channels")
	Println("\nUnderstanding How Go Executes Code")
	// Reread the explanation below this title in the book!!!

	//	Println("main function started")
	//	CalcStoreTotal(products)
	//	Println("main function complete")

	dispatchChannel := make(chan DispatchNotification, 100)
	// Println("\nRestricting Channel Argument Direction")
	//	var sendOnlyCh chan<- DispatchNotification = dispatchChannel
	//	var receiveOnlyCh <-chan DispatchNotification = dispatchChannel
	//	go DispatchOrders(sendOnlyCh)
	//	receiveDispatches(receiveOnlyCh)
	// or
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	// receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
	/*
	   Using Select Statements
	   The select keyword is used to group operations that will send or receive
	   from channels, which allows for complex arrangements of goroutines and
	   channels to be created. There are several uses for select statement.
	   Receiving Without Blocking
	   The simplest use for select statements is to receive from a channel
	   without blocking, ensuring that a goroutine won’t have to wait when the
	   channel is empty.
	   A select statement has a similar structure to a switch statement, except
	   that the case statements are channel operations. When the select
	   statement is executed, each channel operation is evaluated until one
	   that can be performed without blocking is reached. The channel operation
	   is performed, and the statements enclosed in the case statement are
	   executed. If none of the channel operations can be performed, the
	   statements in the default clause are executed.
	*/
	// Receiving from Multiple Channels
	ch1 := make(chan *Product, 2)
	ch2 := make(chan *Product, 2)
	go enumerateProducts(ch1, ch2)
	time.Sleep(time.Second)
	for p := range ch1 {
		Printf("Received product: %s via ch1 \n", p.Name)
	}
	for p := range ch2 {
		Printf("Received product: %s via ch2 \n", p.Name)
	}
	/*
			openChannels := 2
			for {
				select {
				case details, ok := <-dispatchChannel:
					if ok {
						Printf("Dispatch to %s: %d x %s\n",
							details.Customer, details.Quantity, details.Name)
					} else {
						Println("Dispach channel has been closed")
						dispatchChannel = nil
						openChannels--
					}
				case product, ok := <-productChannel:
					if ok {
						Printf("Product: %s\n", product.Name)
					} else {
						Println("Product channel has been closed")
						productChannel = nil
						openChannels--
					}
				default:
					if openChannels == 0 {
						goto alldone
					}
					Println("-- No message ready to be received")
					time.Sleep(time.Millisecond * 500)
				}
			}
		alldone:
			Println("All values received")
	*/
	//	for {
	//		if details, isChannelOpen := <-dispatchChannel; isChannelOpen {
	//			Printf("Dispatch to %s: %d x %s\n",
	//				details.Customer, details.Quantity, details.Name)
	//		} else {
	//			Println("Channel has been closed")
	//			break
	//		}
	//	}

}

func receiveDispatches(channel <-chan DispatchNotification) {
	for orderDetails := range channel {
		// orderDetails.Name==orderDetails.Product.Name
		Printf("Dispatch to %s: %d x %s\n",
			orderDetails.Customer, orderDetails.Quantity, orderDetails.Name)
	}
	Println("Channel has been closed")
}

/*
Sending Without Blocking
A select statement can also be used to send to a channel without blocking, as
shown in Listing 14-30.
*/
func enumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		// The select statement determines when the send operation would block
		// and invokes the default clause instead.
		select {
		case channel1 <- p:
			Printf("Sent product: %s via channel 1\n", p.Name)
		case channel2 <- p:
			Printf("Sent product: %s via channel 2\n", p.Name)
			// If there is no `default` clause, the select statement will block
			// until one of the channels can receive a value.
		default:
			Printf("Discarding product %s\n", p.Name)
			time.Sleep(time.Second)
		}
	}
	Println("Closing *Product channels.")
	close(channel1)
	close(channel2)
}
