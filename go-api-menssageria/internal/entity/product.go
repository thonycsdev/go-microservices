package entity

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price, ID: uuid.New().String()}
}
