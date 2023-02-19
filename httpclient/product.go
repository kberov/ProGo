package main

type Product struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"` // `json:"price,string"`
}

var Kayak = Product{
	Name:     "Kayak", //`json: "name"`
	Category: "Watersports",
	Price:    279,
}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func (p *Product) AddTax() float64 {
	return p.Price * 1.2
}

func (p *Product) ApplyDiscount(amount float64) float64 {
	return p.Price - amount
}
