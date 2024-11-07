package main

import "fmt"

// Interfaces Vehiculo
type Vehiculo interface {
	Mover()
}

// Estructura Auto que implementa Vehiculo
type Auto struct {
	Marca string
}

// Implementación del método Mover para Auto
func (a *Auto) Mover() {
	fmt.Println("El auto", a.Marca, "esta en movimiento")
}

// Estructura Bicicleta que también implementa Vehiculo
type Bicicleta struct {
	Marca string
}

func (a *Bicicleta) Mover() {
	fmt.Println("La Bici", a.Marca, "esta en movimiento")
}

// Función que utiliza la interfaz Vehiculo
func iniciarmovimiento(v Vehiculo) {
	v.Mover() // auto.Mover()
}

func main() {
	auto := Auto{
		Marca: "Toyota",
	}
	bici := Bicicleta{
		Marca: "Giant",
	}
	//auto.Mover()
	//bici.Mover()

	iniciarmovimiento(&auto)
	iniciarmovimiento(&bici)
}
