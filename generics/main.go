package main

import (
	. "fmt"
)

func main() {
	Println("Hello Generics - This chapter is from 'A tour of Go' _ not in the book.")
	Println("\nType parameters")

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	i := 15
	Printf("Find %d in %v: %d\n", i, si, Index(si, i))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	s := "hello"
	Printf("Find '%s' in %v: %d\n", s, ss, Index(ss, s))
	Println("\nGeneric types")
	// In addition to generic functions, Go also supports generic types.
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
