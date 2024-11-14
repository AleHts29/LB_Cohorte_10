package main

import "testing"

func TestInventario(t *testing.T) {
	inventario := NuevoInventario()

	// Casos de prueba para AgregarProducto
	casosAgregar := []struct {
		nombre   string
		cantidad int
	}{
		{"Manzanas", 10},
		{"Naranjas", 5},
		{"Bananas", 8},
		{"Manzanas", 2}, // Caso para probar acumulación de stock
	}

	//	iteramos los casos de pruebas
	//TEST:: 01 --> Agrega productos
	for _, caso := range casosAgregar {

		//	Ejecutar el test
		t.Run("Agregar Producto "+caso.nombre, func(t *testing.T) {
			inventario.AgregarProducto(caso.nombre, caso.cantidad)

			// inventario.stock[Manzanas] = 9
			if inventario.stock[caso.nombre] < caso.cantidad {
				t.Errorf("AgregarProducto fallo: se esperaba al menos %d de %s, pero se tiene %d",
					caso.cantidad, caso.nombre, inventario.stock[caso.nombre])
			} else {
				t.Logf("AgregarProducto pasó para %s con cantidad %d", caso.nombre, inventario.stock[caso.nombre])
			}
		})
	}

	//TEST:: 02 --> Verifica stock
	casosVerificar := []struct {
		nombre            string
		cantidad          int
		esperaSuficiencia bool
		esperaError       bool
	}{
		{"Manzanas", 5, true, false},
		{"Naranjas", 10, true, false},
		{"Bananas", 8, true, false},
		{"Manzanas", -1, false, true}, // Caso para probar error con cantidad negativa
	}

	for _, caso := range casosVerificar {
		t.Run("VerificarStock "+caso.nombre, func(t *testing.T) {
			tieneStock, err := inventario.VerificarStock(caso.nombre, caso.cantidad)

			if caso.esperaError {
				if err == nil {
					t.Errorf("Se esperaba un error para %s con cantidad %d, pero no se produjo", caso.nombre, caso.cantidad)
				} else {
					t.Logf("Se recibió el error esperado para %s con cantidad %d: %v", caso.nombre, caso.cantidad, err)
				}
			} else {
				if tieneStock != caso.esperaSuficiencia {
					t.Errorf("VerificarStock fallo: para %s con cantidad %d, se esperaba %v, pero se obtuvo %v",
						caso.nombre, caso.cantidad, caso.esperaSuficiencia, tieneStock)
				} else {
					t.Logf("VerificarStock pasó para %s con cantidad %d", caso.nombre, caso.cantidad)
				}
			}
		})

	}

}
