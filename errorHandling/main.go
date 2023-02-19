package main

import (
	. "fmt"
)

type CategoryCountMessage struct {
	Category      string
	Count         int
	TerminalError interface{}
}

/*
Recovering from Panics in Go Routines
A panic works its way up the stack only to the top of the current goroutine, at which point it causes
termination of the application. This restriction means that panics must be recovered within the code that a
goroutine executes, as shown in Listing 15-17.
*/
func processCategories(categories []string, outChannel chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			Println(arg)
			outChannel <- CategoryCountMessage{
				TerminalError: arg,
			}
			close(outChannel)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChannel <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
		close(outChannel)
	}
}

func main() {
	//	defer func() {
	//		if arg := recover(); arg != nil {
	//			// type assertion: arg.(error) - `is this an error`
	//			if err, ok := arg.(error); ok {
	//				Println("Recovering from Panics!..")
	//				Println("... realizing that the error is not recoverable")
	//				panic(err)
	//				// type assertion: arg.(string) - `is this a string`
	//			} else if str, ok := arg.(string); ok {
	//				Printf("Message: %s\n", str)
	//			} else {
	//				Println("Panic recovered!")
	//			}
	//		}
	//	}()
	// `()` are required to invoke — rather than just define — the anonymous
	// function.

	categories := []string{"Watersports", "Chess", "Running"}
	Println("Dealing with Recoverable Errors\n\nReporting Errors via Channels\n",
		"Using the Error Convenience Functions")
	channel := make(chan CategoryCountMessage, 10)
	go processCategories(categories, channel)
	for message := range channel {
		// To ignore the error, use the blank identifier `_`:
		// total, _ := Products.TotalPrice(cat)
		if message.TerminalError == nil {
			Printf("%s Total: %d\n", message.Category, message.Count)
		} else {
			Println("Dealing with Unrecoverable Errors: Use `panic`.")
			Println("A terminal error occured!")
			//panic(message.CategoryError)
			//Printf("No such category '%s'!\n", message.Category)
		}
	}
}
