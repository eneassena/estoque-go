package productservicecreate

import (
	"github.com/eneassena10/go-api-estoque/internal/domain/product/entities"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository"
)

type IProductServiceCreate interface {
	Execute(product ProductInputCreate)
}

type ProductServiceCreate struct {
	productRepository *repository.ProductRepository
}

func NewProductServiceCreate(repository *repository.ProductRepository) IProductServiceCreate {
	return &ProductServiceCreate{productRepository: repository}
}

type ProductInputCreate struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func (p *ProductServiceCreate) Execute(product ProductInputCreate) {
	productInput := entities.Product{
		Name:  product.Name,
		Price: product.Price,
		Count: product.Count,
	}
	_ = p.productRepository.CreateProducts(&productInput)
}
