create database biblioteca;

use biblioteca;

-- -- Tabla Authors
create table Authors (
AuthorID INT primary key,
Name varchar(100) not null,
Coutry varchar(50)
);


-- -- Tabla Books
create table Books (
BookID INT primary key,
Title varchar(150) not null,
AuthorID INT,
Genre varchar(50),
PublishedYear YEAR,
foreign key (AuthorID) references Authors(AuthorID)
);

-- -- Tabla Readers
create table Readers (
ReaderID INT primary key,
Name varchar(100) not null,
MembershipDate DATE
);

-- -- Tabla Loans
create table Loans (
LoanID INT primary key,
BookID INT,
ReaderID INT,
LoanDate DATE,
ReturnDate DATE,
foreign key (BookID) references Books(BookID),
foreign key (ReaderID) references Readers(ReaderID)
);



-- Insertar datos en Authors
INSERT INTO Authors (AuthorID, Name, Coutry)
VALUES
	(1, 'Gabriel García Márquez', 'Colombia'),
    (2, 'George Orwell', 'Reino Unido');


-- Insertar datos en Books
INSERT INTO Books (BookID, Title, AuthorID, Genre, PublishedYear)
VALUES
    (1, 'Cien años de soledad', 1, 'Novela', 1967),
    (2, '1984', 2, 'Distopía', 1949);
    
-- Insertar datos en Readers
INSERT INTO Readers (ReaderID, Name, MembershipDate)
VALUES
    (1, 'Ana López', '2023-01-15'),
    (2, 'Carlos Pérez', '2023-03-10');

-- Insertar datos en Loans
INSERT INTO Loans (LoanID, BookID, ReaderID, LoanDate, ReturnDate)
VALUES
    (1, 1, 1, '2024-12-01', NULL),
    (2, 2, 2, '2024-11-20', '2024-12-01');
    
    
    
    
-- CONSULTAS
-- 1. ¿Cuáles son los libros disponibles (no prestados actualmente)?
select Books.BookID, Books.Title
FROM Books
LEFT JOIN Loans ON Books.BookID = Loans.BookID AND Loans.ReturnDate IS NULL
WHERE Loans.BookID IS NULL;



-- 2. ¿Cuántos libros hay de cada género?

-- 3. ¿Cuáles son los lectores que tienen libros prestados actualmente?

-- 4. Muestra los nombres de los autores y los títulos de sus libros.

-- 5. ¿Qué libros se publicaron antes del año 2000?


