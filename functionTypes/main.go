package main

import (
	. "fmt" // revive:disable dot-imports
)

/* The `calc` variable defined in Listing 9-4 can be assigned any value that
* matches its type, which means any function that has the right number and type
* of arguments and results */
var calc func(float64) float64

/*
Using function types can be verbose and repetitive, which produces code that
can be hard to read and maintain. Go supports type aliases, which can be used
to assign a name to a function signature so that the parameter and result types
are not specified every time the function type is used, as shown in Listing
9-8.
You don’t have to use aliases for function types, but they can simplify code
and make the use of a specific function signature easier to identify.
*/
type Calc func(float64) float64

//var calc Calc

func main() {
	Println("Function Types\n")
	Println(`
Understanding Function Types
Functions have a data type in Go, which means they can be assigned to variables
and used as function parameters, arguments, and results. Listing 9-4 shows a
simple use of a function data type.` + "\n")

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	Println("Using a Function Data Type in the main.go File in the functionTypes Folder")
	for product, price := range products {
		Printf("A function is assigned to  `calc`? %t; %T\n", calc != nil, calc)
		if price > 100 {
			// To assign a specific function to a variable, the function’s name is used.
			calc = calcWithTax
		} else {
			calc = calcWithoutTax
		}
		Printf("A function is assigned to  `calc`? %t; %T\n", calc != nil, calc)
		// Once a function has been assigned to a variable, it can be invoked
		// as though the variable’s name was the function’s name.
		totalPrice := calc(price)
		Printf("Product: %s; Price: %.2f\n", product, totalPrice)
	}

	Println("\nUsing Functions as Arguments")

	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}

	Println("\nUsing Functions as Results")
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}

	Println("\nCreating Function Type Aliases")
	for product, price := range products {
		printPrice_(product, price, selectCalculator_(price))
	}

	Println("\nUsing the Literal Function Syntax\nUsing Functions Values Directly")

	for product, price := range products {
		printPrice_(product, price, selectCalculator__(price))
	}

	Println("\nUsing Functions Values Directly\nUnderstanding Function Closure")
	for product, price := range products {
		printPrice_(product, price, func(price float64) float64 {

			if price < 101 {
				return price
			}
			return price + (price * 0.2)
		})
	}

}

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax(price float64) float64 {
	return price
}

// Funcs as arguments
func printPrice(product string, price float64, calc func(float64) float64) {
	Printf("Product: %s; Price: %.2f\n", product, calc(price))
}

// Using Functions as Results
func selectCalculator(price float64) func(float64) float64 {
	if price < 101 {
		return calcWithoutTax
	}
	return calcWithTax
}

// Creating Function Type Aliases
func printPrice_(product string, price float64, calc Calc) {
	Printf("Product: %s; Price: %.2f\n", product, calc(price))
}

// Using Functions as Results
func selectCalculator_(price float64) Calc {
	if price < 101 {
		return calcWithoutTax
	}
	return calcWithTax
}

// Using the Literal Function Syntax
// Using Functions Values Directly
/* The literal syntax omits a name so that the func keyword is followed by the
* parameters, the result type, and the code block, as shown in Figure 9-7.
* Because the name is omitted, functions defined this way are called anonymous
* functions.
Note go does not support arrow functions, where functions are expressed more concisely using the => operator, without the func keyword and a code block surrounded by braces. in go, functions must always be defined with the keyword and a body.
 Literal functions can also be used with the short variable declaration syntax:
 withTax := func...
*/

// Understanding Function Closure
/* Functions defined using the literal syntax can reference variables from the
* surrounding code, a feature known as closure. The closure feature allows a
* function to access variables—and parameters—in the surrounding code. */
/* Forcing Early Evaluation
Evaluating closures when the function is invoked can be useful, but if you want
to use the value that was current when the function was created, then copy the
value, as shown in Listing 9-16. */
func selectCalculator__(price float64) Calc {
	if price < 101 {
		return func(price float64) float64 {
			return price
		}

	}
	return func(price float64) float64 {
		return price + (price * 0.2)
	}
}
