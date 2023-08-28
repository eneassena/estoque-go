package controllers

import (
	"net/http"

	productservicecreate "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_create"
	productservicedelete "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_delete"
	productservicefindall "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_find_all"
	productservicefindbyid "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_find_by_id"
	productserviceupdatecount "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_update"
	"github.com/gin-gonic/gin"
)

type IControllers interface {
	GetProductsAll(ctx *gin.Context)
	GetProductsByID(ctx *gin.Context)
	CreateProducts(ctx *gin.Context)
	UpdateProductsCount(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

/*
Controllers de Products
*/
type Controllers struct {
	ProductServiceGetAll      productservicefindall.IProductServiceFindAll
	ProductServiceFindByID    productservicefindbyid.IProductServiceFindByID
	ProductServiceCreate      productservicecreate.IProductServiceCreate
	ProductServiceUpdateCount productserviceupdatecount.IProductServiceUpdateCount
	ProductServiceDelete      productservicedelete.IProductServiceDelete
}

func NewControllers(
	productServiceGetById productservicefindbyid.IProductServiceFindByID,
	productServiceGetAll productservicefindall.IProductServiceFindAll,
	productServiceCreate productservicecreate.IProductServiceCreate,
	productServiceUpdateCount productserviceupdatecount.IProductServiceUpdateCount,
	productServiceDelete productservicedelete.IProductServiceDelete,
) IControllers {
	return &Controllers{
		ProductServiceGetAll:      productServiceGetAll,
		ProductServiceFindByID:    productServiceGetById,
		ProductServiceCreate:      productServiceCreate,
		ProductServiceUpdateCount: productServiceUpdateCount,
		ProductServiceDelete:      productServiceDelete,
	}
}

func (c *Controllers) GetProductsAll(ctx *gin.Context) {
	products := c.ProductServiceGetAll.Execute()

	response := Response{Code: http.StatusOK, CountItems: len(products), Data: products}
	ctx.JSON(response.Code, response)
}

func (c *Controllers) GetProductsByID(ctx *gin.Context) {
	var productByID ProductRequestByID

	if err := ctx.BindJSON(&productByID); err != nil {
		response := Response{Code: http.StatusInternalServerError, Data: err.Error()}
		ctx.JSON(response.Code, response)
		return
	}
	productResult := c.ProductServiceFindByID.Execute(productservicefindbyid.ProductInputFindByID(productByID))
	if productResult == nil {
		response := Response{Code: http.StatusNoContent, Data: nil}
		ctx.JSON(response.Code, response)
		return
	}

	response := Response{Code: http.StatusOK, Data: productResult}
	ctx.JSON(response.Code, response)
}

func (c *Controllers) CreateProducts(ctx *gin.Context) {
	var productRequest ProductRequestCreate

	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		response := Response{Code: http.StatusInternalServerError, Data: err.Error()}
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	c.ProductServiceCreate.Execute(productservicecreate.ProductInputCreate(productRequest))

	response := Response{Code: http.StatusNoContent, Data: nil}
	ctx.JSON(response.Code, response)
}

func (c *Controllers) UpdateProductsCount(ctx *gin.Context) {
	var product ProductRequestUpdateCount
	if err := ctx.BindJSON(&product); err != nil {
		response := Response{Code: http.StatusInternalServerError, Data: err.Error()}
		ctx.JSON(response.Code, response)
		return
	}

	c.ProductServiceUpdateCount.Execute(productserviceupdatecount.ProductInputUpdateCount(product))

	response := Response{Code: http.StatusNoContent, Data: nil}
	ctx.JSON(response.Code, response)
}

func (c *Controllers) DeleteProducts(ctx *gin.Context) {
	var product ProductRequestByID

	if err := ctx.BindJSON(&product); err != nil {
		response := Response{Code: http.StatusInternalServerError, Data: err.Error()}
		ctx.JSON(response.Code, response)
		return
	}

	c.ProductServiceDelete.Execute(productservicedelete.ProductInputDelete(product))

	response := Response{Code: http.StatusNoContent, Data: nil}
	ctx.JSON(response.Code, response)
}
