package productuc

type ProductInput struct {
	Name  string
	Price float64
}
type ProductOutput struct {
	ID    string
	Name  string
	Price float64
}

func CreateProductOutput(name string, price float64, id string) *ProductOutput {
	return &ProductOutput{Name: name, Price: price, ID: id}
}
