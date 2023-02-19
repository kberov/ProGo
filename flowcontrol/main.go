package main

import (
	"fmt"
	"strconv"
)

func main() {
	kayakPrice := 275.00
	if kayakPrice > 100 &&
		kayakPrice < 500 {
		fmt.Println("Price is greater than 100 and less than 500")
	}
	/* Go allows an if statement to use an initialization statement, which is
	 * executed before the if statement’s expression is evaluated. The
	 * initialization statement is restricted to a Go simple statement, which
	 * means—in broad terms—that the statement can define a new variable, assign a
	 * new value to an existing variable, or invoke a function. */
	priceString := "275_"
	if kayakPr, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPr)
	} else {
		fmt.Printf("Error:%s, Price: %d\n", err, kayakPr)
	}

	/* The for keyword is used to create loops that repeatedly execute
	* statements. The most basic for loops will repeat indefinitely unless
	* interrupted by the break keyword, as shown in Listing 6-11. (The return
	* keyword can also be used to terminate a loop.) */
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}
	counter = 0
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
		// if (counter > 3) {
		// break
		// }
	}
	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter:", counter)
		// counter++
	}
	//RE-CREATING DO...WHILE LOOPS
	for counter := 0; true; counter++ {
		fmt.Println("Counter:", counter)
		if counter > 3 {
			break
		}
	}

	/* The continue keyword can be used to terminate the execution of the for
	* loop’s statements for the current value and move to the next iteration.
	* */
	for counter := 0; counter <= 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter:", counter)
	}
	product := "KayakИсега"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}
	// enumerate only indexes
	for index := range product {
		fmt.Println("Index:", index)
	}
	// only values
	for _, character := range product {
		fmt.Println("Character:", string(character))
	}

	/* The range keyword can also be used with the built-in data structures that
	 * Go provides—arrays, slices, and maps—all of which are described in Chapter
	 * 7, including examples using the for and range keywords. */
	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}

	fmt.Println("\nUsing switch Statements")
	/* A switch statement provides an alternative way to control execution
	* flow, based on matching the result of an expression to a specific value,
	* as opposed to evaluating a true or false result, as shown in Listing
	* 6-19. This can be a concise way to perform multiple comparisons,
	* providing a less verbose alternative to a complex if/elseif/else
	* statement. */
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("K at position", index)
			fmt.Printf("Good %c\n", character)
		case 'y':
			fmt.Println("y at position", index)
		}
	}

	/* Go switch statements do not fall through automatically, but multiple
	* values can be specified with a comma-separated list, as shown in Listing
	* 6-20. */
	for index, character := range product {
		switch character {
		case 'И', 'с', 'е', 'г', 'а':
			fmt.Printf("Good! %c at position %d\n", character, index)
		case 'K', 'k':
			fmt.Println("K or k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}

	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}

	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}
	/* A switch statement can be defined with an initialization statement,
	* which can be a helpful way of preparing the comparison value so that it
	* can be referenced within case statements.  */
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Printf("%d/2; Prime value:%d\n", counter, val)
		default:
			fmt.Printf("%d/2; Non-prime value:%d\n", counter, val)
		}
	}
	// Listing 6-26. Using Expressions in a switch Statement in the main.go
	// File in the flowcontrol Folder
	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}
	fmt.Println("\nUsing Label Statements")
	counter = 0
target:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
}
