# Especialización en Back End III

## Punteros y Canales en Go

### Ejercitación Individual

**Nivel de complejidad:** Medio 🔥🔥

#### Calcular Precio

Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento. Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda, y para optimizar la velocidad, requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.

Se requieren tres estructuras:

1. Productos:
    - Nombre
    - Precio
    - Cantidad

2. Servicios:
    - Nombre
    - Precio
    - Minutos trabajados

3. Mantenimiento:
    - Nombre
    - Precio

Y se requieren tres funciones:

1. Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).

2. Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada. Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).

3. Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).
