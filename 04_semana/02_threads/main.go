package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Obtener el número de procesadores disponibles
	numCPUs := runtime.NumCPU()
	fmt.Printf("Número de CPUs disponibles: %d\n", numCPUs)

	// Ajustar GOMAXPROCS al número de CPUs disponibles
	runtime.GOMAXPROCS(3)
	fmt.Printf("GOMAXPROCS configurado a: %d\n", runtime.GOMAXPROCS(0))

	// Usar un WaitGroup para esperar a que todas las gorutinas terminen
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1) // Incrementa el contador de gorutinas
		go func(id int) {
			defer wg.Done() // Decrementa el contador de gorutinas

			fmt.Printf("Gorutina %d ejecutándose\n", id)
		}(i)
	}

	// Espera para que las gorutinas terminen - 20 hilos de ejecucion
	wg.Wait()
	fmt.Println("Todos las gorutinas terminaron")

}
