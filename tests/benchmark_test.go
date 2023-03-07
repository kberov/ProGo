package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 250
	data := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		b.StartTimer()
		sortAndTotal(data)
	}
}

var sizes = []int{10, 100, 250}

func BenchmarkSortSubBenches(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array Size %v", size),
			func(subB *testing.B) { theBench(subB, size) })
	}
}

func theBench(subB *testing.B, size int) {
	data := make([]int, size)
	// 850 Chapter 31 ■ Unit Testing, Benchmarking, and Logging
	subB.ResetTimer()
	for i := 0; i < subB.N; i++ {
		subB.StopTimer()
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		subB.StartTimer()
		sortAndTotal(data)
	}
}
