// Package store provides types and methods commonly required for
// online sales
// The name specified by the package statement should match
// the name of the folder in which the code files are created,
// which is store in this case.
package store

// taxRate redeclared in this block
//
//	store/product.go:8:5: other declaration of taxRate
// var taxRate = newTaxRate(0.25, 20)
var standartTax = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
	Name, Category string
	price          float64
}

/* Go examines the first letter of the names given to the features in a code
* file, such as types, functions, and methods. If the first letter is
* lowercase, then the feature _can be used only within the *package*_ that defines
* it. Features are exported for use outside of the package by giving them an
* uppercase first letter.
The name of the struct type in Listing 12-4 is
* Product, which means the type can be used outside the store package. The
* names of the Name and Category fields also start with an uppercase letter,
* which means they are also exported. The price field has a lowercase first
* letter, which means that it can be accessed only within the store package.
* */

// Listing 12-7. Defining Methods in the product.go File in the store Folder
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (self *Product) Price() float64            { return standartTax.calcTax(self) }
func (self *Product) SetPrice(newPrice float64) { self.price = newPrice }
