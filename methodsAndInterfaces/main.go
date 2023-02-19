package main

import (
	. "fmt"
)

/*
Understanding Method Overloading
Go does not support method overloading, where multiple methods can be defined
with the same name but different parameters. Instead, each combination of
method name and receiver type must be unique, regardless of the other
parameters that are defined. In Listing 11-7, I have defined methods that have
the same name but different receiver types.
*/

/*
Defining and Using Methods
Methods are functions that can be invoked via a value and
are a convenient way of expressing functions that operate on
a specific type.
*/
/*
Understanding Method Overloading
Go does not support method overloading, where multiple methods can be defined
with the same name but different parameters. Instead, each combination of
method name and receiver type must be unique, regardless of the other
parameters that are defined.
*/
type Supplier struct {
	name, city string
}

// func (self *Supplier) printDetails() {
// 	Printf("*Supplier (%s): %#v\n", self.name, self)
// }

func (self *Supplier) printDetails(showName ...bool) {
	if len(showName) > 0 && showName[0] {
		Printf("*Supplier (%s): %#v\n", self.name, self)
	} else {
		Printf("*Supplier: %#v\n", self)
	}
}

/*
Understanding Pointer and Value Receivers
A method whose receiver is a pointer type can also be invoked through a regular
value of the underlying type, meaning that a method whose type is *Product, for
example, can be used with a Product value, as shown in Listing 11-9.  Listing
11-9.
*/
/*
Defining Methods for Type Aliases
Methods can be defined for any type defined in the current package. I explain
how to add packages to a project in Chapter 12, but for this chapter, there is
a single code file containing a single package, which means that methods can be
defined only for types defined in the main.go file.  But this doesn’t limit
methods to just structs, because the type keyword can be used to create aliases
to any type, and methods can be defined for the alias. (I introduced the type
keyword in Chapter 9 as a way to simplify dealing with function types.) Listing
11-11 creates an alias and a method
*/
type ProductList []Product

func (products *ProductList) calcCategoryTotals() (totals map[string]float64) {
	totals = make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

/*
Defining and Using Interfaces
It is easy to imagine a scenario where the Product and Service types defined in
the previous section are used together...

Defining an Interface
This problem is addressed using interfaces, which describe a set of methods
without specifying the implementation of those methods. If a type implements
all the methods defined by the interface, then a value of that type can be used
wherever the interface is permitted. The first step is to define an interface,
as shown Listing 11-16.
*/
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

/*
Implementing an Interface
To implement an interface, all the methods specified by the interface must be
defined for a struct type, as shown in Listing 11-17.
See files product.go and service.go

Using an interface
...
Variables whose type is an interface have two types: the static type and the
dynamic type. The static type is the interface type. The dynamic type is the
type of value assigned to the variable that implements the interface, such as
Product or Service in this case. The static type never changes—the static type
of an Expense variable is always Expense, for example—but the dynamic type can
change by assigning a new value of a different type that implements the
interface.

Using an Interface in a Function
*/
func calcTotal(expenses []Expense) (total float64) {
	for _, it := range expenses {
		total += it.getCost(true)
	}
	return
}

/*
Using an Interface for Struct Fields
Interface types can be used for struct fields, which means that fields can be
assigned values of any type that implements the methods defined by the
interface, as shown in Listing 11-21.
*/
type Account struct {
	accountNumber int
	expenses      []Expense
}

type Person struct {
	name, city string
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		Println("Product:", value.name, "Price:", value.price)
	case *Product:
		Println("Product Pointer:", value.name, "Price:", value.price)
	case Service:
		Println("Service:", value.description, "Price:",
			value.monthlyFee*float64(value.durationMonths))
	case Person:
		Println("Person:", value.name, "City:", value.city)
	case *Person:
		Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		Println("Built-in type:", value)
	default:
		Printf("Default:%#v\n", value)
	}
}

func processItems(items ...interface{}) {
	for _, item := range items {
		processItem(item)
	}
}

