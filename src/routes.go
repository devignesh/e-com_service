package routes

import (
	orders "e-com/src/order/controller"
	product "e-com/src/product/controller"

	"github.com/gin-gonic/gin"
)

var productController = new(product.ProductController)
var orderController = new(orders.OrderController)

func ProductRoutes(r *gin.Engine) {

	product := r.Group("/product")
	order := r.Group("/order")

	//product routes
	{
		product.POST("", productController.CreateProduct)
	}
	{
		product.GET("/:id", productController.GetProductByID)
	}
	{
		product.GET("", productController.GetProductList)
	}
	{
		product.PUT("/:id", productController.UpdateProduct)
	}

	//order routes

	{
		order.POST("", orderController.CreateOrder)
	}
	{
		order.GET("/:id", orderController.GetOrderByID)
	}

	{
		order.PUT("/:id", orderController.UpdateOrder)
	}

}
