package main

import (
	. "fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func main() {
	Printfln("Hello, Math and Sorting")
	val1 := 279.00
	val2 := 48.95
	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	var ev int8
	ev = int8(math.RoundToEven(val2))
	Printfln("%d,%T", ev, ev)
	Printfln("%s, %s", reflect.TypeOf(ev), Sprintf("%d", ev))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))
	Printfln("math.MaxInt16: %d", math.MaxInt16)

	Println("\nGenerating Random Numbers")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Printfln("Value %v: %v", i, rand.Int())
	}

	Println("\nGenerating a Random Number Within a Specific Range")
	for i := 5; i < 15; i++ {
		Printfln("Value %v: %v", i, rand.Intn(i))
	}

	for i := 5; i < 15; i++ {
		min := i - 5
		Printfln("Value min:%v ; max: %v: random in range %v", min, i, IntRange(min, i))
	}

	Println("\nShuffling Elements")
	var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}
	Printfln("initial order: %v", names)
	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})

	for i, name := range names {
		Printfln("Index: %d Name: %s", i, name)
	}

	println("\nSorting Data\nSorting Number and String Slices")
	Printfln("The elements are sorted?:%v", sort.StringsAreSorted(names))
	sort.Strings(names)
	Printfln("The elements are sorted?:%v", sort.StringsAreSorted(names))
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)
	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)
	ints = []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)

	Printfln("\nSearching Sorted Data")
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v", indexOf4)
	Printfln("Index of 3: %v", indexOf3)
	Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
	Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3)

	Println("\nSorting Custom Data Types")
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlices(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}

	Println("\nSorting Using Different Fields")

	ProductSlicesByName(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
	products[0], products[1] = products[1], products[0]
	Println("Specifying the Comparison Function\n", products)
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}
