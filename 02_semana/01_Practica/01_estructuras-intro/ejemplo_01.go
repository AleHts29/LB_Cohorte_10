package main

import "fmt"

// Struc Persona
type Persona struct {
	Nombre string
	Edad   int
}

// Metodos
func (p Persona) Description() {
	fmt.Printf("%s tiene %d años de edad\n", p.Nombre, p.Edad)
}

func main() {
	p := Persona{
		Nombre: "Jhon",
		Edad:   30,
	}

	p.Description()
}
