package main

type Named interface{ GetName() string }

type Person struct{ PersonName string }

func (p *Person) GetName() string             { return p.PersonName }
func (dp *DiscountedProduct) GetName() string { return dp.Name }
