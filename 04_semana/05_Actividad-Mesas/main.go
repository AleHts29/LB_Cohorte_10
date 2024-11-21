package main

import (
	"C_07A/models"
	"fmt"
	"sync"
)

func main() {
	// Test data
	products := []models.Product{
		{"Producto 1", 10.0, 2},
		{"Producto 2", 5.0, 3},
		{"Producto 3", 8.0, 1},
	}
	services := []models.Service{
		{"Service 1", 10.0, 1800},
		{"Service 2", 5.0, 300},
		{"Service 3", 8.0, 520},
	}
	maintenances := []models.Maintenance{
		{"Maintenance 1", 10.0},
		{"Maintenance 2", 15.0},
		{"Maintenance 3", 20.0},
	}

	//	Variables to storage results
	var totalProducts, totalServices, totalMaintenance float64

	//	waitGroup to wait for gorutines to complete
	var wg sync.WaitGroup
	wg.Add(3)

	//    goroutine to calculate total products
	go func() {
		defer wg.Done()
		totalProducts = models.SumProducts(products)
	}()

	//    goroutine to calculate total services
	go func() {
		defer wg.Done()
		totalServices = models.SumServices(services)
	}()

	//    goroutine to calculate total maintenance
	go func() {
		defer wg.Done()
		totalMaintenance = models.SumMaintenance(maintenances)
	}()

	//    wait for all goroutines to complete
	wg.Wait()

	//    calculate total invoice
	total := totalProducts + totalServices + totalMaintenance

	//    print total invoice
	// Display results
	fmt.Printf("Total Products: %.2f\n", totalProducts)
	fmt.Printf("Total Services: %.2f\n", totalServices)
	fmt.Printf("Total Maintenances: %.2f\n", totalMaintenance)
	fmt.Printf("Final Total: %.2f\n", total)
}
