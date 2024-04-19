package repository

import (
	"fmt"
	"go-api-menssage/internal/entity"
	"go-api-menssage/internal/infra/database"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
func (pr *ProductRepository) Create(product *entity.Product) error {

	query := `INSERT INTO products(id, name, price) values ($1,$2,$3)`
	conn, err := database.MakeConn()
	defer conn.Close()
	result := conn.QueryRow(query, product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	if result.Err() != nil {
		return result.Err()
	}
	fmt.Println("Product Added!")
	return nil

}
func (pr *ProductRepository) FindAll() (*[]entity.Product, error) { return nil, nil }
