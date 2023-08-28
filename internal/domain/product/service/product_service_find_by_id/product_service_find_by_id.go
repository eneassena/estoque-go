package productservicefindbyid

import (
	"github.com/eneassena10/go-api-estoque/internal/domain/product/entities"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository"
)

type IProductServiceFindByID interface {
	Execute(product ProductInputFindByID) *ProductOutputFindByID
}

type ProductServiceFindByID struct {
	productRepository *repository.ProductRepository
}

func NewProductServiceFindByID(repository *repository.ProductRepository) IProductServiceFindByID {
	return &ProductServiceFindByID{productRepository: repository}
}

type ProductOutputFindByID struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}
type ProductInputFindByID struct {
	ID int `json:"id"`
}

func (p *ProductServiceFindByID) Execute(productById ProductInputFindByID) *ProductOutputFindByID {
	productRepository := entities.Product{ID: productById.ID}

	productOutput := p.productRepository.GetProductByID(&productRepository)

	if !productOutput.Validator() {
		return nil
	}

	return &ProductOutputFindByID{
		Name:  productOutput.Name,
		Price: productOutput.Price,
		Count: productOutput.Count,
	}
}
