package interfaces

import (
	"go-api-menssage/internal/entity"
)

type InterfaceProductRepository interface {
	Create(product *entity.Product) error
	FindAll() (*[]entity.Product, error)
}
