package main

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

// Ejemplo 2 - Pruebas para Métodos de una Estructura
/*
Supongamos que tenemos una estructura CuentaBancaria con métodos para depositar y retirar dinero. Queremos escribir pruebas unitarias para asegurarnos de que esos métodos funcionen como se espera.
*/

type CuentaBancaria struct {
	saldo float64
}

// Depositar
func (b *CuentaBancaria) Depositar(monto float64) {
	if monto > 0 {
		b.saldo += monto
	}
}

// Retirar
func (b *CuentaBancaria) Retirar(monto float64) error {
	if monto > b.saldo {
		return fmt.Errorf("saldo insufficiente")
	}
	b.saldo -= monto
	return nil
}

// ConsultarSaldo
func (b *CuentaBancaria) ConsultarSaldo() float64 {
	return b.saldo
}
