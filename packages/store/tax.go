package store

//import (
//	. "fmt"
//)

const defaultTaxRate float64 = 0.2
const minThreshold = 10

/* Using Package Initialization Functions
* Each code file can contain an initialization function that is executed
* only when all packages have been loaded and all other initialization—such
* as defining constants and variables—has been done. The most common use
* for initialization functions is to perform calculations that are
* difficult to perform or that require duplication to perform, as shown in
* Listing 12-17. */

var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

// The initialization function is called init, and it is defined without
// parameters and a result. The init function is called automatically and
// provides an opportunity to prepare the package for use.

// The init function is not a regular Go function and cannot be invoked
// directly. And, unlike regular functions, a single file can define multiple
// init functions, all of which will be executed.
func init() {
	//Println("in init() of store/tax.go")
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + price*defaultTaxRate
	}
}

type taxRate struct{ rate, threshold float64 }

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(prod *Product) (price float64) {
	if prod.price > taxRate.threshold {
		price = prod.price + (prod.price * taxRate.rate)
	} else {
		price = prod.price
	}

	if max, ok := categoryMaxPrices[prod.Category]; ok && price > max {
		price = max
	}

	return
}
