package configuracao

import (
	"log"

	productControllers "github.com/eneassena10/go-api-estoque/internal/domain/product/controllers"
	"github.com/gin-gonic/gin"
)

type ControllersEntrypoint struct {
	ProductController productControllers.IControllers
}

func NewHandlers(productController any) *ControllersEntrypoint {
	return &ControllersEntrypoint{
		ProductController: productController.(productControllers.IControllers),
	}
}

func filterQuery(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		log.Println("field id required")
	} else {
		log.Println("fields id:" + id)
	}
}

func (h *ControllersEntrypoint) MapRoutes(router *gin.Engine) {
	g := router.Group("api/v1", filterQuery)
	g.GET("/product", h.ProductController.GetProductsByID)
	g.GET("products", h.ProductController.GetProductsAll)
	g.DELETE("/products", h.ProductController.DeleteProducts)
	g.POST("/products", h.ProductController.CreateProducts)
	g.PATCH("/products", h.ProductController.UpdateProductsCount)

}
