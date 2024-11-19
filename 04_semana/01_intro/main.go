package main

import (
	"fmt"
	"sync"
)

// 01_ Ejemplo con time.Sleep(1 * time.Second)
//func main() {
//	//go func() {
//	//	fmt.Println("Soy otra gorutina")
//	//}()
//
//	go testFunction() // Esta función se ejecuta en otra gorutina y no bloquea al hilo principal
//
//	//time.Sleep(1 * time.Second)
//
//	//inicio := time.Now()
//	//simularUnaTareaLarga()
//	//duracion := time.Since(inicio) // Registrar el tiempo final y calcular la duración
//	//fmt.Printf("Soy el hilo principal - la tarea tardo %v tiempo", duracion)
//
//	fmt.Println("Soy el hilo principal")
//}
//
//func testFunction() {
//	fmt.Println("Soy otra gorutina")
//	//	Actualizar registros en una DB.
//	//	Envia una notificacion a otro programa
//}
//
//func simularUnaTareaLarga() {
//	for i := 0; i < 10000000000; i++ {
//		// Simulación de una tarea larga
//	}
//}

// // 02_ Ejemplo con waitGroup
var wg sync.WaitGroup

func main() {
	wg.Add(1) // establecemos el contador regresivo en 1
	n := 3

	//ch := make(chan int)
	go func() {
		res := testFunction(n) // ejecutamos la función en otra gorutina
		//ch <- res
		fmt.Println(res)
	}()

	//fmt.Println(<-ch)

	fmt.Println("Soy la gorutina principal")
	wg.Wait() // esperamos que la gorutina principal termine
}

func testFunction(num int) int {
	res := num * 2

	fmt.Println("Soy otra gorutina")

	wg.Done() // establecemos el contador regresivo en 0
	return res
}

// 03_ Ejemplo con Canales
//func main() {
//	n := 3
//
//	ch := make(chan int)
//	go multiplicarPorDos(n, ch) // ejecutamos la función en otra gorutina
//
//	fmt.Println(<-ch) // esperamos a que la gorutina termine y lea el canal y lo imprima
//
//	fmt.Println("Soy la gorutina principal")
//}
//
//func multiplicarPorDos(num int, ch chan int) {
//	res := num * 2
//	ch <- res
//}
