package main

import "fmt"

// Nodo representa un nodo en el árbol BST
type Nodo struct {
	valor int
	izq   *Nodo
	der   *Nodo
}

// ArbolBST representa el árbol de búsqueda binaria
type ArbolBST struct {
	raiz *Nodo
}

// Insertar agrega un nuevo valor al árbol BST
func (arbol *ArbolBST) Insertar(valor int) {
	arbol.raiz = insertarNodo(arbol.raiz, valor)
}

// Insertar Nodo
func insertarNodo(node *Nodo, valor int) *Nodo {
	if node == nil {
		return &Nodo{valor: valor}
	}

	if valor < node.valor {
		node.izq = insertarNodo(node.izq, valor)
	} else {
		node.der = insertarNodo(node.der, valor)
	}
	return node
}

// InOrder realiza un recorrido in-order del árbol BST
func (arbol *ArbolBST) InOrder() {
	inOrderRecorrido(arbol.raiz)
	fmt.Println()
}

func inOrderRecorrido(nodo *Nodo) {
	if nodo != nil {
		inOrderRecorrido(nodo.izq)
		fmt.Printf("%d ", nodo.valor)
		inOrderRecorrido(nodo.der)
	}
}

// Buscar verifica si un valor existe en el árbol BST
func (arbol *ArbolBST) Buscar(valor int) bool {
	return buscarNodo(arbol.raiz, valor)
}

func buscarNodo(nodo *Nodo, valor int) bool {
	if nodo == nil {
		return false
	}

	if valor < nodo.valor {
		return buscarNodo(nodo.izq, valor)
	} else if valor > nodo.valor {
		return buscarNodo(nodo.der, valor)
	} else {
		return true
	}
}

func main() {
	valores := []int{53, 12, 60, 5, 23, 57, 80, 21, 26, 71}

	arbol := &ArbolBST{}
	for _, valor := range valores {
		arbol.Insertar(valor)
	}

	fmt.Println("Recorrido in-order del árbol BST:")
	arbol.InOrder()

	buscarValor := 40
	fmt.Printf("¿El valor %d está en el árbol? %v\n", buscarValor, arbol.Buscar(buscarValor))

	buscarValor = 26
	fmt.Printf("¿El valor %d está en el árbol? %v\n", buscarValor, arbol.Buscar(buscarValor))
}
