package main

import (
	"errors"
	"fmt"
)

// Pila Definición de la estructura de la pila
type Pila struct {
	elementos []interface{}
}

// Método push: inserta un elemento en el tope de la pila
func (p *Pila) Push(e interface{}) {
	p.elementos = append(p.elementos, e)
}

// Método pop: elimina el elemento del tope de la pila y lo retorna
func (p *Pila) Pop() (interface{}, error) {
	// si la pila esta vacia retornamos error
	if p.IsEmpty() {
		return nil, errors.New("error: la pila esta vacia")
	}

	//	Accedo al elemento del tope de la pila
	elemento := p.elementos[len(p.elementos)-1]

	// Actualizamos el slice p.elementos para que contenga todos los elementos excepto el último.
	p.elementos = p.elementos[:len(p.elementos)-1]

	return elemento, nil
}

// Método isEmpty: retorna verdadero si la pila está vacía, falso en caso contrario
func (p *Pila) IsEmpty() bool {
	return len(p.elementos) == 0
}

// Método top: retorna el elemento del tope de la pila sin eliminarlo
func (p *Pila) Top() (interface{}, error) {
	// si la pila esta vacia retornamos error
	if p.IsEmpty() {
		return nil, errors.New("error: la pila esta vacia")
	}

	//	retornamos el elemento del tope de la pila sin eliminarlo
	return p.elementos[len(p.elementos)-1], nil
}

// Método size: retorna el número de elementos en la pila
func (p *Pila) Size() int {
	return len(p.elementos)
}

func main() {
	//instaciamos la PILA
	pila := &Pila{}

	// Ejemplo de uso
	pila.Push(10)
	pila.Push(20)
	pila.Push(30)
	pila.Push(40)

	//	Resuldos
	fmt.Println("Tamaño de la pila: ", pila.Size())

	elemento, err := pila.Top()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Elemento en el tope: ", elemento)
	}

	elemento2, _ := pila.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Elemento en el tope y que se elimina de la misma: ", elemento2)
	}

	fmt.Println("Tamaño de la pila: ", pila.Size())

	elemento3, err := pila.Top()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Elemento en el tope: ", elemento3)
	}
}
