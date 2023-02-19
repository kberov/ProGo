package main

import "fmt"

func main() {
	fmt.Println("Hello, Functions!")
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)

	fmt.Println(`Omitting Parameter Names`)
	printPrice_("Kayak", 275, 0.2)
	printPrice_("Lifejacket", 48.95, 0.2)
	printPrice_("Soccer Ball", 19.50, 0.15)
	printPrice__("Soccer Ball", 19.50, 0.15)

	fmt.Println("\nDefining Variadic Parameters")
	printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	printSuppliers("Lifejacket", []string{"Sail Safe Co"})
	printSuppliers_("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers_("Lifejacket", "Sail Safe Co")
	printSuppliers_("Soccer Ball")

	fmt.Println("Using Slices as Values for Variadic Parameters")
	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers_("Kayak", names...)

	fmt.Println("\nUsing Pointers as Function Parameters")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues(val1, val2)
	fmt.Println("After calling function", val1, val2)
	fmt.Println("Before calling function", val1, val2)
	swapValues_(&val1, &val2)
	fmt.Println("After calling function", val1, val2)

	fmt.Printf("\nDefining and Using Function Results\n")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
		"Hat":        10.01,
	}
	for product, price := range products {
		fmt.Printf("Product: %s; Price: %.2f\n", product, calcTax(price))
	}

	fmt.Println("\nReturning Multiple Function Results")
	fmt.Println("Before calling function", val1, val2)
	val1, val2 = swapValues__(val1, val2)
	fmt.Println("After calling function", val1, val2)

	fmt.Printf("\nUsing Multiple Results Instead of Multiple Meanings\n")
	for product, price := range products {
		taxAmount, taxDue := calcTax_(price)
		if taxDue {
			fmt.Printf("Product: %s; Tax: %.2f\n", product, taxAmount)
		} else {
			fmt.Printf("Product: %s; No tax due.\n", product)
		}
	}

	fmt.Println("\nUsing Named Results")
	tot1, tax1 := calcTotalPrice(products, 10)
	fmt.Printf("Total 1: %.2f; Tax 1: %.2f\n", tot1, tax1)
	tot2, tax2 := calcTotalPrice(nil, 10)
	fmt.Printf("Total 2: %.2f; Tax 2: %.2f\n", tot2, tax2)

	fmt.Println("\nUsing the Blank Identifier to Discard Results")
	fmt.Printf("and\nUsing the defer Keyword\n")
	_, total := calcTotalPriceNoCount(products)
	fmt.Printf("Total: %.2f\n", total)

} // end main()

// Note! Go does not support optional parameters or default values for parameters.
// The type can be omitted when adjacent parameters have the same type, as
// shown in Listing 8-6.
func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Printf("Product: %s; Price: %.2f; Tax: %.3f\n", product, price, taxAmount)
}

// Omitting Parameter Names
/*
The underscore is known as the blank identifier, and the result is a parameter
for which a value must be provided when the function is called but whose value
cannot be accessed inside the function’s code block.  This may seem like an odd
feature, but it can be a useful way to indicate that a parameter is not used
within a function, which can arise when implementing the methods required by an
interface. */
func printPrice_(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func printPrice__(string, float64, float64) {
	fmt.Println("No parameters")
}

// Variadic Parameters
func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

// Variadic parameters allow a function to receive a variable number of
// arguments more elegantly, as shown in Listing 8-10
func printSuppliers_(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
		return
	}

	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

// Using Pointers as Parameters

func swapValues(first, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

// Listing 8-15. Defining a Function with Pointers
func swapValues_(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)

}

// Defining and Using Function Results
func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

// Returning Multiple Function Results
func swapValues__(first, second int) (int, int) {
	return second, first
}

// Using Multiple Results Instead of Multiple Meanings
func calcTax_(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false

}

// Using Named Results
/* A function’s results can be given names, which can be assigned values during
* the function’s execution. When execution reaches the return keyword, the
* current values assigned to the results are returned, as shown in Listing
* 8-22. */
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax_(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

// Using the Blank Identifier to Discard Results
// Using the defer Keyword
func calcTotalPriceNoCount(products map[string]float64) (count int, total float64) {
	fmt.Println("Func calcTotalPriceNoCount started")
	defer fmt.Println("First `defer` call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second `defer` call")
	fmt.Println("Func calcTotalPriceNoCount is about to return")
	return
}
