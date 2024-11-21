package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	parChan := make(chan int, len(numbers))
	impChan := make(chan int, len(numbers))

	//	La espera de grupo (sync.WaitGroup) se utiliza para asegurarse de que ambas goroutines hayan terminado antes de que el programa finalice.
	var wg sync.WaitGroup

	wg.Add(2)

	go par(parChan, &wg)
	go impar(impChan, &wg)

	//	Pre-procesamiento
	for _, num := range numbers {
		if num%2 == 0 {
			parChan <- num
		} else {
			impChan <- num
		}
	}

	close(parChan)
	close(impChan)

	wg.Wait()

}

func par(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Printf("Numero Par: %d\n", item)
	}
}

func impar(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Printf("Numero Impar: %d\n", item)
	}
}
