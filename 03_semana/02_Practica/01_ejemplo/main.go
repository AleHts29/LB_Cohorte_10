package main

import "errors"

// Inventario representa un inventario de productos
type Inventario struct {
	stock map[string]int
}

// NuevoInventario crea un nuevo inventario vacío
func NuevoInventario() *Inventario {
	return &Inventario{stock: make(map[string]int)}
}

// AgregarProducto agrega una cantidad específica de un producto al inventario
func (i *Inventario) AgregarProducto(nombre string, cantidad int) {
	i.stock[nombre] += cantidad
}

// VerificarStock verifica si hay suficiente stock para un producto dado
func (i *Inventario) VerificarStock(nombre string, cantidad int) (bool, error) {
	if cantidad < 0 {
		return false, errors.New("cantidad no puede ser negativa")
	}

	if i.stock[nombre] >= cantidad {
		return true, nil
	}

	return false, nil
}
