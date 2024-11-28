package main

import "fmt"

// Nodo representa cada nodo en el trie
type Nodo struct {
	hijos        map[rune]*Nodo
	finDePalabra bool
}

// Trie representa el trie completo
type Trie struct {
	raiz *Nodo
}

// NuevaTrie inicializa un trie vacío
func NuevaTrie() *Trie {
	return &Trie{
		raiz: &Nodo{hijos: make(map[rune]*Nodo)},
	}
}

// Insertar añade una palabra al trie
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

// ComienzaCon verifica si hay alguna palabra en el trie que comienza con el prefijo dado
func (t *Trie) ComienzaCon(prefijo string) bool {
	nodoActual := t.raiz
	for _, ch := range prefijo {
		if _, existe := nodoActual.hijos[ch]; !existe {
			return false
		}
		nodoActual = nodoActual.hijos[ch]
	}
	return true
}

// ---- Otras funcionalidades ----

// PalabrasConPrefijo devuelve todas las palabras en el trie que comienzan con el prefijo dado
func (t *Trie) PalabrasConPrefijo(prefijo string) []string {
	var resultados []string
	nodoActual := t.raiz
	for _, ch := range prefijo {
		if _, existe := nodoActual.hijos[ch]; !existe {
			return resultados
		}
		nodoActual = nodoActual.hijos[ch]
	}
	t.recogerPalabras(nodoActual, prefijo, &resultados)
	return resultados
}

// recogerPalabras es una función auxiliar para recolectar palabras desde un nodo dado
func (t *Trie) recogerPalabras(nodo *Nodo, prefijo string, resultados *[]string) {
	if nodo.finDePalabra {
		*resultados = append(*resultados, prefijo)
	}
	for ch, hijo := range nodo.hijos {
		t.recogerPalabras(hijo, prefijo+string(ch), resultados)
	}
}

func main() {
	trie := NuevaTrie()

	palabras := []string{"hola", "mundo", "holaquetal", "holanda"}
	for _, palabra := range palabras {
		trie.Insertar(palabra)
	}

	fmt.Println(trie.Buscar("hola"))    // true
	fmt.Println(trie.Buscar("mundo"))   // true
	fmt.Println(trie.Buscar("holanda")) // true
	fmt.Println(trie.Buscar("hol"))     // false

	// DESAFIO
	//fmt.Println(trie.ComienzaCon("hol")) // true
	//fmt.Println(trie.ComienzaCon("m"))   // true
	//fmt.Println(trie.ComienzaCon("mun")) // true
	//fmt.Println(trie.ComienzaCon("mu"))  // true
	//fmt.Println(trie.ComienzaCon("mua")) // false

	fmt.Println(trie.PalabrasConPrefijo("hol")) // ["hola", "holaquetal", "holanda"]
	fmt.Println(trie.PalabrasConPrefijo("m"))   // ["mundo"]
}

////En este código, se itera sobre cada carácter de la cadena "hola" y se convierte cada uno a su valor int32, que corresponde al valor Unicode del carácter.
//
//func main() {
//	// La cadena que deseas convertir
//	str := "你好"
//
//	// Convertir cada carácter de la cadena a int32
//	for _, char := range str {
//		fmt.Println(rune(char))
//	}
//}
