package main

import "fmt"

/*
Registro de estudiantes
Una universidad necesita registrar a los estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos, de la siguiente manera:
	- Nombre: [Nombre del alumno]
	- Apellido: [Apellido del alumno]
	- DNI: [DNI del alumno]
	- Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos. Para ello es necesario generar una estructura Alumno con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle.

*/

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a *Alumno) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha de ingreso: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	p := Alumno{
		Nombre:   "Jhon",
		Apellido: "Doe",
		DNI:      12345678,
		Fecha:    "15/08/2023",
	}

	p.Detalle()
}
