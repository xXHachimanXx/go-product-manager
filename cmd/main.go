package main

import (
	"product-manager/controller"
	"product-manager/db"
	"product-manager/repository"
	"product-manager/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	productUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
