package configuracao

import (
	"database/sql"
	"log"

	"github.com/eneassena10/go-api-estoque/internal/domain/product/controllers"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository"
	"github.com/eneassena10/go-api-estoque/internal/domain/product/repository/mysql"
	productservicecreate "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_create"
	productservicedelete "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_delete"
	productservicefindall "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_find_all"
	productservicefindbyid "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_find_by_id"
	productserviceupdatecount "github.com/eneassena10/go-api-estoque/internal/domain/product/service/product_service_update"
	dbmysql "github.com/eneassena10/go-api-estoque/pkg/conexao/db_mysql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var DB *sql.DB

type ServicePlay interface {
	MapRoutes(e *gin.Engine)
}

func Start() ServicePlay {
	controllersService := startDependence()

	s := NewHandlers(controllersService)
	return s
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".Env cant be load")
	}

	if DB == nil {
		database := dbmysql.IDataBase{}
		DB = database.Open()
	}
}

func startDependence() controllers.IControllers {

	// dependence of product
	dbOperation := mysql.NewMysqlRepository(DB)
	productRepository := repository.NewProductRepository(dbOperation)

	// services
	useCaseProductServiceGetById := productservicefindbyid.NewProductServiceFindByID(productRepository)
	useCaseProductServiceGetAll := productservicefindall.NewProductServiceFindAll(productRepository)
	useCaseProductServiceCreate := productservicecreate.NewProductServiceCreate(productRepository)
	useCaseProductServiceUpdateCount := productserviceupdatecount.NewProductServiceUpdateCount(productRepository)
	useCaseProductServiceDelete := productservicedelete.NewProductServiceDelete(productRepository)

	productController := controllers.NewControllers(
		useCaseProductServiceGetById,
		useCaseProductServiceGetAll,
		useCaseProductServiceCreate,
		useCaseProductServiceUpdateCount,
		useCaseProductServiceDelete,
	)

	return productController
}
