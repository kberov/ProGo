package main

import (
	"fmt"
	. "fmt"
)

func main() {
	Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
	//println("Kayak is:", Kayak.String())

	Println("Formatting Strings")
	Printf("Product: %v, Price: $%6.2f\n", Kayak.Name, Kayak.Price)
	name, _ := getProductName(1)
	Println(name)
	_, err := getProductName(10)
	Println(err.Error())

	Println("\nUnderstanding the Formatting Verbs")
	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	Println("\nControlling Struct Formatting")
	Printfln("Value with fields: %+v", Kayak)

	Println("\nFormatting Numbers")
	number := 250
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)

	Println("\nUsing the Floating-Point Formatting Verbs")
	float := 279.00
	Printfln("Decimalless with exponent: %b", float)
	Printfln("Decimal with exponent: %e", float)
	Printfln("Decimal without exponent: %f", float)
	Printfln("Hexadecimal: %x, %X", float, float)
	Printfln("Decimal without exponent: >>%8.2f<<", float)
	Printfln("Decimal without exponent: >>%.2f<<", float)
	Printfln("Sign: >>%+.2f<<", float)
	Printfln("Zeros for Padding: >>%010.2f<<", float)
	Printfln("Right Padding: >>%-8.2f<<", float)

	Println("\nUsing the String and Character Formatting Verbs")
	име := "Краси"
	Printfln("String: %s", име)
	// to convert a string to runes, simply create a rune slice from it:
	// runes:=[]rune(име)
	Printfln("Character: %c", []rune(име)[0])
	Printfln("Unicode: %U", []rune(име)[0])
	Println("\nUsing the Pointer Formatting Verb")
	Printfln("Pointer: %p", &име)

	Println("\n\nScanning Strings..............Hurray!!!")
	println("Use `< ./skanln_in.txt ` as input.")
	Println("Please enter a text to scan:")
	var _name string
	var category string
	var price float64
	/*
	 */
	numvals, err := Scan(&_name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", numvals)
		Printfln("Name: %v, Category: %v, Price: %.2f", _name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
	// Scan into a slice of values - needs the folowing trick.
	// because the string slice can’t be properly decomposed for use with the
	// variadic parameter.
	vals := make([]string, 3)
	ivals := make([]any, 3)
	for i := 0; i < len(vals); i++ {
		ivals[i] = &vals[i]
	}
	fmt.Print("Enter text to scan: ")
	fmt.Scan(ivals...)
	Printfln("Name: %v", vals)
	// By default, scanning treats newlines in the same way as spaces, acting as
	// separators between values.
	Println("\nDealing with Newline Characters")
	/*
		The Scan function doesn’t stop looking for values until after it has received
		the number it expects and the first press of the Enter key is treated as a
		separator and not the termination of the input. The functions whose name ends
		with ln in Table 17-12, such as Scanln, change this behavior. Listing 17-20
		uses the Scanln function.
	*/
	// if Enter is pressed before all three values are collecte,
	// we see the error: `Error: unexpected newline`.
	// If more values are entered: `Error: expected newline`
	nums, err := Scanln(&_name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", nums)
		Printfln("Name: %v, Category: %v, Price: %.2f", _name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}

	Println("\nUsing a Different String Source")
	/*
		The functions described in Table 17-12 scan strings from three sources: the
		standard input, a reader (described in Chapter 20), and a value provided as an
		argument. Providing a string as the argument is the most flexible because it
		means the string can arise from anywhere.
	*/
	source := "Lifejacket Watersports 48.95"
	n, err := Sscan(source, &name, &category, &price)
	Printfln("%#v\n", []any{n, err, name, category, price})

	Println("\nUsing a Scanning Template")
	source = "Product Lifejacket Watersports 42.95"
	template := "Product %s %s %f"
	nu, err := Sscanf(source, template, &name, &category, &price)
	Printfln("%#v\n", []any{nu, err, name, category, price})
}

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = Errorf("Error for index %v", index)
	}
	return
}
func Printfln(template string, values ...interface{}) {
	Printf(template+"\n", values...)
}

// Formatting and Scanning Strings

/*
In this chapter, I describe the standard library features for formatting and scanning strings. Formatting is the
process of composing a new string from one or more data values, while scanning is the process of parsing
values from a string. Table 17-1 puts formatting and scanning strings in context.
Table 17-1. Putting Formatting and Scanning Strings in Context

Question Answer
What are they? Formatting is the process of composing values into a string. Scanning is the process of parsing a string for the values it contains.
Why are they useful? Formatting a string is a common requirement and is used to produce strings for everything from logging and debugging to presenting the user with information.  Scanning is useful for extracting data from strings, such as from HTTP requests or user input.
How are they used? Both sets of features are provided through functions defined in the fmt package.
Are there any pitfalls or limitations? The templates used to format strings can be hard to read, and there is no built-in function that allows a formatted string to be created to which a newline character is appended automatically.
Are there any alternatives? Larger amounts of text and HTML content can be generated using the template features described in Chapter 23.

Table 17-2 summarizes the chapter.
Table 17-2. Chapter Summary
Problem
Solution
Listing
Combine data values to form a string -> Use the basic formatting functions provided by the fmt package5, 6
Specify the structure of a string -> Use the fmt functions that use formatting templates and use the formatting verbs7–9, 11–18
Change the way custom data types are represented -> Implement the Stringer interface10
Parse a string to obtain the data values it contains -> Use the scanning functions provided by the fmt package

The Println and Fprintln functions add spaces between all the values, but the
Print and Fprint functions only add spaces between values that are not strings.
This means that the pairs of functions in Table 17-3 differ in more than just
adding a newline character, as shown in Listing 17-5.
*/
