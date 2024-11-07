package main

import (
	"encoding/json"
	"fmt"
)

type PersonaJson struct {
	PrimerNombre string `json:"primer_nombre"`
	Apellido     string `json:"apellido"`
	Telefono     string `json:"telefono"`
	Direccion    string `json:"direccion"`
	Email        string `json:"-"` // ignora el campo
	Edad         int    `json:"edad"`
}

func main() {
	p := PersonaJson{
		PrimerNombre: "Alejandro",
		Apellido:     "Huertas",
		Telefono:     "111111122",
		Direccion:    "Calle 123",
		Email:        "test@test.com",
		Edad:         37,
	}

	miJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(miJson))
}
