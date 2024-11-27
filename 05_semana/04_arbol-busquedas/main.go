package main

import "fmt"

type Nodo struct {
	hijos        map[rune]*Nodo
	finDePalabra bool
}

type Trie struct {
	raiz *Nodo
}

// NuevaTrie inicializa un trie vacío
func NuevaTrie() *Trie {
	return &Trie{
		raiz: &Nodo{hijos: make(map[rune]*Nodo)},
	}
}

func (t *Trie) Insertar(palabra string) {
	nodoActual := t.raiz
	for _, ch := range palabra {
		if _, existe := nodoActual.hijos[ch]; !existe {
			nodoActual.hijos[ch] = &Nodo{hijos: make(map[rune]*Nodo)}
		}
		nodoActual = nodoActual.hijos[ch]
	}
	nodoActual.finDePalabra = true
}

// Buscar verifica si una palabra está en el trie
func (t *Trie) Buscar(palabra string) bool {
	nodoActual := t.raiz
	for _, ch := range palabra {
		if _, existe := nodoActual.hijos[ch]; !existe {
			return false
		}
		nodoActual = nodoActual.hijos[ch]
	}
	return nodoActual.finDePalabra
}

func main() {
	trie := NuevaTrie()

	palabras := []string{"hola", "mundo", "holaquetal", "holanda"}
	for _, palabra := range palabras {
		trie.Insertar(palabra)
	}

	fmt.Println(trie.Buscar("hola"))       // true
	fmt.Println(trie.Buscar("holaquetal")) // true
	fmt.Println(trie.Buscar("estonoesta")) // true

	// DESAFIO
	fmt.Println(trie.ComienzaCon("hol")) // true
	fmt.Println(trie.ComienzaCon("m"))   // true
	fmt.Println(trie.ComienzaCon("mun")) // true
	fmt.Println(trie.ComienzaCon("mua")) // false

	fmt.Println(trie.PalabrasConPrefijo("hol")) // ["hola", "holaquetal", "holanda"]
	fmt.Println(trie.PalabrasConPrefijo("m"))   // ["mundo"]
}
