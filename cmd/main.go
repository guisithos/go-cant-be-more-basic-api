package main

import (
	"go-cant-be-more-basic-api/controller"
	"go-cant-be-more-basic-api/database"
	"go-cant-be-more-basic-api/repo"
	"go-cant-be-more-basic-api/usage"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repo.NewProductRepository(dbConnection)
	ProductUseCase := usage.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/product", ProductController.CreateProduct)
	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":3030")

}
