package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	// Textos a procesar
	texts := []string{
		"Este es el primer texto de ejemplo",
		"Go es un lenguaje concurrente y rápido",
		"Este texto tiene más palabras que los demás textos de prueba",
		"Último texto de prueba para este",
	}

	//	Canal para recibir los resultados
	resultChannel := make(chan int, len(texts))

	// WaitGroup para sincronizar las goroutines
	var wg sync.WaitGroup
	wg.Add(len(texts))

	for _, text := range texts {
		go countWords(text, &wg, resultChannel)
	}

	//	Esperamos a quye temiene todas las gorutinas
	wg.Wait()
	close(resultChannel)

	// Calcular el total de palabras
	totalWords := 0
	for count := range resultChannel {
		totalWords += count
	}

	// Mostrar resultados
	//fmt.Printf("Total de palabras sin iterar resultChannel: %i\n", resultChannel) // revisar
	fmt.Printf("Total de palabras: %d\n", totalWords)
}

func countWords(text string, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()

	//Dividir el texto en palabras
	words := strings.Fields(text)

	// Contar las palabras y enviar el resultado al canal
	result <- len(words)
}
