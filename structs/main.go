package main

import (
	"encoding/json"
	. "fmt"
	"strings"
)

func main() {
	Println("Hello Structs")
	Println("\nDefining and Using a Struct")
	kayak := Product{name: "Kayak", category: "Watersports", price: 275}
	Printf("Name: %s; Category: %s; Price: %.2f\n",
		kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	Println("\nUsing a Struct Value\nChanged price: ", kayak.price)

	Println("\nPartially Assigning Struct Values")

	partialKayak := Product{name: "Kayak", category: "Watersports"}
	Printf("Name: %s; Category: %s; Price: %.2f\n",
		partialKayak.name, partialKayak.category, partialKayak.price)
	var lifejacket Product
	Println("Name is zero value:", lifejacket.name == "")
	Println("Category is zero value:", lifejacket.category == "")
	Println("Price is zero value:", lifejacket.price == 0)

	Println("\nUsing Field Positions to Create Struct Values")
	// using pointer to the struct
	kayakP := &Product{"Kayak", "Watersports", 270.00}
	Printf("Name: %s\nCategory: %s\nPrice: %.2f\n",
		kayakP.name, kayakP.category, kayakP.price)

	Println("\nDefining Embedded Fields")
	stockItem := &StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.01},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	Printf("Stock Item Name, %s\n", stockItem.Product.name)
	Printf("Stock Item Name, %s\n", stockItem.name) //can be accessed as if not embedded
	Printf("Stock Item Name, %s\n", stockItem.Alternate.name)
	Printf("Stock Item Count, %d\n", stockItem.count)

	Println("\nComparing Struct Values")
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	Println("p1 == p2:", p1 == p2)
	Println("p1 == p3:", p1 == p3)

	Println("\nConverting Between Struct Types")
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	Println("prod == item:", prod == Product(item))
	Printf("Product: %v\nItem: %v\n Product(item): %v\n", prod, item, Product(item))

	Println("\nDefining Anonymous Struct Types")
	item = Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod)
	writeName(item)

	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	if err == nil {
		Printf("Anonimous struct as JSON: %s\n", builder.String())
	} else {
		Printf("An error occured: %s\n", err)
	}

	Println("\nCreating Arrays, Slices, and Maps Containing Struct Values")
	arrayOfStockLevels := [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	Println("Array:", arrayOfStockLevels[0].Product.name)
	// The struct type can be omitted when populating arrays, slices, and maps
	// with struct values
	slice := []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	Println("Slice:", slice[0].Product.name)
	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	Printf("Map: %#v\n", kvp)

	Println("\nUnderstanding Structs and Pointers")
	// Assigning a struct to a new variable or using a struct as a function
	// parameter creates a new value that copies the field values, as
	// demonstrated in Listing 10-16.
	p1 = Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p2 = p1
	p1.name = "Original Kayak"
	Println("P1:", p1.name)
	Println("P2:", p2.name)

	Println("Listing 10-17. Using a Pointer to a Struct:")
	p2_ := &p1 //p2_ type is *Product, meaning a pointer to a Product value.
	p1.name = "Original Kayak"
	Println("P1:", p1.name)
	// We have to use parentheses to follow the pointer to the struct value and
	// then read the value of the name field.
	Println("P2:", (*p2_).name)

	Println("\nUnderstanding the Struct Pointer Convenience Syntax")
	calcTax(&kayak)
	Printf("Product: %#v\n", kayak)
	// Go follows the pointer
	calcTax_(&kayak)
	Printf("Product: %#v\n", kayak)

	Println("\nUnderstanding Pointers to Values")
	kayak_ := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax_(kayak_)
	Printf("Product: %#v\n", kayak_)

	Println("Listing 10-21. Using Pointers Directly")
	kayak_ = calcTax__(&Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	})
	Printf("Product: %#v\n", kayak_)

	Println("\nUnderstanding Struct Constructor Functions")
	products := []*Product{
		NewProduct("Kayak", "Watersports", 275),
		NewProduct("Hat", "Skiing", 42.50),
	}

	for i, p := range products {
		i++
		Printf("%d: %#v\n", i, p)
	}

	Println("\nUsing Pointer Types for Struct Fields")
	acme := &Supplier{"Acme Co", "New York"}

	products_ := []*Product_{
		NewProduct_("Kayak", "Watersports", 275, acme),
		NewProduct_("Hat", "Skiing", 42.50, acme),
	}

	for i, p := range products_ {
		i++
		Printf("%d: %#v\nSupplier: %#v\n", i, p, p.Supplier)
	}

	Println("\nUnderstanding Pointer Field Copying")
	p1__ := NewProduct_("Kayak", "Watersports", 275, acme)
	p2__ := *p1__
	p1__.name = "Original Kayak"
	p1__.Supplier.name = "BoatCo"
	for _, p := range []Product_{*p1__, p2__} {
		Println("Name:", p.name, "Supplier:",
			p.Supplier.name, p.Supplier.city)
	}
	p2__ = copyProduct_(p1__)
	p1__.name = "Original Kayak1"
	p1__.Supplier.name = "BoatCo1"
	for _, p := range []Product_{*p1__, p2__} {
		Println("Name:", p.name, "Supplier:",
			p.Supplier.name, p.Supplier.city)
	}

	Println("\nUnderstanding Zero Value for Structs and Pointers to Structs")
	var prod_ Product
	var prodPtr *Product
	Println("Value:", prod_.name, prod_.category, prod_.price)
	Println("Pointer:", prodPtr)
	var prod__ Product_
	var prodPtr_ *Product_
	Println("Value:", prod__.name, prod__.category, prod__.price, prod__.Supplier.name)
	Println("Pointer:", prodPtr_)

} //End of main()

