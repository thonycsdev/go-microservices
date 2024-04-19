package productuc

import (
	"encoding/json"
	"go-api-menssage/internal/communication"
	"go-api-menssage/internal/entity"
	"go-api-menssage/internal/infra/interfaces"
)

type CreateProductUseCase struct {
	repository interfaces.InterfaceProductRepository
}

func (puc CreateProductUseCase) CreateProduct(input *ProductInput) error {
	product := entity.NewProduct(input.Name, input.Price)
	payload, err := json.Marshal(product)
	puc.repository.Create(product)
	if err != nil {
		return err
	}
	communication.SendDataToRabbitMQ(&payload)
	return nil
}

func MakeCreateProductUserCase(repository interfaces.InterfaceProductRepository) *CreateProductUseCase {
	puc := new(CreateProductUseCase)
	puc.repository = repository
	return puc
}
