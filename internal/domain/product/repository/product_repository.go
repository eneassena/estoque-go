package repository

import (
	"fmt"

	"github.com/eneassena10/go-api-estoque/internal/domain/product/entities"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository/mysql"
)

//go:generate mockgen -source=./product_repository.go -destination=../../../test/mocks/mockgen/product_repository_mock.go -package=mockgen
type IProductRepository interface {
	ProductFindAll() *[]entities.Product
	ProductFindByID(product *entities.Product) *entities.Product
	ProductSave(product *entities.Product) error
	ProductUpdateCount(product *entities.Product) error
	ProductDestroy(product *entities.Product) error
}

type ProductRepository struct {
	dbOperation mysql.IDataBaseOperations
}

func NewProductRepository(dbOperation mysql.IDataBaseOperations) *ProductRepository {
	return &ProductRepository{dbOperation: dbOperation}
}

func (r *ProductRepository) GetProductsAll() *[]entities.Product {
	result, err := r.dbOperation.Query(mysql.QUERY_SELECT)
	if err != nil {
		return &[]entities.Product{}
	}

	var products []entities.Product
	for result.Next() {
		var product entities.Product
		err := result.Scan(&product.Name, &product.Price, &product.Count)
		if err != nil {
			return &[]entities.Product{}
		}
		products = append(products, product)
	}

	return &products
}

func (r *ProductRepository) GetProductByID(product *entities.Product) *entities.Product {
	result := r.dbOperation.QueryRow(mysql.QUERY_SELECT_BY_ID, product.ID)

	var productDb entities.Product
	if err := result.Scan(&productDb.Name, &productDb.Price, &productDb.Count); err != nil {
		return &entities.Product{}
	}
	return &productDb
}

func (r *ProductRepository) CreateProducts(product *entities.Product) error {
	result, err := r.dbOperation.Prepare(mysql.QUERY_INSERT, product.Name, product.Price, product.Count)
	if err != nil {
		return err
	}

	rowsAffected, errAffected := result.RowsAffected()

	if errAffected != nil && rowsAffected == 0 {
		return fmt.Errorf("error: %s, rows affected %d", errAffected.Error(), rowsAffected)
	}
	return nil
}

func (r *ProductRepository) UpdateProductsCount(product *entities.Product) error {
	// buscar o product a ser alterado
	productOld := r.GetProductByID(product)
	productOld.ID = product.ID

	// realizar alteração no produto
	countForRename := 0

	if productOld.Count+product.Count >= 0 {
		countForRename = productOld.Count + product.Count
	}

	productOld.Count = countForRename

	// salvar o product
	_, errUpdate := r.dbOperation.Exec(mysql.QUERY_UPDATE_COUNT, productOld.Count, productOld.ID)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (r *ProductRepository) DeleteProducts(product *entities.Product) error {
	// buscar p product a ser removido
	productOld := r.GetProductByID(product)
	productOld.ID = product.ID

	if productOld.Validator() {
		// fazer a remoção do product
		result, errUpdate := r.dbOperation.Exec(mysql.QUERY_DELETE, productOld.ID)

		if errUpdate != nil {
			return errUpdate
		}
		rowsAffected, errDelete := result.RowsAffected()
		if rowsAffected == 0 && errDelete != nil {
			return errDelete
		}
	}

	return nil
}
