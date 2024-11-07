package main

import "fmt"

type Vehiculo02 struct {
	km     float64
	tiempo float64
}

func (v Vehiculo02) detalle() {
	fmt.Printf("km:\t%f\ntiempo:\t%f\n", v.km, v.tiempo)
}

// --> Auto
type Auto02 struct {
	v     Vehiculo02
	color string
}

func (a *Auto02) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}
func (a *Auto02) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

// --> Moto
type Moto struct {
	v Vehiculo02
}

func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 80
}

func (m *Moto) Detalle() {
	fmt.Printf("\nV:\tMoto\n")
	m.v.detalle()
}

func main() {
	//	Auto
	auto := Auto02{}
	auto.Correr(30)
	auto.Detalle()

	//	Moto
	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()

}
