package models

import "math"

type Service struct {
	Name             string
	PricePerHalfHour float64
	MinutesWorked    int
}

func SumServices(services []Service) float64 {
	total := 0.0
	for _, service := range services {
		// Esto representa cuántas "medias horas" (bloques de 30 minutos) se trabajaron. Ceil hace redondeo hacia arriba
		billedHours := math.Ceil(float64(service.MinutesWorked) / 30.0)
		total += service.PricePerHalfHour * billedHours
	}
	return total
}
