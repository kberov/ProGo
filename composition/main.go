package main

import (
	"composition/store"
	. "fmt"
)

/*
Table 13-1. Putting Type and Interface Composition in Context
Question Answer
What is it? Composition is the process by which new types are created by combining
structs and interfaces.
Why is it useful? Composition allows types to be defined based on existing types.
How is it used? Existing types are embedded in new types.
Are there any pitfalls or
limitations? Composition doesn’t work in the same way as inheritance, and care must be
taken to achieve the desired outcome.
Are there any alternatives?  Composition is optional, and you can create entirely independent types.
Table 13-2. Chapter Summary
Problem Solution Listing
Compose a struct typeAdd an embedded field7-9, 14–17
Build on an already composed typeCreate a chain of embedded types10–13
Compose an interface typeAdd the name of the existing interface to the new
interface definition25–26
*/
func main() {
	Println("Type and Interface Composition")
	Println("\nDefining the Base Type") // See store/product.go - type ...
	Println("\nDefining a Constructor") // See store/product.go NewProduct...
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product{kayak, lifejacket} {
		Printf("Name: %s Category: %s Price: %.2f\n",
			p.Name, p.Category, p.Price(0.2))
	}

	Println("\nComposing Types")
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}

	for _, b := range boats {
		Printf("Conventional: %s Direct: b.Name: %s Price: %.2f\n",
			b.Product.Name, b.Name, b.Price(0.2))
	}

	Println("\nCreating a Chain of Nested Types")
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	/* Go promotes fields from the nested Boat and Product types so they can be
	* accessed through the top- level RentalBoat type, which allows the Name
	* field to be read in Listing 13-11. Methods are also promoted to the
	* top-level type, which is why I can use the Price method, even though it
	* is defined on the *Product type, which is at the end of the chain.*/
	for _, r := range rentals {
		Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2),
			"Captain:", r.Captain)
	}

	Println("\nUsing Multiple Nested Types in the Same Struct")
	// See type Crew in rentalboats.go

	Println("\nUnderstanding When Promotion Cannot Be Performed")
	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()
	Println("Name:", Name)
	Println("Price field:", price)
	/* The third result is the one that can cause problems. Go can promote the
	* Price method, but when it is invoked, it uses the price field from the
	* Product and not the SpecialDeal.  It is easy to forget that field and
	* method promotion is just a convenience feature. So we MUST overwrite the
	* Price method for type SpecialDeal. */
	Println("Price method:", Price)

	Println("\nUnderstanding Promotion Ambiguity")
	kayak = store.NewProduct("Kayak", "Watersports", 279)
	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}
	/*
		The OfferBundle type has two embedded fields, both of which have Price methods.
		Go cannot differentiate between the methods, and the code in Listing 13-17
		produces the following error when it is compiled (BUT at RUNTIME, when
		the method price is invoked):
		.\main.go:22:33: ambiguous selector bundle.Price
	*/
	//	bundle := OfferBundle{
	//		store.NewSpecialDeal("Weekend Special", kayak, 50),
	//		store.NewProduct("Lifrejacket", "Watersports", 48.95),
	//	}
	//	Println("Price:", bundle.Price(0))

	Println("\nUnderstanding Composition and Interfaces")
	//	products := map[string]*store.Product{
	//		"Kayak": store.NewBoat("Kayak", 279, 1, false),
	//		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	//	}
	//	for _, p := range products {
	//		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	//	}
	//
	// ./main.go:103:12: cannot use store.NewBoat("Kayak", 279, 1, false) (value of type *store.Boat) as type *store.Product in map literal
	/*
		This can seem similar to writing classes in other languages, but there is an
		important difference, which is that each composed type is distinct and cannot
		be used where the types from which it is composed are required, as shown in
		Listing 13-18.

		The Go compiler will not allow a Boat to be used as a value in a slice where
		Product values are required.  In a language like C# or Java, this would be
		allowed because Boat would be a subclass of Product, but this is not how Go
		deals with types.
	*/

	Println("\nUsing Composition to Implement Interfaces")
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for k, p := range products {
		Println("Key:", k, "Price:", p.Price(0.2))
	}

	Println("\nUnderstanding the Type Switch Limitation")
	// ... case statements that specify multiple types will match values of all
	// of those types but will not perform type assertion.
	// main.go|139 col 30| item.Name undefined (type store.ItemForSale has no field or method Name) (typecheck)
	//	for key, p := range products {
	//		switch item := p.(type) {
	//		case *store.Product, *store.Boat:
	//			fmt.Println("Name:", item.Name, "Category:", item.Category,
	//				"Price:", item.Price(0.2))
	//		default:
	//			fmt.Println("Key:", key, "Price:", p.Price(0.2))
	//		}
	//	}
	// A type assertion is performed by the case statement when a single type
	// is specified, albeit it can lead to duplication as each type is
	// processed.
	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		case *store.Boat:
			Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		default:
			Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			Println("Name:", item.GetName(), "Category:", item.GetCategory(),
				"Price:", item.Price(0.2))
		default:
			Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	Println("\nComposing Interfaces (See loop above)")
	/*
		One interface can enclose another, with the effect that types must
		implement all the methods defined by the enclosing and enclosed
		interfaces. Interfaces are simpler than structs, and there are no
		fields or method to promote. The result of composing interfaces is a
		union of the method defined by the enclosing and enclosed types. See
		Describable interface in store/product.go. Now we removed the assertion
		item.(store.ItemForSale) in the loop above.
		A value of any type that implements the Describable interface must have
		a Price method because of the composition performed in Listing 13-25,
		which means the method can be called without a potentially risky type
		assertion.
	*/
}
