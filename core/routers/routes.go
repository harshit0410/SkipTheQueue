package routers

import (
	"skipthequeue/core/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)

	router.GET("/outlets", controllers.FindAllOutlet)
	router.POST("/outlets", controllers.CreateOutlet)
	router.GET("/outlets/:id", controllers.FindOutletById)
	router.PATCH("/outlets/:id", controllers.UpdateOutlet)
	router.DELETE("/outlets/:id", controllers.DeleteOutlet)

	router.POST("/outlet/dish", controllers.CreateDish)
	router.GET("/outlet/:outletId/dish", controllers.FindAllDish)
	router.GET("/outlet/:outletId/dish/:dishId", controllers.FindDishById)
	router.PATCH("/outlet/:outletId/dish/:dishId", controllers.UpdateDish)
	router.DELETE("/outlet/:outletId/dish/:dishId", controllers.DeleteDish)

	router.POST("/order", controllers.CreateOrder)
	router.PATCH("/order/:orderId", controllers.UpdateOrderStatus)
	router.GET("/order/:orderId", controllers.FindOrderById)
}
