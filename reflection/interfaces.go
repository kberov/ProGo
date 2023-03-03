package main

import "fmt"

type NamedItem interface {
	GetName() string
	unexportedMethod()
}

// © Adam Freeman 2022
// A. Freeman, Pro Go, https://doi.org/10.1007/978-1-4842-7355-5_29
// 785 Chapter 29 ■ Using Reflection, Part 3
type CurrencyItem interface {
	GetAmount() string
	currencyName() string
}

func (p *Product) GetName() string {
	return p.Name
}
func (c *Customer) GetName() string {
	return c.Name
}
func (p *Product) GetAmount() string {
	return fmt.Sprintf("$%.2f", p.Price)
}
func (p *Product) currencyName() string {
	return "USD"
}
func (p *Product) unexportedMethod() {}
