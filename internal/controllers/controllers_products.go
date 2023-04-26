package controllers

import (
	"net/http"

	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
)

/*
Controllers de Products
*/
type Controllers struct {
	fileStore store.IStore
}

func NewControllers(fileStore store.IStore) IControllers {
	return &Controllers{fileStore: fileStore}
}

func (c *Controllers) GetProductsAll(ctx *gin.Context) {
	var productFileJson *[]Product
	if err := c.fileStore.Read(&productFileJson); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, productFileJson})
}

func (c *Controllers) GetProductsByID(ctx *gin.Context) {
	var p *Product
	if err := ctx.BindJSON(&p); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	product := c.getProductByProductID(p)
	if product == nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, nil})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) CreateProducts(ctx *gin.Context) {
	var productRequest *Product
	err := ctx.ShouldBindJSON(&productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	var productList []*Product
	err = c.fileStore.Read(&productList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	if len(productList) > 0 {
		productRequest.ID = productList[len(productList)-1].ID + 1
	} else {
		productRequest.ID = 1
	}
	productList = append(productList, productRequest)
	err = c.fileStore.Write(&productList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, productList})
}

func (c *Controllers) UpdateProductsQuantidade(ctx *gin.Context) {
	var product *Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	var productList []*Product
	err := c.fileStore.Read(&productList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	for _, p := range productList {
		if p.ID == product.ID && (p.Quantidade+product.Quantidade) >= 0 {
			p.Quantidade += product.Quantidade
			product = p
			break
		}
	}

	if err := c.fileStore.Write(&productList); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) DeleteProducts(ctx *gin.Context) {
	var product *Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	pList, err := c.loadListProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	for i, p := range pList {
		if p.ID == product.ID {
			pList = append(pList[:i], pList[i+1:]...)
		}
	}

	err = c.fileStore.Write(&pList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *Controllers) getProductByProductID(product *Product) *Product {
	products, _ := c.loadListProducts()
	for _, p := range products {
		if p.ID == product.ID {
			return p
		}
	}
	return nil
}

func (c *Controllers) loadListProducts() ([]*Product, error) {
	var p []*Product
	if err := c.fileStore.Read(&p); err != nil {
		return p, err
	}
	return p, nil
}
