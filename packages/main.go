package main

import (
	// When using a dot import, you must ensure that the names of the features
	// imported from the package are not defined in the importing package.
	// I must NOT define any of the fmt identifiers - e.g. Println,Printf etc.
	// A dot import uses a period as the package alias.
	// The dot import allows me to access all the public identifiers of the
	// package without using a prefix.
	. "fmt"
	_ "packages/data" // just to execute the init() function
	cFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
)

func main() {
	Println("Creating and Using Packages")

	Println("Hello, Packages and Modules")
	prod := store.Product{Name: "Kayak", Category: "Watersports"}
	Printf("Name: %s\nCategory: %s\n", prod.Name, prod.Category)
	// prod.price undefined (type store.Product has no field or method price)
	// Println(product.price)

	Println("\nUnderstanding Package Access Control")
	produ := store.NewProduct("Kayak", "Watersports", 279)
	Printf("Name: %s\nCategory: %s\nPrice: %.2f\n",
		produ.Name, produ.Category, produ.Price())

	Println("\nAdding Code Files to Packages")
	Printf("Now the price is calculated by store.calcTax. p.Price(): %.2f\n", produ.Price())
	Println("\nDealing with Package Name Conflicts")
	Printf("Price: %s\n", cFmt.ToCurrency(produ.Price()))

	Println("\nCreating Nested Packages")
	/* Packages can be defined within other packages, making it easy to break
	* up complex features into as many units as possible. */
	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*produ},
	}
	Printf("Name: %s\nTotal: %s\n",
		cart.CustomerName, cFmt.ToCurrency(cart.GetTotal()))
	Println("\nUsing Package Initialization Functions")
	// See store/tax.go

	Println("\nImporting a Package Only for Initialization Effects")
	// use _ as package alias when importing a package.

	Println("\nUsing External Packages")
	/* Projects can be extended using packages developed by third parties.
	* Packages are downloaded and installed using the go get command. Run the
	* command shown in Listing 12-21 in the packages folder to add a package to
	* the example project */
	color.Green(Sprintf("Name: %s\n", cart.CustomerName))
	color.Cyan(Sprintf("Total: %s\n", cFmt.ToCurrency(cart.GetTotal())))

	Printf("\nManaging External Packages\n")
	/* The go get command adds dependencies to the go.mod file, but these are not removed automatically if the external package is no longer required. To update the go.mod file to reflect the change, run the `go mod tidy` command.*/
	//	Println("Name:", cart.CustomerName)
	//	Println("Total:", cFmt.ToCurrency(cart.GetTotal()))
}
