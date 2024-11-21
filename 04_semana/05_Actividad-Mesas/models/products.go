package models

type Product struct {
	Name  string
	Price float64
	Cant  int
}

func SumProducts(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.Price * float64(product.Cant)
	}
	return total
}
