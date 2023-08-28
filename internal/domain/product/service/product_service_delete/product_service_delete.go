package productservicedelete

import (
	"github.com/eneassena10/go-api-estoque/internal/domain/product/entities"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository"
)

type IProductServiceDelete interface {
	Execute(product ProductInputDelete)
}

type ProductInputDelete struct {
	ID int `json:"id"`
}
type ProductServiceDelete struct {
	productRepository *repository.ProductRepository
}

func NewProductServiceDelete(repository *repository.ProductRepository) IProductServiceDelete {
	return &ProductServiceDelete{productRepository: repository}
}

func (p *ProductServiceDelete) Execute(product ProductInputDelete) {
	productInput := entities.Product{
		ID: product.ID,
	}
	_ = p.productRepository.DeleteProducts(&productInput)

}
