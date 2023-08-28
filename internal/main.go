package main

import (
	"github.com/eneassena10/go-api-estoque/internal/configuracao"
	"github.com/gin-gonic/gin"
)

func main() {

	// create instance
	route := gin.Default()

	configuracao.Start().MapRoutes(route)

	// start app
	route.Run(":8080")
}
