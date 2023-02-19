package store

type SpecialDeal struct {
	Name string
	*Product
	price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
	return &SpecialDeal{name, p, p.price - discount}
}

func (d *SpecialDeal) GetDetails() (string, float64, float64) {
	return d.Name, d.price, d.Price(0)
}

/* The new Price method stops Go from promoting the Product method. */
func (d *SpecialDeal) Price(taxRate float64) float64 {
	return d.price
}
