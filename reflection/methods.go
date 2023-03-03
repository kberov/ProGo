package main

func (p Purchase) calcTax(taxRate float64) float64 {
	return p.Price * taxRate
}

// 786 Chapter 29 ■ Using Reflection, Part 3
func (p Purchase) GetTotal() float64 {
	return p.Price + p.calcTax(.20)
}
