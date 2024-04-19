package main

import (
	"fmt"
	"go-api-menssage/internal/infra/repository"
	"go-api-menssage/internal/useCases/productUC"
)

func main() {
	pr := repository.NewProductRepository()
	pucc := productuc.MakeCreateProductUserCase(pr)
	input := productuc.ProductInput{Name: "IPhone", Price: 3999}
	err := pucc.CreateProduct(&input)

	if err != nil {
		fmt.Println(err.Error())
	}

}
