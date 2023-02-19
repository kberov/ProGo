package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, Collections!")
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
	fmt.Println("\nUsing the Array Literal Syntax")
	/* The number of elements specified with the literal syntax can be less
	* than the capacity of the array. Any position in the array for which a
	* value is not provided will be assigned the zero value for the array type.
	* */
	names_ := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names_)
	var coords [3][3]int
	coords[1][2] = 10
	fmt.Println(coords)
	// cannot use names (variable of type [3]string) as [4]string value in variable declaration (compile)
	// var otherArray [4]string = names
	// NB: The capacity is part of the type!!!
	// Let the compiler infer the capacity by using `...` instead of a number for literal array creation. page 152.
	// A COPY of names is created!!!
	var otherArray = names

	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("otherArray:", otherArray)
	// using pointer!
	anotherArray := &names
	names[0] = "Boat"
	fmt.Println("names:", names)
	fmt.Println("anotherArray:", *anotherArray)

	fmt.Println("\nComparing Arrays")
	// Arrays are equal if they are of the same type and contain equal elements
	// in the same order
	_names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := _names == moreNames
	fmt.Println("comparison:", same)

	fmt.Println("\nEnumerating Arrays")
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}
	// Listing 7-11. Discarding the Current Index in the main.go File in the
	// collections Folder
	/* The underscore is known as the blank identifier and is used when a
	* feature returns values that are not subsequently used and for which a
	* name should not be assigned. */
	for _, value := range names {
		fmt.Println("Value:", value)
	}

	fmt.Println("\nWorking with Slices")
	/* The best way to think of slices is as a variable-length array because
	* they are useful when you don’t know how many values you need to store or
	* when the number changes over time. One way to define a slice is to use
	* the built-in make function, as shown in Listing 7-12. */
	boats := make([]string, 3)
	boats[0] = "Kayak"
	boats[1] = "Lifejacket"
	boats[2] = "Paddle"
	fmt.Println(boats)
	// literal assignment
	sailors_things := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(sailors_things)

	fmt.Println("\nAppending Elements to a Slice")
	sailors_things = append(sailors_things, "Hat", "Gloves")
	fmt.Println(sailors_things)
	appendedNames := append(sailors_things, "Hat", "Gloves")
	sailors_things[0] = "Canoe"
	fmt.Println("sailors_things:", sailors_things)
	fmt.Println("appendedNames:", appendedNames)

	// Allocating Additional Slice Capacity
	Names := make([]string, 3, 6)
	Names[0] = "Kayak"
	Names[1] = "Lifejacket"
	Names[2] = "Paddle"
	fmt.Println("len:", len(Names))
	fmt.Println("cap:", cap(Names))
	/* Caution if you define a slice variable but don’t initialize it, then the
	* result is a slice that has a length of zero and a capacity of zero, and
	* this will cause an error when an element is appended to it. */
	amendedNames := append(Names, "Hat", "Gloves")
	Names[0] = "Canoe"
	fmt.Println("Names:", Names)
	fmt.Println("amendedNames:", amendedNames)
	/* The result of the append function is a slice whose length has increased
	* but is still backed by the same underlying array. The original slice
	* still exists and is backed by the same array, with the effect that there
	* are now two views onto a single array, as shown in Figure 7-13. */

	fmt.Println("\nAppending One Slice to Another")
	_moreNames := []string{"Hat Gloves"}
	amendedNames = append(Names, _moreNames...)
	fmt.Println("amendedNames:", amendedNames)

	fmt.Println("\nCreating Slices from Existing Arrays")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	fmt.Println("products:", products)
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)

	fmt.Println("\nAppending Elements When Using Existing Arrays for Slices")
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	someNames = append(someNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))

	fmt.Println("\nSpecifying Capacity When Creating a Slice from an Array")
	/* Ranges can include a maximum capacity, which provides some degree of
	* control over when arrays will be duplicated, as shown in Listing 7-23. */
	_someNames := products[1:3:3]

	fmt.Println("_someNames len:", len(_someNames), "cap:", cap(_someNames))

	fmt.Println("\nCreating Slices from Other Slices")
	_allnames := products[1:]
	some_Names := _allnames[1:3]
	_allnames = append(_allnames, "Gloves")
	_allnames[1] = "Canoe"
	fmt.Println("some_Names:", some_Names)
	fmt.Println("_allnames", _allnames)

	fmt.Println("\nUsing the copy Function to Ensure Slice Array Separation")
	/* The copy function can be used to duplicate an existing slice, selecting
	* some or all the elements but ensuring that the new slice is backed by its
	* own array, as shown in Listing 7-25. */
	products_ := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames_ := products_[1:]
	someNames_ := make([]string, 2)
	copy(someNames_, allNames_)
	fmt.Println("someNames_:", someNames_)
	fmt.Println("allNames_", allNames_)

	fmt.Println("\nCopying Slices with Different Sizes")
	_products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacement_products := []string{"Canoe", "Boots"}
	copy(_products, replacement_products)
	fmt.Println("_products:", _products)

	fmt.Println("\nDeleting Slice Elements")
	/* There is no built-in function for deleting slice elements, but this
	 * operation can be performed using the ranges and the append function, as
	 * Listing 7-30 demonstrates. */
	// this is a new slice
	products_without_deleted_element := append(products[:2], products[3:]...)
	fmt.Println("products_without_deleted_element:", products_without_deleted_element)

	fmt.Println("\nEnumerating Slices")
	for index, value := range products[1:] {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println("\nSorting")
	sort.Strings(_products)
	for index, value := range _products {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println("\nWorking with Maps")
	/* Maps are a built-in data structure that associates data values with keys.
	 * Unlike arrays, where values are associated with sequential integer locations,
	 * maps can use other data types as keys, as shown in Listing 7-36. */
	mproducts := make(map[string]float64, 10)
	mproducts["Kayak"] = 279
	mproducts["Lifejacket"] = 48.95
	// The number of items stored in the map is obtained using the built-in len
	// function, like this:
	fmt.Println("Map size:", len(mproducts))
	fmt.Println("Price:", mproducts["Kayak"])
	// NB! The zero value for the map’s value type is returned if the map
	// doesn’t contain the key.
	fmt.Println("Price:", mproducts["Hat"])
	fmt.Printf("map: %#v\n", mproducts)
	fmt.Printf("Using the Map Literal Syntax\n")
	things := map[string]string{
		"нещо":  "първо нещо",
		"друго": "второ нещо",
	}
	fmt.Printf("map: %#v\n", things)
	fmt.Printf("Checking for Items in a Map\n")
	if val, exists := things["treto"]; exists {

		fmt.Println("treto:", val)
	} else {

		fmt.Println("treto:", val, " - nema treto")
	}
	fmt.Printf("Removing Items from a Map\n")
	delete(things, "нещо")
	fmt.Printf("map: %#v\n", things)
	if value, ok := things[`нещо`]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value for key `нещо`")
	}
	// Enumerating
	for key, value := range mproducts {
		fmt.Println("Key:", key, "Value:", value)
	}
	// in order
	keys := make([]string, 0, len(mproducts))
	for key := range mproducts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", mproducts[key])
	}

	fmt.Printf("\nDual Nature of strings\n")
	var price string = "$48.95"
	var currency byte = price[0]
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
	// Slicing a string produces another string, but an explicit conversion is
	// required to interpret the byte as the character it represent
	fmt.Println("Currency strng(price[0]):", string(price[0]))
	fmt.Println("len(price):", len(price))
	price = "€48.95"
	fmt.Println("len(price):", len(price))
	fmt.Println("\nConverting a String to Runes")
	/* The rune type represents a Unicode code point, which is essentially a
	* single character. To avoid slicing strings in the middle of characters,
	* an explicit conversion to a rune slice can be performed, as shown in
	* Listing 7-48. */
	price_rune := []rune(price)
	// slicing a rune slice produces another rune slice.
	currency_rune := price_rune[0]
	fmt.Println("len(price_rune):", len(price_rune))
	/* As explained in Chapter 4, the rune type is an alias for int32, which
	* means that printing out a rune value will display the numeric value used
	* to represent the character. This means, as with the byte example
	* previously, I have to perform an explicit conversion of a single rune
	* into a string, like this: */
	fmt.Println("price_rune:", price_rune)
	fmt.Println("string(price_rune):", string(price_rune))
	fmt.Println("currency_rune:", currency_rune)
	fmt.Println("string(currency_rune):", string(currency_rune))

	fmt.Println("price_rune[1:]:", price_rune[1:])
	fmt.Println("string(price_rune[1:]):", string(price_rune[1:]))
	fmt.Println("\nEnumerating Strings")
	for index, char := range price {
		fmt.Println(index, char, string(char))
	}

	/* If you want to enumerate the underlying bytes without them being
	* converted to characters, then you can perform an explicit conversion to a
	* byte slice, as shown in Listing 7-50. */
	fmt.Println(`Enumerating the Bytes in the String`)
	for index, char := range []byte(price) {
		fmt.Println(index, char, string(char))
	}
}
