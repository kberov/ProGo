package main

import (
	. "fmt"
)

//type CategoryError struct {
//	requestedCat string
//}
//
//func (e *CategoryError) Error() string {
//	return Sprintf("Category `%s` does not exist!", e.requestedCat)
//}

type ChannelMessage struct {
	Category      string
	Total         float64
	CategoryError error
}

func (slice ProductSlice) TotalPrice(category string) (total float64,
	err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = Errorf("No such category :%v\n", category)
	}
	return
}

func (slice ProductSlice) TotalPriceAsync(categories []string,
	channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
