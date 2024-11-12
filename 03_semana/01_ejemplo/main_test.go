package main

import "testing"

func TestSuma(t *testing.T) {
	//Given
	var a, b int = 3, 5
	esperado := 8

	// Then
	resultado := Sumar(a, b)

	//Assert
	if resultado != esperado {
		t.Errorf("Suma(%d, %d) = %d; se esperaba %d", a, b, resultado, esperado)
	} else {
		t.Log("TestSuma: Suma correcta")
	}
}

// 02_ Pruebas Unitarias para CuentaBancaria

// al ejecutar las pruebas con `go test -v`, se obtendrá una salida detallada mostrando qué prueba específica se está ejecutando y si pasa o falla, facilitando el diagnóstico de problemas.

// TEST_01 - TestDepositar
func TestDepositar(t *testing.T) {
	t.Run("Depositar 500.00 en cuenta vacia", func(t *testing.T) {
		//	Given
		cuenta := CuentaBancaria{}
		deposit := 500.00

		//	Then
		cuenta.Depositar(deposit)

		//	Assert
		saldoEsperado := deposit
		if cuenta.ConsultarSaldo() != saldoEsperado {
			t.Errorf("Error en TestDepositar: ConsultarSaldo() = %f; se esperaba %f ", cuenta.ConsultarSaldo(), saldoEsperado)
		} else {
			t.Logf("TestDepositar: Depositar 500 en cuenta vacía pasó correctamente")
		}
	})
}

// TEST_02 - TestRetirar
func TestRetirar(t *testing.T) {
	t.Run("Retirar 200.00 de una cuenta con un saldo de 500.00", func(t *testing.T) {
		//	Given
		cuenta := CuentaBancaria{}
		retirar := 200.00

		//	Then
		cuenta.Depositar(500.00)
		err := cuenta.Retirar(retirar)
		if err != nil {
			t.Errorf("Error en TestRetirar: Se produjo un error inesperado: %v", err)
		}

		//	Assert
		saldoEsperado := 300.00
		if cuenta.ConsultarSaldo() != saldoEsperado {
			t.Errorf("Error en TestRetirar: ConsultarSaldo() = %f; se esperaba %f ", cuenta.ConsultarSaldo(), saldoEsperado)
		} else {
			t.Logf("TestRetirar: Retirar 200.00 de una cuenta con 500.00 de saldo - Paso correctamente")
		}
	})
}

// TEST_03 - TestRetirarSaldoInsuficiente
func TestRetirarSaldoInsuficiente(t *testing.T) {
	t.Run("Intento de retiro de 200 en cuenta con 100 de saldo", func(t *testing.T) {
		//	Given
		cuenta := CuentaBancaria{}
		retirar := 200.00

		//	Then
		cuenta.Depositar(100.00)
		err := cuenta.Retirar(retirar)
		if err == nil {
			t.Error("Error en TestRetirarSaldoInsuficiente: Se esperaba un error por saldo insuficiente, pero no se produjo")
		} else {
			t.Logf("TestRetirarSaldoInsuficiente: Error por saldo insuficiente fue correctamente detectado")
		}
	})
}
