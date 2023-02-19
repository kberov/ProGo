package main

import (
	. "fmt"
)

type Product struct {
	name, category string
	price          float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (self *Product) printDetails() {
	Printf("*Product(%s): %#v\nprice with VAT: %.2f\n",
		self.name, self, self.calcTax(0.2, 100))
}

func printDetails(self *Product) {
	Printf("*Product(%s): %#v\n", self.name, self)
}

/*
Note: one effect of this feature is that value and pointer types are considered

	the same when it comes to method overloading, meaning that a method named
	printDetails whose receiver type is Product will conflict with a printDetails
	method whose receiver type is *Product.

	func printDetails(self Product) {
		Printf("*Product(%s): %#v\n", self.name, self)
	}
*/
func (self *Product) calcTax(rate, threshold float64) float64 {
	if self.price > threshold {
		return self.price + self.price*rate
	}
	return self.price
}

// methods to implement the Expense interface
func (p Product) getName() string {
	return p.name
}

// the `annual bool` parameter is not used but needs to conform to the
// signature defined in the Expense interface
func (p Product) getCost(_ bool) float64 {
	return p.price
}