// The type can be defined anywhere, because main is executd after the whole
// code is compiled. The type definition is visible from main during execution.
type Product struct {
	name, category string
	price          float64
	// Listing 10-11. Adding an Incomparable Field in the main.go File in the structs Folder
	// Compile time errors:
	// too few values in Product{…}
	// invalid operation: p1 == p2 (struct containing []string cannot be compared)
	// otherNames []string
}

type Product_ struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct{ name, city string }

/*
UNDERSTANDING STRUCT TAGS
The struct type can be defined with tags, which provide additional information
about how a field should be processed. Struct tags are just strings that are
interpreted by the code that processes struct values, using the features
provided by the reflect package. See Chapter 21 for an example of how struct
tags can be used to change the way that structs are encoded in JSOn data, and
see Chapter 28 for details of how to access struct tags yourself.
*/

/*
USING THE NEW FUNCTION TO CREATE STRUCT VALUES
You may see code that uses the built-in new function to create struct values,
like this:
...
var lifejacket = new(Product)
...
the result is a pointer to a struct value whose fields are initialized with
their type’s zero value. this is
equivalent to this statement:
...
var lifejacket = &Product{}
...
these approaches are interchangeable, and choosing between them is a matter of
preference.
*/

/*
Defining Embedded Fields
If a field is defined without a name, it is known as an embedded field, and it
is accessed using the name of its type, as shown in Listing 10-8.
*/
type StockLevel struct {
	Product
	Alternate Product
	count     int
}

/*
Converting Between Struct Types
A struct type can be converted into any other struct type that has the same
fields, meaning all the fields have the same name and type and are defined in
the same order, as demonstrated in Listing 10-12.
*/
type Item struct {
	name     string
	category string
	price    float64
}

func writeName(val struct {
	name, category string
	price          float64
}) {
	Printf("Name: %s\n", val.name)
}

/* Creating Arrays, Slices, and Maps Containing Struct Values
The struct type can be omitted when populating arrays, slices, and maps with
struct values, as shown in Listing 10-15.` */

// Understanding the Struct Pointer Convenience Syntax

func calcTax(product *Product) {
	if (*product).price > 100 {
		(*product).price += (*product).price * 0.2
	}
}

// To simplify this type of code, Go will follow pointers to struct fields
// without needing an asterisk character.
// The asterisk and the parentheses are not required, allowing a pointer to a
// struct to be treated as though it were a struct value, as illustrated by
// Figure 10-7.
func calcTax_(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}

}

// Listing 10-21. Using Pointers Directly

func calcTax__(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

// Understanding Struct Constructor Functions
// Constructor functions return struct pointers, and the address operator is
// used directly with the literal struct syntax.
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Using Pointer Types for Struct Fields

func NewProduct_(name, category string, price float64, suppl *Supplier) *Product_ {
	return &Product_{name, category, price, suppl}
}

func copyProduct_(product *Product_) Product_ {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

// Understanding Zero Value for Structs and Pointers to Structs
// The zero value for a struct type is a struct value whose fields are assigned
// their zero type. The zero value for a pointer to a struct is nil, as
// demonstrated in Listing 10-27.
