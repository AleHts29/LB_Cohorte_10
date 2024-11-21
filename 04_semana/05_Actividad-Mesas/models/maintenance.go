package models

type Maintenance struct {
	Name  string
	Price float64
}

func SumMaintenance(maintenances []Maintenance) float64 {
	total := 0.0
	for _, product := range maintenances {
		total += product.Price
	}
	return total
}
