### Esplicacion
Para entender gráficamente la estructura de un trie (árbol de búsqueda de prefijos) con las palabras "hola", "mundo", "holaquetal" y "holanda", podemos dibujar el árbol de la siguiente manera:
```text

       (raíz)
       /  |  \
      h   m   n
     /    |    \
    o     u     d
   /      |      \
  l       n       a
 / \      |        (✓)
a   q     d        
|    \    |         
(✓)  u    o         
      \   |         
      e   (✓)       
       \
        t
         \
          a
           \
            l
             \
              (✓)

```


#### Usando byte:
```go
    palabra := "hola"
    for i := 0; i < len(palabra); i++ {
        fmt.Printf("%c ", palabra[i])  // Salida: h o l a 
    }
    
    palabraChino := "你好"
    for i := 0; i < len(palabraChino); i++ {
        fmt.Printf("%c ", palabraChino[i])  // Salida incorrecta debido a que cada carácter chino ocupa más de un byte
    }
```


-----

Ejemplo de inserción: Palabra “hola”

Estado inicial:

	•	Nodo raíz vacío (raiz).

Iteración por cada carácter:

	1.	Insertar ‘h’:
	•	No existe un nodo hijo correspondiente a 'h' en la raíz, por lo que se crea uno.
	•	Se avanza al nodo 'h'.
	2.	Insertar ‘o’:
	•	No existe un nodo hijo correspondiente a 'o' en 'h', por lo que se crea uno.
	•	Se avanza al nodo 'o'.
	3.	Insertar ‘l’:
	•	No existe un nodo hijo correspondiente a 'l' en 'o', por lo que se crea uno.
	•	Se avanza al nodo 'l'.
	4.	Insertar ‘a’:
	•	No existe un nodo hijo correspondiente a 'a' en 'l', por lo que se crea uno.
	•	Se avanza al nodo 'a'.
	5.	Marcar el fin de palabra:
	•	Se marca el nodo 'a' como un nodo de fin de palabra.

Resultado:

El Trie tiene una rama que representa la palabra “hola”.

```text
raíz
└── h
    └── o
        └── l
            └── a (fin de palabra)
```