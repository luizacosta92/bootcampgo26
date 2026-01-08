package main

type Size string
const (
	Small Size = "small"
	Medium Size = "medium"
	Large Size = "large"
)
type Product struct {
	Size  Size
	Price float64
}

type Pricer interface {
	CalculatePrice() float64
}
func (p *Product) CalculatePrice() float64 {
	switch p.Size {
	case "small":
		return p.Price
	case "medium":
		return (p.Price * 1.3 ) * 1.3
	case "large":
		return (p.Price * 1.6) + 2500
	default:
		return 0
	}

}

func factory(size Size, price float64) *Product {
	return &Product{
		Size:  size,
		Price: price,
	}
}