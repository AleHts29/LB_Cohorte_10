Aquí tienes un ejercicio práctico para practicar bases de datos SQL. Este ejercicio abarca conceptos básicos como creación de tablas, inserción de datos, consultas simples, y consultas más avanzadas.

---

## **Ejercicio: Sistema de Gestión de Biblioteca**

### **Descripción**
Vas a diseñar un sistema básico para gestionar una biblioteca. Este sistema debe manejar información sobre libros, autores, lectores y préstamos de libros.

---

### **1. Crea las tablas necesarias**
Escribe las sentencias SQL para crear las siguientes tablas:

1. **Tabla `Authors`**:
    - `AuthorID` (INT, PRIMARY KEY)
    - `Name` (VARCHAR(100), NO NULL)
    - `Country` (VARCHAR(50))

2. **Tabla `Books`**:
    - `BookID` (INT, PRIMARY KEY)
    - `Title` (VARCHAR(150), NO NULL)
    - `AuthorID` (INT, FOREIGN KEY que referencia `Authors.AuthorID`)
    - `Genre` (VARCHAR(50))
    - `PublishedYear` (YEAR)

3. **Tabla `Readers`**:
    - `ReaderID` (INT, PRIMARY KEY)
    - `Name` (VARCHAR(100), NO NULL)
    - `MembershipDate` (DATE)

4. **Tabla `Loans`**:
    - `LoanID` (INT, PRIMARY KEY)
    - `BookID` (INT, FOREIGN KEY que referencia `Books.BookID`)
    - `ReaderID` (INT, FOREIGN KEY que referencia `Readers.ReaderID`)
    - `LoanDate` (DATE)
    - `ReturnDate` (DATE)

---

### **2. Inserta datos en las tablas**

Inserta datos en las tablas utilizando sentencias `INSERT`. Aquí tienes algunos ejemplos de datos:

- **Autores (Authors):**
    - `AuthorID`: 1, `Name`: "Gabriel García Márquez", `Country`: "Colombia"
    - `AuthorID`: 2, `Name`: "George Orwell", `Country`: "Reino Unido"

- **Libros (Books):**
    - `BookID`: 1, `Title`: "Cien años de soledad", `AuthorID`: 1, `Genre`: "Novela", `PublishedYear`: 1967
    - `BookID`: 2, `Title`: "1984", `AuthorID`: 2, `Genre`: "Distopía", `PublishedYear`: 1949

- **Lectores (Readers):**
    - `ReaderID`: 1, `Name`: "Ana López", `MembershipDate`: "2023-01-15"
    - `ReaderID`: 2, `Name`: "Carlos Pérez", `MembershipDate`: "2023-03-10"

- **Préstamos (Loans):**
    - `LoanID`: 1, `BookID`: 1, `ReaderID`: 1, `LoanDate`: "2024-12-01", `ReturnDate`: NULL
    - `LoanID`: 2, `BookID`: 2, `ReaderID`: 2, `LoanDate`: "2024-11-20", `ReturnDate`: "2024-12-01"

---

### **3. Realiza consultas SQL**

Responde las siguientes preguntas utilizando consultas SQL:

1. ¿Cuáles son los libros disponibles (no prestados actualmente)?
2. ¿Cuántos libros hay de cada género?
3. ¿Cuáles son los lectores que tienen libros prestados actualmente?
4. Muestra los nombres de los autores y los títulos de sus libros.
5. ¿Qué libros se publicaron antes del año 2000?

---

### **4. Nivel avanzado**

1. Encuentra los lectores que han tomado prestados libros de un autor específico (por ejemplo, "Gabriel García Márquez").
2. Calcula el número total de préstamos que se han hecho para cada libro.
3. Obtén los lectores que no han tomado ningún libro en préstamo.





--- 
### RELACIONES ENTRE TABLAS
```sql
create table Empleados(
id_empleado INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(100),
apellido VARCHAR(100),
fecha_nacimiento DATE,
direccion VARCHAR(255),
telefono VARCHAR(15),
id_departamento INT,
foreign key (id_departamento) references Departamentos(id_departamento)
);

CREATE TABLE Departamentos (
    id_departamento INT PRIMARY KEY AUTO_INCREMENT,
    nombre_departamento VARCHAR(100)
);
```