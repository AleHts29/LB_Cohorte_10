package models

var Products = []Product{
	{ID: 1, Name: "Producto 1", Price: 19.99, Description: "Description 1", Category: "A"},
	{ID: 2, Name: "Producto 2", Price: 9.99, Description: "Description 2", Category: "A"},
	{ID: 3, Name: "Producto 3", Price: 10.05, Description: "Description 3", Category: "B"},
}

type Product struct {
	ID          int     `json:"ID"`
	Name        string  `json:"Name"`
	Price       float64 `json:"Price"`
	Description string  `json:"Description"`
	Category    string  `json:"Category"`
}
