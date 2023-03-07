package main

import (
	"fmt"
	"log"
	"sort"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.LUTC)
}

func sortAndTotal(vals []int) (sorted []int, total int) {
	logger := log.New(log.Writer(), "sortAndTotal: ",
		log.Flags()|log.Lmsgprefix)
	logger.Printf("Invoked with %v values", len(vals))
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	logger.Printf("Sorted data: %v", sorted)
	for _, val := range sorted {
		total = total + val
	}
	logger.Printf("Total: %v", total)
	return
}

func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := sortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)

	fmt.Println("\n Benchmarking Code")
	fmt.Println("\n Logging Data")
	log.Print("Sorted Data: ", sorted)
	log.Print("Total: ", total)
}
