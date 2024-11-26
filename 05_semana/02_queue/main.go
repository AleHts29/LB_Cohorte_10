package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements []interface{}
}

// IsEmpty: Retorna verdadero si la cola no contiene elementos y falso en caso contrario
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Enqueue: Inserta el elemento en el rabo de la cola
func (q *Queue) Enqueue(e interface{}) {
	q.elements = append(q.elements, e)
}

// Front: Retorna el elemento del frente de la cola
func (q *Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("error: la cola está vacía")
	}
	return q.elements[0], nil
}

// Size: Retorna la cantidad de elementos de la cola
func (q *Queue) Size() int {
	return len(q.elements)
}

// Dequeue: Elimina el elemento del frente de la cola y lo retorna
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("error: la cola está vacía")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]

	return element, nil
}

func main() {

	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Tamaño de la cola:", queue.Size())

	front, _ := queue.Front()
	fmt.Println("Elemento del frente:", front)

	dequeued, _ := queue.Dequeue()
	fmt.Println("Elemento desencolado:", dequeued)

	isEmpty := queue.IsEmpty()
	fmt.Println("¿La cola está vacía?", isEmpty)

	fmt.Println("Tamaño de la cola después de desencolar:", queue.Size())

}
