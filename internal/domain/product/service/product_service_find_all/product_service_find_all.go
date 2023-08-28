package productservicefindall

import "github.com/eneassena10/go-api-estoque/internal/domain/product/repository"

type IProductServiceFindAll interface {
	Execute() []ProductOutputFindAll
}

type ProductServiceFindAll struct {
	productRepository *repository.ProductRepository
}

type ProductOutputFindAll struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func NewProductServiceFindAll(repository *repository.ProductRepository) IProductServiceFindAll {
	return &ProductServiceFindAll{productRepository: repository}
}

func (p *ProductServiceFindAll) Execute() []ProductOutputFindAll {
	productOutput := []ProductOutputFindAll{}
	products := p.productRepository.GetProductsAll()

	for _, p := range *products {
		productOutput = append(productOutput, ProductOutputFindAll{
			Name:  p.Name,
			Price: p.Price,
			Count: p.Count,
		})
	}
	return productOutput
}
