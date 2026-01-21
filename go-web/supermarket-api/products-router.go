package main

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	products := router.Group("/products")
	{
		products.GET("", getAllProductsHandler) // "" ao invés de "/"
		products.GET("/:id", getProductByIDHandler)
		products.POST("", createProductHandler) // "" ao invés de "/"
		products.PUT("/:id", updateProductHandler)
		products.DELETE("/:id", deleteProductHandler)

	}
}
