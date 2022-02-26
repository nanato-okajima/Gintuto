package app

import "Gintuto/api/controllers/products"

func mapUrls() {
	router.GET("/products/:product_id", products.GetProduct)
	router.POST("/products", products.CreateProduct)
	router.PUT("/products/:product_id", products.UpdateProduct)
	router.PATCH("/products/:product_id", products.UpdateProduct)
}