/*
	Understanding the Effect of Pointer Method Receivers

The methods defined by the Product and Service types have value receivers,
which means that the methods will be invoked with copies of the Product or
Service value. This can be confusing, so Listing 11-22 provides a simple
example.
*/
func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	//aKayak := Product{"Kayak", "Watersports", 275}
	for _, p := range products {
		Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}
	Println("\nDefining and Using Methods")
	Println("Using function printDetails(p)")
	for _, p := range products {
		printDetails(p)
	}

	products = []*Product{
		newProduct("Kayak", "Watersports", 276),
		newProduct("Lifejacket", "Watersports", 48.96),
		newProduct("Soccer Ball", "Soccer", 19.51),
	}

	Println("Using method p.printDetails()")
	Println("Defining Method Parameters and Results")
	for _, p := range products {
		p.printDetails()
	}
	Println("\nUnderstanding Method Overloading")
	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails(true)
	}

	Println("Understanding Pointer and Value Receivers")
	kayak := Product{"Kayak", "Watersports", 271}
	/* The kayak variable is assigned a Product value but is used with the
	* printDetails method, whose receiver is *Product. Go takes care of the
	* mismatch and invokes the method seamlessly. The opposite process is also
	* true so that a method that receives a value can be invoked using a
	* pointer, as shown in Listing 11-10. */
	kayak.printDetails()
	Println("\nINVOKING METHODS VIA THE RECEIVER TYPE!!!")
	(*Product).printDetails(&kayak)

	Println("\nDefining Methods for Type Aliases")
	products_ := ProductList(getProducts())
	/* The result from the getProducts function is []Product, which is
	* converted to ProductList with an explicit conversion, allowing the method
	* defined on the alias to be used. */
	for category, total := range products_.calcCategoryTotals() {
		Printf("Category: %s, Total: %.2f\n", category, total)
	}

	Println("\nPutting Types and Methods in Separate Files")
	kayak = Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50, []string{""}}
	Println("Product:", kayak.name, "Price:", kayak.price)
	Println("Service:", insurance.description, "Price:",
		insurance.monthlyFee*float64(insurance.durationMonths))

	Println("\n\nDefining and Using Interfaces\nImplementing an Interface\n",
		"Using an Interface")
	expenses := []Expense{
		Product{"Kay", "Watersp.", 275},
		Service{"BoatIns", 12, 89.50, []string{""}},
	}
	/* The for loop deals only with the static type—Expense—and doesn’t know
	* (and doesn’t need to know) the dynamic type of those values. The use of
	* the interface has allowed me to group disparate dynamic types together
	* and use the common methods specified by the static interface type. */
	for _, exp := range expenses {
		Printf("Expense: %s Cost: %.2f\n", exp.getName(), exp.getCost(true))
	}
	Println("\nUsing an Interface in a Function")
	Printf("Total Expenses: %.2f\n", calcTotal(expenses))

	Println("\nUsing an Interface for Struct Fields")
	acc := Account{
		accountNumber: 12345,
		expenses:      expenses,
	}

	Printf("Total Expenses for account #%d: %.2f\n",
		acc.accountNumber, calcTotal(acc.expenses))

	Println("\nUnderstanding the Effect of Pointer Method Receivers")
	product := Product{"Kayak", "Watersports", 275}
	/* Using a pointer means that a reference to the Product value is assigned
	* to the Expense variable, but this does not change the interface variable
	* type, which is still Expense. */
	var expense Expense = &product
	product.price = 100
	/* Now the change to the price field is reflected in the result from the
	* getCost method. It can also be counterintuitive because the variable type
	* is always Expense, regardless of whether it is assigned a Product or
	* *Product value */
	Println("Product field value:", product.price)
	Println("Expense method result:", expense.getCost(false))

	Println("\nComparing Interface Values")
	/* Interface values can be compared using the Go comparison operators, as shown
	 * in Listing 11-26. Two interface values are equal if they have the same
	 * dynamic type and all of their fields are equal. */

	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}
	//var e3 Expense = Service{description: "Boat Cover"}
	/* The first two Expense values are not equal. That’s because the dynamic
	* type for these values is a pointer type, and pointers are equal only if
	* they point to the same memory location. The second two Expense values are
	* equal because they are simple struct values with the same field values.
	 */
	//var e4 Expense = Service{description: "Boat Cover"}
	Println("e1 == e2", e1 == e2)
	// panic: runtime error: comparing uncomparable type main.Service
	// Println("e3 == e4", e3 == e4)

	Println("\nPerforming Type Assertions")
	/* We use type assertion to access the dynamic Service value from a slice
	* of Expense interface types. Once I have a Service value to work with, I
	* can use all the fields and methods defined for the Service type, and not
	* just the methods that are defined by the Expense interface. */
	expenses = []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses {
		// A type assertion is performed by applying a period after a value,
		// followed by the target type in parentheses: expense.(Service)
		if s, ok := expense.(Service); ok {
			Printf("Service: %s, Price: %.2f\n",
				s.description, s.monthlyFee*float64(s.durationMonths))
		} else {
			Printf("Expence: %s, Cost: %.2f \n", expense.getName(), expense.getCost(true))
		}
	}

	Println("\nSwitching on Dynamic Types")
	for _, expense := range expenses {
		// The switch statement uses a special type assertion that uses the
		// type keyword, as illustrated in Figure 11-7.
		switch value := expense.(type) {
		case Service:
			Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case *Product:
			// the methods in the product.go file use pointer receivers!!!
			Println("Product:", value.name, "Price:", value.price)
		default:
			Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}

	Println("\nUsing the Empty Interface")

	var expence Expense = &Product{"Kayak", "Watersports", 48.59}
	data := []any{ // synonim of interface{}
		expence,
		Product{"LifeJ.", "WaterS.", 48.95},
		Service{"BoatC.", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "NY"},
		"Just a string",
		100,
		true,
	}

	for _, item := range data {
		switch val := item.(type) {
		case Product:
			Printf("Product: %s Price: %.2f\n", val.name, val.price)
		case *Product:
			Printf("*Product: %s Price: %.2f\n", val.name, val.price)
		case Person:
			Println("Person:", val.name, "City:", val.city)
		case *Person:
			Println("Person Pointer:", val.name, "City:", val.city)
		case string, bool, int:
			Println("Built-in type:", val)
		default:
			Println("Default:", val)

		}
	}

	Println("\nUsing the Empty Interface for Function Parameters")
	/* The empty interface can be used as the type for a function parameter,
	* allowing a function to be called with any value, as shown in Listing
	* 11-33. */

	for _, item := range data {
		processItem(item)
	}
	/* The empty interface can also be used for variadic parameters, which
	* allows a function to be called with any number of arguments, each of
	* which can be any type, as shown in Listing 11-34. */
	Println("\n...can also be used for variadic parameters")
	processItems(data...)
}
