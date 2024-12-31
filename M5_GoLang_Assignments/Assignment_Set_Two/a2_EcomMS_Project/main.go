package main

import (
	"a2-ecomms-project/config"
	"a2-ecomms-project/controller"
	"a2-ecomms-project/middleware"
	"a2-ecomms-project/repository"
	"a2-ecomms-project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	repo := &repository.ProductRepository{DB: db}
	service := &service.ProductService{Repo: repo}
	controller := &controller.ProductController{Service: service}

	r := gin.Default()

	r.Use(middleware.Logging())
	r.Use(middleware.JWTAuth())
	r.Use(middleware.RateLimit())

	r.POST("/product", controller.AddProduct)
	r.GET("/product/:id", controller.GetProductByID)
	r.PUT("/product/:id", controller.UpdateStock)
	r.DELETE("/product/:id", controller.DeleteProduct)

	r.Run(":8080")
}
