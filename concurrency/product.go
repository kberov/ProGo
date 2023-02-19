/*
This file defines a custom type named Product, along with type aliases that I
use to create a map that organizes products by category. I use the Product type
in a slice and a map, and I rely on an init function, described in Chapter 12,
to populate the map from the contents of the slice, which is itself populated
using the literal syntax. This file also contains a ToCurrency function that
formats float64 values into dollar currency strings, which I will use to format
the results in this chapter.
*/
package main

import str "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + str.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range ProductList {
		if _, ok := products[p.Category]; ok {
			products[p.Category] = append(products[p.Category], p)
		} else {
			products[p.Category] = ProductGroup{p}
		}
	}
}
