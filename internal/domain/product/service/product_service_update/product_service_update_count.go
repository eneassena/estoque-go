package productserviceupdate

import (
	"github.com/eneassena10/go-api-estoque/internal/domain/product/entities"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository"
)

type IProductServiceUpdateCount interface {
	Execute(product ProductInputUpdateCount)
}

type ProductInputUpdateCount struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}
type ProductServiceUpdateCount struct {
	productRepository *repository.ProductRepository
}

func NewProductServiceUpdateCount(repository *repository.ProductRepository) IProductServiceUpdateCount {
	return &ProductServiceUpdateCount{productRepository: repository}
}

func (p *ProductServiceUpdateCount) Execute(product ProductInputUpdateCount) {
	productInput := entities.Product{
		ID:    product.ID,
		Count: product.Count,
	}
	_ = p.productRepository.UpdateProductsCount(&productInput)

}
