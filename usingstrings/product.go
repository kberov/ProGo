package main

import . "fmt"

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    275,
}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

// String reperesantation for Product when used with Printf %+v
/*
The fmt package supports custom struct formatting through an interface named
Stringer that is defined as follows:
	type Stringer interface {
	    String() string
	}

This method implements Stringer.

Tip: If you define a GoString method that returns a string, then your type will
conform to the GoStringer interface, which allows custom formatting for the %#v
verb.
*/
func (p Product) String() string {
	return Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}
